package main

import (
	"fmt"
	"os"
)

func main() {
	// example cute cat files
	urls := []string{
		"https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Cat_August_2010-4.jpg/2880px-Cat_August_2010-4.jpgx",
		"https://upload.wikimedia.org/wikipedia/commons/6/68/Orange_tabby_cat_sitting_on_fallen_leaves-Hisashi-01A.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/b/bc/Juvenile_Ragdoll.jpg",
	}

	if len(os.Args) > 1 {
		urls = os.Args[1:]
	} else {
		fmt.Println("No urls provided via cli, using default urls")
	}

	DownloadFiles(urls)
}
