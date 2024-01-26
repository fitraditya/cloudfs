package filesystem

import (
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/fitraditya/cloudfs/storage"
	"github.com/spf13/afero"
)

type option struct {
	basePath string
	fs       afero.Fs
}

func (s *option) Open(name string) (afero.File, error) {
	return s.fs.Open(filepath.Join(s.basePath, name))
}

func (s *option) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	return s.fs.OpenFile(filepath.Join(s.basePath, name), flag, perm)
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

		s.fs = afero.NewOsFs()

		return s, nil
	})
}
