package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fitraditya/cloudfs"
	_ "github.com/fitraditya/cloudfs/storage/dropbox"
	_ "github.com/fitraditya/cloudfs/storage/filesystem"
	_ "github.com/fitraditya/cloudfs/storage/gdrive"
	_ "github.com/fitraditya/cloudfs/storage/s3"
)

func main() {
	// Init fs
	fs := cloudfs.NewFs("dropbox")

	// Open original file
	fp, err := fs.OpenFile("/users.csv", os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(fp)

	// Print original file
	for s.Scan() {
		fmt.Println(s.Text())
	}

	fp.Close()

	// Reopen original file
	fp, err = fs.OpenFile("/users.csv", os.O_WRONLY|os.O_TRUNC, 0700)
	if err != nil {
		panic(err)
	}

	// Override original file
	_, err = fp.WriteString("Assalamualaikum!")
	if err != nil {
		panic(err)
	}

	fp.Close()

	// Open modified file
	fp, err = fs.OpenFile("/users.csv", os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}

	s = bufio.NewScanner(fp)

	// Print modified file
	for s.Scan() {
		fmt.Println(s.Text())
	}

	fp.Close()
}
