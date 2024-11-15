package gtfs

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

type feedFiles struct {
	RouteFile []byte
}

func Parse(url string) *feedFiles {
	reader, _ := FetchZip(url)
	feed := ReadZip(reader)
	return feed
}

func FetchZip(url_or_path string) (*zip.Reader, error) {
	var body []byte

	file, err := os.Open(url_or_path)
	if err != nil {
		body, _ = download(url_or_path)
	} else {
		body, _ = io.ReadAll(file)
		defer file.Close()
	}
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, err
	}
	return zipReader, nil
}

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ReadZip(r *zip.Reader) *feedFiles {
	feed := feedFiles{}

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()
		if f.Name == "routes.txt" {
			feed.RouteFile, _ = io.ReadAll(rc)
		}
	}

	return &feed
}
