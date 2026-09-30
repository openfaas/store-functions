package function

import (
	"encoding/json"
	"regexp"
	"testing"
)

func TestHandleReturnsCorrectResponse(t *testing.T) {
	expected := "www.google.com"
	resp := handle([]byte("www.google.com/about/"), "")

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}

func TestHandleReturnsMultiSanResponse(t *testing.T) {
	expected := ".stefanprodan.com"
	resp := handle([]byte("stefanprodan.com"), "")

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}

func TestHandleReturnsJSONResponse(t *testing.T) {
	resp := handle([]byte("www.google.com/about/"), "output=json")

	var out struct {
		Host       string `json:"Host"`
		CommonName string `json:"CommonName"`
	}
	if err := json.Unmarshal([]byte(resp), &out); err != nil {
		t.Fatalf("expected JSON response, got: %v (err: %s)", resp, err)
	}

	if out.Host == "" {
		t.Fatalf("expected a non-empty host in JSON response, got: %v", resp)
	}
}
