package function

import (
	"regexp"
	"testing"
)

func TestHandleReturnsCorrectResponse(t *testing.T) {
	expected := "www.google.com"
	resp := handle([]byte("www.google.com/about/"))

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}

func TestHandleReturnsMultiSanResponse(t *testing.T) {
	expected := ".stefanprodan.com"
	resp := handle([]byte("stefanprodan.com"))

	r := regexp.MustCompile("(?m:" + expected + ")")
	if !r.MatchString(resp) {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
	}
}
