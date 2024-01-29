package filesystem

import (
	"errors"

	"github.com/fitraditya/cloudfs/storage"
	"github.com/spf13/afero"
)

type option struct {
	basePath string
	fs       afero.Fs
}

func (s *option) Fs() afero.Fs {
	return s.fs
}

// BasePath option setter
func BasePath(s string) storage.Option {
	return func(b *option) {
		b.basePath = s
	}
}

func init() {
	storage.Register("filesystem", func(opts ...storage.Option) (storage.Storage, error) {
		s := &option{}

		for _, opt := range opts {
			switch f := opt.(type) {
			case func(*option):
				f(s)
			default:
				return nil, errors.New("filesystem: unknown option")
			}
		}

		if s.basePath == "" {
			return nil, errors.New("filesystem: missing base path")
		}

		s.fs = afero.NewBasePathFs(afero.NewOsFs(), s.basePath)

		return s, nil
	})
}
