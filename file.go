package cloudfs

import (
	"os"

	"github.com/fitraditya/cloudfs/config"
	"github.com/fitraditya/cloudfs/storage"
	"github.com/spf13/afero"
)

// Opens the file.
func OpenFile(path string, flag int) (afero.File, error) {
	var perm os.FileMode = 0700

	if flag == os.O_RDONLY {
		perm = 0400
	}

	cfs, err := storage.New(config.GetStorage(), config.GetStorageOption()...)
	if err != nil {
		return nil, NewIOError(err.Error())
	}

	fp, err := cfs.OpenFile(path, flag, perm)
	if err != nil {
		return nil, NewIOError(err.Error())
	}

	return fp, nil
}

// Closes the file.
func Close(fp afero.File) (err error) {
	defer func() { _ = fp.Close() }()

	return nil
}
