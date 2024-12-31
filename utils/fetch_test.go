package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetch_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	}))
	defer server.Close()

	body, err := Fetch(server.URL)

	require.NoError(t, err)
	assert.Equal(t, []byte("Hello, world!"), body)
}

func TestFetch_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	body, err := Fetch(server.URL)

	require.Error(t, err)
	assert.Nil(t, body)
	assert.Contains(t, err.Error(), "unexpected status code: 500")
}

func TestFetch_InvalidURL(t *testing.T) {
	body, err := Fetch("http://invalid-url")

	require.Error(t, err)
	assert.Nil(t, body)
	assert.Contains(t, err.Error(), "failed to fetch URL")
}
