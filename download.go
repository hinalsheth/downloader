package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func DownloadFiles(urls []string) {
	// create a waitgroup
	var wg sync.WaitGroup

	// add the number of urls to the waitgroup
	wg.Add(len(urls))

	for i, url := range urls {
		go downloadFile(i, url, &wg)
	}

	// wait for all downloads to complete
	wg.Wait()

	fmt.Println("All files downloaded")
}

func downloadFile(i int, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// get the file
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading", url, "-", err)
		return
	}

	// close body stream
	defer resp.Body.Close()

	// create the file
	filename := "downloaded_" + fmt.Sprintf("%d", i) + ".jpg"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file", filename, "-", err)
		return
	}

	// close file at the end of the function
	defer file.Close()

	// write the body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file", filename, "-", err)
		return
	}

	fmt.Println("Downloaded", filename, "from", url)
}
