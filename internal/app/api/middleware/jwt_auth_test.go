package middleware

import "testing"

func Test_parseAuthHeader(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		wantType   string
		wantToken  string
		wantErr    bool
	}{
		{
			name:       "Valid Bearer Token",
			authHeader: "Bearer abcdef123456",
			wantType:   "Bearer",
			wantToken:  "abcdef123456",
			wantErr:    false,
		},
		{
			name:       "Empty Authorization Header",
			authHeader: "",
			wantType:   "",
			wantToken:  "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotToken, err := parseAuthHeader(tt.authHeader)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseAuthHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotType != tt.wantType {
				t.Errorf("parseAuthHeader() gotType = %v, want %v", gotType, tt.wantType)
			}

			if gotToken != tt.wantToken {
				t.Errorf("parseAuthHeader() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
