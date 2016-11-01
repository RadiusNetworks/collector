package output

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pganalyze/collector/state"
	"github.com/pganalyze/collector/util"
)

type s3UploadResponse struct {
	Location string
	Bucket   string
	Key      string
}

func uploadToS3(grant state.Grant, logger *util.Logger, compressedData bytes.Buffer, filename string) (string, error) {
	var err error

	if !grant.Valid {
		return "", fmt.Errorf("Error - can't upload without valid S3 grant")
	}

	if grant.S3URL == "" && grant.LocalDir != "" {
		location := grant.LocalDir + filename
		err = os.MkdirAll(filepath.Dir(location), 0755)
		if err != nil {
			logger.PrintError("Error creating target directory: %s", err)
			return "", err
		}

		err = ioutil.WriteFile(location, compressedData.Bytes(), 0644)
		if err != nil {
			logger.PrintError("Error writing local file: %s", err)
			return "", err
		}
		return location, nil
	}

	logger.PrintVerbose("Successfully prepared S3 request - size of request body: %.4f MB", float64(compressedData.Len())/1024.0/1024.0)

	var formBytes bytes.Buffer

	writer := multipart.NewWriter(&formBytes)

	for key, val := range grant.S3Fields {
		err = writer.WriteField(key, val)
		if err != nil {
			return "", err
		}
	}

	part, _ := writer.CreateFormFile("file", filename)
	_, err = part.Write(compressedData.Bytes())
	if err != nil {
		return "", err
	}

	writer.Close()

	req, err := http.NewRequest("POST", grant.S3URL, &formBytes)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("Bad S3 upload return code %s (should be 201 Created), body: %s", resp.Status, body)
	}

	var s3Resp s3UploadResponse
	err = xml.Unmarshal(body, &s3Resp)
	if err != nil {
		return "", err
	}

	return s3Resp.Key, nil
}
