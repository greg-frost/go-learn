package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTextResponse(t *testing.T) {
	handler := http.HandlerFunc(textResponse)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	handler.ServeHTTP(rec, req)

	wantCode, wantRes := http.StatusOK, text
	gotCode, gotRes := rec.Code, rec.Body.String()

	if wantCode != gotCode {
		t.Errorf("Want code = %d, got code = %d", wantCode, gotCode)
	}
	if wantRes != gotRes {
		t.Errorf("Want body = %q, got body = %q", wantRes, gotRes)
	}
}

func TestHtmlResponse(t *testing.T) {
	handler := http.HandlerFunc(htmlResponse)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/html", nil)

	handler.ServeHTTP(rec, req)

	wantCode, wantRes := http.StatusOK, []byte(html)
	gotCode, gotRes := rec.Code, rec.Body.Bytes()

	if wantCode != gotCode {
		t.Errorf("Want code = %d, got code = %d", wantCode, gotCode)
	}
	if !reflect.DeepEqual(wantRes, gotRes) {
		t.Errorf("Want: %v Got: %v", string(wantRes), string(gotRes))
	}
}
