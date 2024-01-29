package gdrive

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"

	gdrive "github.com/fclairamb/afero-gdrive"
	"github.com/fclairamb/afero-gdrive/oauthhelper"
	"github.com/fitraditya/cloudfs/storage"
	"github.com/spf13/afero"
	"golang.org/x/oauth2"
)

type option struct {
	clientId     string
	clientSecret string
	token        string
	fs           afero.Fs
}

func (s *option) Fs() afero.Fs {
	return s.fs
}

// ClientId option setter
func ClientId(s string) storage.Option {
	return func(b *option) {
		b.clientId = s
	}
}

// ClientSecret option setter
func ClientSecret(s string) storage.Option {
	return func(b *option) {
		b.clientSecret = s
	}
}

// Token option setter
func Token(s string) storage.Option {
	return func(b *option) {
		b.token = s
	}
}

func init() {
	storage.Register("gdrive", func(opts ...storage.Option) (storage.Storage, error) {
		s := &option{}

		for _, opt := range opts {
			switch f := opt.(type) {
			case func(*option):
				f(s)
			default:
				return nil, errors.New("gdrive: unknown option")
			}
		}

		if s.clientId == "" {
			return nil, errors.New("gdrive: missing google client id")
		}

		if s.clientSecret == "" {
			return nil, errors.New("gdrive: missing google client secret")
		}

		if s.token == "" {
			return nil, errors.New("gdrive: missing access token")
		}

		helper := oauthhelper.Auth{
			ClientID:     s.clientId,
			ClientSecret: s.clientSecret,
			Authenticate: func(url string) (string, error) {
				return "", gdrive.ErrNotSupported
			},
		}

		token, err := base64.StdEncoding.DecodeString(s.token)
		if err != nil {
			return nil, err
		}

		helper.Token = new(oauth2.Token)
		json.Unmarshal(token, helper.Token)

		client, err := helper.NewHTTPClient(context.Background())
		if err != nil {
			return nil, err
		}

		fs, err := gdrive.New(client)
		if err != nil {
			return nil, err
		}

		s.fs = fs.AsAfero()

		return s, nil
	})
}
