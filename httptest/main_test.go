package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTextHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	textHandler(rec, req)
	wantCode, wantRes := http.StatusOK, text
	gotCode, gotRes := rec.Code, rec.Body.String()

	if wantCode != gotCode {
		t.Errorf("Ожидается код = %d, получен код = %d", wantCode, gotCode)
	}
	if wantRes != gotRes {
		t.Errorf("Ожидается = %q, получено = %q", wantRes, gotRes)
	}
}

func TestHtmlHandler(t *testing.T) {
	handler := http.HandlerFunc(htmlHandler)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/html", nil)

	handler.ServeHTTP(rec, req)
	wantCode, wantRes := http.StatusOK, []byte(html)
	gotCode, gotRes := rec.Code, rec.Body.Bytes()

	if wantCode != gotCode {
		t.Errorf("Ожидается код = %d, получен код = %d", wantCode, gotCode)
	}
	if !reflect.DeepEqual(wantRes, gotRes) {
		t.Errorf("Ожидается: %v Получено: %v", string(wantRes), string(gotRes))
	}
}

func TestTextClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(textHandler))
	defer srv.Close()

	wantCode, wantRes := http.StatusOK, text
	gotCode, gotRes := Client(srv.URL)

	if wantCode != gotCode {
		t.Errorf("Ожидается код = %d, получен код = %d", wantCode, gotCode)
	}
	if !reflect.DeepEqual(wantRes, gotRes) {
		t.Errorf("Ожидается: %v Получено: %v", string(wantRes), string(gotRes))
	}
}

func TestHtmlClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(htmlHandler))
	defer srv.Close()

	wantCode, wantRes := http.StatusOK, html
	gotCode, gotRes := Client(srv.URL)

	if wantCode != gotCode {
		t.Errorf("Ожидается код = %d, получен код = %d", wantCode, gotCode)
	}
	if !reflect.DeepEqual(wantRes, gotRes) {
		t.Errorf("Ожидается: %v Получено: %v", string(wantRes), string(gotRes))
	}
}
