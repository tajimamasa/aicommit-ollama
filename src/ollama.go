package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func checkOllamaVersion(host string) (string, error) {
	resp, err := http.Get(host + "api/version")
	if err != nil {
		return "", fmt.Errorf("error fetching Ollama version: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var versionResp VersionResponse
	if err := json.Unmarshal(body, &versionResp); err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}

	return versionResp.Version, nil
}
