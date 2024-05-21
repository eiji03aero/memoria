package util

import (
	"io"
	"net/http"
	"os"
)

type DownloadFileDTO struct {
	FileName string
	URL      string
}

func DownloadFile(dto DownloadFileDTO) (file *os.File, err error) {
	filePath := "*-" + dto.FileName
	file, err = os.CreateTemp("", filePath)
	if err != nil {
		return
	}

	resp, err := http.Get(dto.URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// TODO: probably we should make it streaming?
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return
	}

	return
}
