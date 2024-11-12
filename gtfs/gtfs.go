package gtfs

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Parse(url string) string {
	reader, _ := FetchZip(url)
	ReadZip(reader)
	return "Hello, world."
}

func FetchZip(url string) (*zip.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	zipReader, err := zip.NewReader(bytes.NewReader(body), resp.ContentLength)
	if err != nil {
		return nil, err
	}
	return zipReader, nil
}

func ReadZip(r *zip.Reader) {
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()
		scanner := bufio.NewScanner(rc)
		for scanner.Scan() {
			//fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
