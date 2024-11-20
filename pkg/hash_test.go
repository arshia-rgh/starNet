package pkg

import "testing"

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
			name:    "ValidPassword",
			args:    args{password: "password123"},
			wantErr: false,
		},
		{
			name:    "EmptyPassword",
			args:    args{password: ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
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
		name string
		args args
		want bool
	}{
		{
			name: "ValidPassword",
			args: args{password: "password123", hash: func() string { hashedPassword, _ := HashPassword("password123"); return hashedPassword }()},
			want: true,
		},
		{
			name: "InvalidPassword",
			args: args{password: "password123", hash: func() string { hashedPassword, _ := HashPassword("another password "); return hashedPassword }()},
			want: false,
		},
		{
			name: "EmptyPassword",
			args: args{password: "", hash: func() string { hashedPassword, _ := HashPassword("another password "); return hashedPassword }()},
			want: false,
		},
		{
			name: "EmptyHash",
			args: args{password: "password123", hash: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPassword(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
