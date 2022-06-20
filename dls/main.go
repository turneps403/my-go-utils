package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	MAX_DEEP   = 100
)

func main() {
	initDir := "./"
	if len(os.Args) > 1 {
		initDir = os.Args[len(os.Args)-1]
	}
	if initDir == "-h" || initDir == "-help" || initDir == "h" || initDir == "help" {
		fmt.Println(`shows list of directories and subdirectories by path and max deep 100`)
		os.Exit(0)
	}
	if initDir[0] != '/' {
		initDir = filepath.Join(basepath, initDir)
	}
	recurSearch(initDir, MAX_DEEP)
}

func recurSearch(initDir string, deep int) {
	if deep < 1 {
		return
	}
	files, err := ioutil.ReadDir(initDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fullDirPath := filepath.Join(initDir, f.Name())
			fmt.Println(fullDirPath)
			recurSearch(fullDirPath, deep-1)
		}
	}
}
