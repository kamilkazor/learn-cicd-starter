package auth

import (
    "reflect"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name     string
		headers  map[string][]string
		expected string
		err      error
	}{
		{
			name:     "no auth header",
			headers:  map[string][]string{},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		{
			name:     "malformed auth header",
			headers:  map[string][]string{"Authorization": {"Bearer"}},
			expected: "",
			err:      ErrMalformedAuthHeader,
		},
		{
			name:     "valid auth header",
			headers:  map[string][]string{"Authorization": {"ApiKey 123"}},
			expected: "123",
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			headers := make(map[string][]string)
			for k, v := range tc.headers {
				headers[k] = v
			}

			got, err := GetAPIKey(headers)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
			if got != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, got)
			}
		})
	}
}