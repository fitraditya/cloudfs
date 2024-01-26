package dropbox

import (
	"errors"
	"io/fs"

	dropbox "github.com/fclairamb/afero-dropbox"
	"github.com/fitraditya/cloudfs/storage"
	"github.com/spf13/afero"
)

type option struct {
	token string
	fs    *dropbox.Fs
}

func (s *option) Open(name string) (afero.File, error) {
	return s.fs.Open(name)
}

func (s *option) OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error) {
	return s.fs.OpenFile(name, flag, perm)
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

		s.fs = dropbox.NewFs(s.token)

		return s, nil
	})
}
