package search

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const fileSearchURL = "https://api.trace.moe/search"

func File(filePath string) error {
	var err error

	if !filepath.IsAbs(filePath) {
		filePath, err = filepath.Abs(filePath)
		if err != nil {
			return fmt.Errorf("convert file path %q to absolute path: %w", filePath, err)
		}
	}

	slog.Debug("file path", "path", filePath)

	res, err := post(filePath)
	if err != nil {
		return err
	}

	var traceResp TraceMoeResponse
	err = json.Unmarshal(res, &traceResp)
	if err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}

	fmt.Printf("%+v", traceResp)

	return err
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

	slog.Debug("request", "method", req.Method, "header", req.Header)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, fmt.Errorf("send post request to %q: %w", fileSearchURL, err)
	}
	defer wclose(resp.Body, &err)

	if err != nil {
		return res, fmt.Errorf("send post request to %q: %w", fileSearchURL, err)
	}
	defer wclose(resp.Body, &err)

	res, err = io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("read repsonse: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		slog.Warn("response", "status", resp.Status, "message", res)
		return res, errors.New(resp.Status)
	}

	return res, err
}

func wclose(w io.Closer, err *error) {
	// respect the existed err
	if e := w.Close(); *err == nil {
		*err = e
	}
}
