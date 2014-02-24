package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func logf(fmt string, v ...interface{}) {
	log.Printf(fmt, v...)
}

// Return the value of an ENV var, or the fallback value if the ENV var is empty/undefined.
func getenvSys(key, fallback string, mapping func(string) string) string {
	if v := mapping(key); v != "" {
		return v
	}
	return fallback	
}

var getenv = func(key, fallback string) string {
	return getenvSys(key, fallback, os.Getenv )
}

// Check if the connection to tcp://addr is readable.
var read = func(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err = conn.Read(make([]byte, 1)); err != nil {
		return err
	}
	return nil
}

// Check if an addr can be successfully connected.
var ping = func(addr string) bool {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// Download the url to the dest path.
var download = func(dest, url string) error {
	rsp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	f, err := os.Create(fmt.Sprintf("%s.download", dest))
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())

	if _, err := io.Copy(f, rsp.Body); err != nil {
		// TODO: display download progress?
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Rename(f.Name(), dest); err != nil {
		return err
	}
	return nil
}

// Get latest release tag name (e.g. "v0.6.0") from a repo on GitHub.
var getLatestReleaseName = func(url string) (string, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	var t []struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(rsp.Body).Decode(&t); err != nil {
		return "", err
	}
	if len(t) == 0 {
		return "", fmt.Errorf("no releases found")
	}
	return t[0].TagName, nil
}
