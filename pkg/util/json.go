package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func JSONMarshalToReader(v interface{}) (io.Reader, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func JSONUnmarshalFromReader(src io.Reader, dst interface{}) error {
	b, err := ioutil.ReadAll(src)
	if err != nil {
		return fmt.Errorf("json unmarshal from reader byte read error: %w", err)
	}

	log.Debug(string(b))
	if err := json.Unmarshal(b, dst); err != nil {
		return fmt.Errorf("json unmarshal from reader error: %w", err)
	}
	return nil
}

// make a json request from body (struct) then unmarshal into vPtr (struct)
func JSONHTTPRequestWithToken(httpClient *http.Client, method, url, token string, extraHeaders map[string]string, body interface{}, vPtr interface{}) error {
	log.WithField("url", url).Debug("request")

	buf, err := JSONMarshalToReader(body)
	if err != nil {
		return fmt.Errorf("client marshal error: %w", err)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("client req error: %w", err)
	}
	defer resp.Body.Close()

	if vPtr == nil {
		return nil
	}

	if err := JSONUnmarshalFromReader(resp.Body, vPtr); err != nil {
		return fmt.Errorf("client resp error: %w", err)
	}

	return nil
}
