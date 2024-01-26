package storage

import (
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
	Fs() afero.Fs
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
		log.For("storage", "register").Fatal("could not register nil storage")
	}

	if _, dup := storages[name]; dup {
		log.For("storage", "register").Fatalf("could not register storage twice (%s)", name)
	}

	storages[name] = s
}
