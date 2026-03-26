package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var BASE_URL = "https://api.aiozpin.network/api/"

func executeRequest(req *http.Request) (map[string]interface{}, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func checkSuccessStatus(resp map[string]interface{}, defaultErr string) error {
	if status, ok := resp["status"].(string); ok && status != "success" {
		if message, ok := resp["message"].(string); ok {
			return errors.New(message)
		}
		return errors.New(defaultErr)
	}
	return nil
}

func MakeRequest(url string, method string, payloadParam string, jwt string) (map[string]interface{}, error) {

	url = BASE_URL + url

	payload := strings.NewReader(payloadParam)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+jwt)
	req.Header.Add("Content-Type", "application/json")

	respRaw, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	status, _ := respRaw["status"].(string)
	if status != "success" {
		return nil, fmt.Errorf("Internal Server Error")
	}

	data, ok := respRaw["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response data")
	}

	return data, nil
}

func UploadPinningFile(url string, filePath string, apiKey string, secretKey string) (map[string]interface{}, error) {
	url = BASE_URL + url

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("pinning_api_key", apiKey)
	req.Header.Add("pinning_secret_key", secretKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	if err := checkSuccessStatus(resp, "Internal Server Error"); err != nil {
		return nil, err
	}

	return resp, nil
}

func MakePinningRequest(url string, method string, payloadParam string, apiKey string, secretKey string) (map[string]interface{}, error) {
	url = BASE_URL + url

	payload := strings.NewReader(payloadParam)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("pinning_api_key", apiKey)
	req.Header.Add("pinning_secret_key", secretKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	if err := checkSuccessStatus(resp, "Internal Server Error"); err != nil {
		return nil, err
	}

	return resp, nil
}

func GetFileFromURL(ctx context.Context, url string) (string, string, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", "", "", err
	}

	client := &http.Client{}
	downloadResp, err := client.Do(req)
	if err != nil {
		return "", "", "", err
	}
	defer downloadResp.Body.Close()

	if downloadResp.StatusCode != http.StatusOK {
		return "", "", "", fmt.Errorf("download failed with status code %d", downloadResp.StatusCode)
	}

	contentType := downloadResp.Header.Get("Content-Type")
	if contentType == "" {
		return "", "", "", fmt.Errorf("downloaded file missing content-type")
	}

	fileName := path.Base(req.URL.Path)
	if fileName == "." || fileName == "/" || fileName == "" {
		fileName = "downloaded-file"
	}

	tempFile, err := os.CreateTemp("", "aioz-pin-*")
	if err != nil {
		return "", "", "", err
	}
	tempPath := tempFile.Name()

	_, err = io.Copy(tempFile, downloadResp.Body)
	closeErr := tempFile.Close()
	if err != nil {
		return "", "", "", err
	}
	if closeErr != nil {
		return "", "", "", closeErr
	}
	return tempPath, fileName, contentType, nil
}
