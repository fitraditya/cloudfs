package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fitraditya/cloudfs"
	_ "github.com/fitraditya/cloudfs/storage/dropbox"
	_ "github.com/fitraditya/cloudfs/storage/filesystem"
	_ "github.com/fitraditya/cloudfs/storage/s3"
)

func main() {
	// Open original file
	fp, err := cloudfs.OpenFile("test.txt", os.O_RDONLY)
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
	fp, err = cloudfs.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		panic(err)
	}

	// Override original file
	_, err = fp.WriteString("Hello world!")
	if err != nil {
		panic(err)
	}

	fp.Close()

	// Open modified file
	fp, err = cloudfs.OpenFile("test.txt", os.O_RDONLY)
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
