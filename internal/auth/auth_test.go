package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthEmpty(t *testing.T) {
	got, _ := GetAPIKey(nil)

	if !reflect.DeepEqual("", got) {
		t.Fatalf("expected: <nil>, got: %v\n", got)
	}

	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 12345678")

	got, err := GetAPIKey(headers)
	if reflect.DeepEqual("", got) {
		t.Fatalf("expected: <nil>, got: %v, err: %v\n", got, err)
	}
}

func TestAuthApi(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 12345678")

	want := "12345678"
	got, _ := GetAPIKey(headers)

	if reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v\n", want, got)
	}
}
