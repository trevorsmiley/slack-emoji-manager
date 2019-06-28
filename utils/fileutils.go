package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func RemoveFolderContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer func() {
		err := d.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func DownloadFile(filepath, uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	return err
}

func CreateOrClearDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, os.ModePerm)
	} else {
		return RemoveFolderContents(dir)
	}
}