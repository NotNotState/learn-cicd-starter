package auth

import (
	"net/http"
	"testing"
)

func CreateCase(key, value string) http.Header {
	return http.Header{
		"Authorization": []string{key + " " + value},
	}
}

func TestAuth(t *testing.T) {
	header := CreateCase("ApiKey", "1234567")
	key, err := GetAPIKey(header)

	if err != nil {
		t.Fatal(err)
	}
	if key != "1234567" {
		t.Fatal("invalid api key")
	}
	header = CreateCase("", "1234567")

	key, err = GetAPIKey(header)
	if err == nil {
		t.Fatal("faild to raise error")
	}

	if key != "" {
		t.Fatal("failed to return nil key")
	}

}
