package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

var Slashes string

func CurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

func ProcessDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func FullPath(filename string) string {
	return filepath.Dir(filename) + Slashes + filepath.Base(filename) + filepath.Ext(filename)
}

func FindDir(dir string) string {
	dir = FullPath(dir)
	currentDir := ProcessDir()
	cwd, _ := os.Getwd()

	if _, err := os.Stat(currentDir + Slashes + dir + Slashes); err == nil {
		currentDir = filepath.Dir(currentDir + Slashes + dir + Slashes)
	} else if _, err := os.Stat(cwd + Slashes + dir + Slashes); err == nil {
		currentDir = filepath.Dir(cwd + Slashes + dir + Slashes)
	}

	return currentDir + Slashes
}

func init() {
	Slashes = filepath.Dir("/")
}