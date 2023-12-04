package search

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/acgtools/trace-moe-go/internal/util"
)

const (
	fileSearchURL = "https://api.trace.moe/search?anilistInfo"
)

func File(filePath string) (*TraceMoeResponse, error) {
	var err error

	res, err := post(filePath)
	if err != nil {
		return nil, err
	}

	var traceResp TraceMoeResponse
	err = json.Unmarshal(res, &traceResp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &traceResp, err
}

func post(filePath string) ([]byte, error) {
	var (
		err error
		res []byte
	)

	img, err := os.Open(filePath)
	if err != nil {
		return res, fmt.Errorf("open file %q: %w", filePath, err)
	}

	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	part, err := mw.CreateFormFile("image", filepath.Base(filePath))
	if err != nil {
		return res, fmt.Errorf("create form data: %w", err)
	}

	_, err = io.Copy(part, img)
	if err != nil {
		return res, fmt.Errorf("upload file: %w", err)
	}

	err = mw.Close()
	if err != nil {
		return res, fmt.Errorf("close multipart wirter: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, fileSearchURL, body)
	if err != nil {
		return res, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, fmt.Errorf("send post request to %q: %w", fileSearchURL, err)
	}
	defer util.Close(resp.Body, &err)

	if err != nil {
		return res, fmt.Errorf("send post request to %q: %w", fileSearchURL, err)
	}
	defer util.Close(resp.Body, &err)

	res, err = io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("read repsonse: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return res, errors.New(resp.Status)
	}

	return res, err
}
