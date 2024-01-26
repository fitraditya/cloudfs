package storage

import (
	"io/fs"
	"sync"

	"github.com/obrel/go-lib/pkg/log"
	"github.com/spf13/afero"
)

var (
	storages = map[string]Factory{}
	lock     sync.RWMutex
)

// Storage structure
type Storage interface {
	Open(name string) (afero.File, error)
	OpenFile(name string, flag int, perm fs.FileMode) (afero.File, error)
}

// Option interface
type Option interface{}

// Factory for storage with given option
type Factory func(...Option) (Storage, error)

// New backup for given storage type
func New(name string, opts ...Option) (Storage, error) {
	lock.RLock()
	s, ok := storages[name]

	defer lock.RUnlock()

	if !ok {
		log.For("storage", "new").Fatal("storage not found")
	}

	return s(opts...)
}

// Register given storage
func Register(name string, s Factory) {
	lock.Lock()

	defer lock.Unlock()

	if s == nil {
		log.For("storage", "register").Fatal("storage: could not register nil type")
	}

	if _, dup := storages[name]; dup {
		log.For("storage", "register").Fatalf("storage: could not register storage twice (%s)", name)
	}

	storages[name] = s
}
