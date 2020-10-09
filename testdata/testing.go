package testdata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

// basepath is the root directory of this package
var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

// path returns the absolute path the given relative file or directory path,
func path(relPath string) string {
	if filepath.IsAbs(relPath) {
		return relPath
	}

	return filepath.Join(basepath, relPath)
}

// GetGolden is a function to get golden file
func GetGolden(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := ioutil.ReadFile(path(filename + ".golden"))
	if err != nil {
		t.Fatal(err)
	}

	return b
}

// HTTPCall form of request message which includes expected response
type HTTPCall struct {
	Header       map[string]string
	Method       string
	Status       int
	ExpectedResp []byte
}

// StartServer running mock http server
func StartServer(t *testing.T, reqs map[string]HTTPCall) (*httptest.Server, func()) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Check the method and the request URI. Eg: GET /some-path?query-param=query-value
		req, ok := reqs[fmt.Sprintf("%s %s", r.Method, r.RequestURI)]
		if !ok {
			t.Fatalf("expected response for request: %s is not found", fmt.Sprintf("%s %s", r.Method, r.RequestURI))
		}

		for k, v := range req.Header {
			w.Header().Set(k, v)
		}
		w.WriteHeader(req.Status)

		_, err := w.Write(req.ExpectedResp)
		if err != nil {
			t.Fatal(err)
		}
	}))

	return server, func() { server.Close() }
}

// RoundTripFunc is a func type to perform executing single HTTP transaction
// Mock http.Client by replacing http.Transport
// Source: http://hassansin.github.io/Unit-Testing-http-client-in-Go#2-by-replacing-httptransport
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip execute a single HTTP transaction
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewClient returns *http.Client with Transport replaced to avoid making real calls
func NewClient(t *testing.T, fn RoundTripFunc) *http.Client {
	t.Helper()

	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
