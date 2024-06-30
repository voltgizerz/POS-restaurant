package utils

import (
	"testing"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid Password",
			args: args{password: "mysecretpassword"},
			wantErr: false,
		},
		{
			name: "Empty Password",
			args: args{password: ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check if the hashed password can be verified with the original password
			if err := bcrypt.CompareHashAndPassword([]byte(got), []byte(tt.args.password)); err != nil {
				t.Errorf("Hashed password does not match the original password: %v", err)
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Correct Password",
			args: args{
				password: "mysecretpassword",
				hash: func() string {
					hash, _ := HashPassword("mysecretpassword")
					return hash
				}(),
			},
			wantErr: false,
		},
		{
			name: "Incorrect Password",
			args: args{
				password: "wrongpassword",
				hash: func() string {
					hash, _ := HashPassword("mysecretpassword")
					return hash
				}(),
			},
			wantErr: true,
		},
		{
			name: "Empty Password",
			args: args{
				password: "",
				hash: func() string {
					hash, _ := HashPassword("")
					return hash
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := VerifyPassword(tt.args.password, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
