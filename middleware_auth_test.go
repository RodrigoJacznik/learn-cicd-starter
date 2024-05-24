package main

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetApiKey_EmptyHeader(t *testing.T) {
    headers := http.Header{}
    _, err := auth.GetAPIKey(headers)

    if !errors.Is(err, auth.ErrNoAuthHeaderIncluded) {
        t.Fatalf("expected: %v, got %v", auth.ErrNoAuthHeaderIncluded, err)
    }
}

func TestGetApiKey_InvalidHeader(t *testing.T) {
    headers := make(http.Header)
    headers["Authorization"] = []string{"foo"}

    _, err := auth.GetAPIKey(headers)

    if !errors.Is(err, auth.ErrMalformedAuthHeader) {
        t.Fatalf("expected: %v, got %v", auth.ErrMalformedAuthHeader, err)
    }
}

func TestGetApiKey_ValidHeader(t *testing.T) {
    headers := make(http.Header)
    expected := "foo"
    headers["Authorization"] = []string{"ApiKey " + expected}

    key, err := auth.GetAPIKey(headers)

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if !reflect.DeepEqual(key, expected) {
        t.Fatalf("expected %v, got %v", expected, key)
    }
}
