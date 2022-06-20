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
	opts := []string{}
	if len(os.Args) > 1 {
		initDir = os.Args[len(os.Args)-1]
		opts = os.Args[1 : len(os.Args)-1]
	}
	if initDir == "-h" || initDir == "-help" || initDir == "h" || initDir == "help" {
		fmt.Println(`-D == dont show directories; -d show only directories`)
		os.Exit(0)
	}
	if initDir[0] != '/' {
		initDir = filepath.Join(basepath, initDir)
	}
	recurSearch(initDir, opts, MAX_DEEP)
}

func inOpts(o string, opts []string) bool {
	for _, v := range opts {
		if v == o {
			return true
		}
	}
	return false
}

func recurSearch(initDir string, opts []string, deep int) {
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
			if !inOpts("-D", opts) {
				fmt.Println(fullDirPath)
			}
			recurSearch(fullDirPath, opts, deep-1)
		}
		if !f.IsDir() && !inOpts("-d", opts) {
			fullDirPath := filepath.Join(initDir, f.Name())
			fmt.Println(fullDirPath)
		}
	}
}
