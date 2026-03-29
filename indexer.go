package main

import (
	"os"
	"path/filepath"
)

func indexDirectory(rootDir string, jobs chan string) {
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			jobs <- path
		}
		return nil
	})
}
