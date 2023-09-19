package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandleForm(t *testing.T) {
	w := httptest.NewRecorder()

	formBody := url.Values{
		"say": []string{"hi"},
		"to":  []string{"rheo"},
	}
	dataReader := formBody.Encode()

	req := httptest.NewRequest(
		http.MethodPost,
		"http://localhost",
		strings.NewReader(dataReader),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	handleForm(w, req)

	responseBody, err := io.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("error reading response body", err)
	}

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected %v, got: %v\n", http.StatusTemporaryRedirect, w.Result().StatusCode)
	}

	expectedResponseBody := "HI RHEO"
	if expectedResponseBody != string(responseBody) {
		t.Fatalf("Expected response: %s, Got:%s", expectedResponseBody, string(responseBody))
	}

}
