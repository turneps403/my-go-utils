package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	files := []string{}
	if len(os.Args) == 2 {
		reader := bufio.NewReader(os.Stdin)
		for {
			fpath, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			fpath = strings.Trim(fpath, " \n")
			files = append(files, fpath)
		}
	}

	if len(files) == 0 && len(os.Args) < 3 {
		fmt.Println(`there shold be at least three arguments`)
		os.Exit(0)
	}
	searchReStr, replaceReStr := parseRe(os.Args[1])
	searchRe := regexp.MustCompile(searchReStr)

	files = files[1:]
	for _, from := range files {
		fileInfo, err := os.Stat(from)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if !fileInfo.IsDir() {
			to := searchRe.ReplaceAllString(from, replaceReStr)
			if from != to {
				// fmt.Printf("%v -> %v\n", from, to)
				err := os.Rename(from, tp)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func parseRe(p string) (string, string) {
	i := 0
	for ; i < len(p)-1; i++ {
		if p[i] != ' ' {
			break
		}
	}
	if i > len(p)-5 || p[i] != 's' {
		log.Fatal("pattern is broken")
	}
	i++

	splitChar := p[i]
	i++

	first := ""
	for j := i; j < len(p)-1; j++ {
		if p[j] == splitChar && p[j-1] != '\\' {
			first = p[i:j]
			i = j + 1
			break
		}
	}
	if first == "" {
		log.Fatal("pattern is broken")
	}

	last := ""
	for j := len(p) - 1; j >= i; j-- {
		if p[j] == splitChar {
			last = p[i:j]
			break
		}
	}

	return first, last
}
