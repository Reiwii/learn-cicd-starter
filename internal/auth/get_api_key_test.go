package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name    string
		input   http.Header
		want    string
		wantErr string
	}{
		{
			name:    "no header",
			input:   http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:    "malformed header",
			input:   http.Header{"Authorization": []string{"noapi cos"}},
			want:    "",
			wantErr: "malformed authorization header",
		},
		{
			name:    "success",
			input:   http.Header{"Authorization": []string{"ApiKey cos"}},
			want:    "cos",
			wantErr: "",
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error '%v', got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Errorf("expected error message '%v', got '%v'", tc.wantErr, err.Error())
				}
				return
			}
			if tc.wantErr == "" {
				if err != nil {
					t.Fatal("did not expect error")
				}
				if got != tc.want {
					t.Errorf("got %v, want %v", got, tc.want)
				}
			}
		})
	}

}
