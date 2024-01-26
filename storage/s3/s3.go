package s3

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	s3 "github.com/fclairamb/afero-s3"
	"github.com/fitraditya/cloudfs/storage"
	"github.com/spf13/afero"
)

type option struct {
	region       string
	bucket       string
	accessKey    string
	accessSecret string
	fs           afero.Fs
}

func (s *option) Fs() afero.Fs {
	return s.fs
}

// Region option setter
func Region(s string) storage.Option {
	return func(b *option) {
		b.region = s
	}
}

// Bucket option setter
func Bucket(s string) storage.Option {
	return func(b *option) {
		b.bucket = s
	}
}

// AccessKey option setter
func AccessKey(s string) storage.Option {
	return func(b *option) {
		b.accessKey = s
	}
}

// AccessSecret option setter
func AccessSecret(s string) storage.Option {
	return func(b *option) {
		b.accessSecret = s
	}
}

func init() {
	storage.Register("s3", func(opts ...storage.Option) (storage.Storage, error) {
		s := &option{}

		for _, opt := range opts {
			switch f := opt.(type) {
			case func(*option):
				f(s)
			default:
				return nil, errors.New("s3: unknown option")
			}
		}

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(s.region),
			Credentials: credentials.NewStaticCredentials(s.accessKey, s.accessSecret, ""),
		})
		if err != nil {
			return nil, err
		}

		s.fs = s3.NewFs(s.bucket, sess)

		return s, nil
	})
}
