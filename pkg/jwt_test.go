package pkg

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		userID    int
		userRole  string
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid token generation",
			args: args{
				userID:    1,
				userRole:  "admin",
				secretKey: "secret",
			},
			wantErr: false,
		},
		{
			name: "Empty secret key",
			args: args{
				userID:    1,
				userRole:  "admin",
				secretKey: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.userID, tt.args.userRole, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("GenerateToken() got an empty token")
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	type args struct {
		token     string
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantErr bool
	}{
		{
			name: "Valid token verification",
			args: args{
				token:     func() string { token, _ := GenerateToken(1, "admin", "secret"); return token }(),
				secretKey: "secret",
			},
			want:    1,
			want1:   "admin",
			wantErr: false,
		},
		{
			name: "Invalid token",
			args: args{
				token:     "invalid.token.here",
				secretKey: "secret",
			},
			want:    0,
			want1:   "",
			wantErr: true,
		},
		{
			name: "Invalid secret key",
			args: args{
				token:     func() string { token, _ := GenerateToken(1, "admin", "secret"); return token }(),
				secretKey: "wrongsecret",
			},
			want:    0,
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := VerifyToken(tt.args.token, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("VerifyToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
