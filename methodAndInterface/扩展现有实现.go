package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}
type GitHubResponse []struct {
	FullName string `json:"full_name"`
	Size     int    `json:"size"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, repo := range resp {
		fmt.Println(repo.FullName)
	}
	return len(p), nil
}

func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}
