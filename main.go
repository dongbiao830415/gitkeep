package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func listAll(path string, curHier int) {
	readerInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(readerInfos) == 0 {
		goto git_keep
	}
	for _, info := range readerInfos {
		if info.IsDir() && info.Name() != ".git" && info.Name() != ".svn" {
			listAll(filepath.Join(path, info.Name()), curHier+1)
		}
	}
	return

git_keep:
	gitKeep := filepath.Join(path, ".gitkeep")
	f, err := os.Create(gitKeep)
	if err != nil {
		fmt.Printf("can't create file %s:%s\n", gitKeep, err.Error())
		return
	}
	f.Close()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s directory\n", filepath.Base(os.Args[0]))
		return
	}
	dir := os.Args[1]
	listAll(dir, 0)
}
