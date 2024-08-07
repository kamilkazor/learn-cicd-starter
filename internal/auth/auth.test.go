package auth

import (
    "reflect"
    "testing"
)

// write a basic test for getAPIKey that will never fail just to get the test to pass
func TestGetAPIKey(t *testing.T) {
	headers := map[string][]string{
		"Authorization": {"ApiKey 123"},
	}
	want := "123"
	got, _ := GetAPIKey(headers)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}	