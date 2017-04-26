package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var expected = "World"
	var result = hello()
	if result != expected {
		t.Error(
			"expected", expected,
			"got", result,
		)
	}
}
func TestGetProduto(t *testing.T) {
	var expected = Produto{"70101000", "SC", 0, "Ampolas de vidro,p/transporte ou embalagem", 13.45, 17.00, 18.02, 0.00, "0", "01/01/2017", "30/06/2017", "W7m9E1", "17.1.A", "IBPT"}
	var result, _ = getProduto("70101000", "SC", "0")
	if result != expected {
		t.Error(
			"expected", expected,
			"got", result,
		)

	}
}

func TestGetProdutoNotFound(t *testing.T) {
	expected := "Produto not found"
	var _, err = getProduto("7010110", "SC", "0")
	if err.Error() != expected {
		t.Error(
			"expected", expected,
			"got", err.Error(),
		)
	}
}

func TestWithoutParameters(t *testing.T) {
	req, _ := http.NewRequest("POST", "/", nil)
	rr := httptest.NewRecorder()

	HandleIndex(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestNotFound(t *testing.T) {
	data := url.Values{}
	data.Set("codigo", "invalid")
	data.Set("uf", "SC")
	data.Set("ex", "0")

	req, _ := http.NewRequest("POST", "/",
		strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length",
		strconv.Itoa(len(data.Encode())))
	rr := httptest.NewRecorder()

	HandleIndex(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestFound(t *testing.T) {
	data := url.Values{}
	data.Set("codigo", "70101000")
	data.Set("uf", "SC")
	data.Set("ex", "0")

	req, _ := http.NewRequest("POST", "/",strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length",strconv.Itoa(len(data.Encode())))

	rr := httptest.NewRecorder()

	HandleIndex(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
