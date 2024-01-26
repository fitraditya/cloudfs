package cloudfs

import (
	"os"
	"time"

	"github.com/fitraditya/cloudfs/config"
	"github.com/fitraditya/cloudfs/storage"
	"github.com/obrel/go-lib/pkg/log"
	"github.com/spf13/afero"
)

// Fs is the cloud filesystem.
type Fs struct {
	name     string
	rootPath string
	storage  storage.Storage
}

func NewFs(name string) *Fs {
	fs := &Fs{
		name: name,
	}

	str, err := storage.New(name, config.GetStorageOptions(name)...)
	if err != nil {
		log.For("cloudfs", "newfs").Error(err)
		return nil
	}

	fs.storage = str

	return fs
}

// Name of the fs.
func (fs *Fs) Name() string {
	return fs.name
}

// Create creates a file.
func (fs *Fs) Create(name string) (afero.File, error) {
	return fs.storage.Fs().Create(name)
}

// Mkdir creates a directory.
func (fs *Fs) Mkdir(name string, perm os.FileMode) error {
	return fs.storage.Fs().Mkdir(name, perm)
}

// MkdirAll creates a directory and all parent directories if necessary.
func (fs *Fs) MkdirAll(name string, perm os.FileMode) error {
	return fs.storage.Fs().MkdirAll(name, perm)
}

// Open a file for reading.
func (fs *Fs) Open(name string) (afero.File, error) {
	return fs.storage.Fs().Open(name)
}

// OpenFile opens a file.
func (fs *Fs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	return fs.storage.Fs().OpenFile(name, flag, perm)
}

// Remove removes a file.
func (fs *Fs) Remove(name string) error {
	return fs.storage.Fs().Remove(name)
}

// RemoveAll removes all files inside a directory.
func (fs *Fs) RemoveAll(name string) error {
	return fs.storage.Fs().RemoveAll(name)
}

// Rename renames a file.
func (fs *Fs) Rename(oldname, newname string) error {
	return fs.storage.Fs().Rename(oldname, newname)
}

// Stat fetches the file info.
func (fs *Fs) Stat(name string) (os.FileInfo, error) {
	return fs.storage.Fs().Stat(name)
}

// Chmod is not supported.
func (fs *Fs) Chmod(name string, mode os.FileMode) error {
	return ErrNotSupported
}

// Chown is not supported.
func (fs *Fs) Chown(name string, uid int, gid int) error {
	return ErrNotSupported
}

// Chtimes is not supported because dropbox doesn't support simply changing a time.
func (fs *Fs) Chtimes(name string, _ time.Time, mtime time.Time) error {
	return ErrNotSupported
}

// SetRootDirectory defines a base directory
// This is mostly useful to isolate tests and can most probably forgotten
// for most use-cases.
func (fs *Fs) SetRootDirectory(fullPath string) {
	fs.rootPath = fullPath
}
