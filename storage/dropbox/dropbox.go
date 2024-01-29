package dropbox

import (
	"errors"

	dropbox "github.com/fclairamb/afero-dropbox"
	"github.com/fitraditya/cloudfs/storage"
	"github.com/spf13/afero"
)

type option struct {
	token string
	fs    afero.Fs
}

func (s *option) Fs() afero.Fs {
	return s.fs
}

// Token option setter
func Token(s string) storage.Option {
	return func(b *option) {
		b.token = s
	}
}

func init() {
	storage.Register("dropbox", func(opts ...storage.Option) (storage.Storage, error) {
		s := &option{}

		for _, opt := range opts {
			switch f := opt.(type) {
			case func(*option):
				f(s)
			default:
				return nil, errors.New("dropbox: unknown option")
			}
		}

		if s.token == "" {
			return nil, errors.New("dropbox: missing access token")
		}

		s.fs = dropbox.NewFs(s.token)

		return s, nil
	})
}
