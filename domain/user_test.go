package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_ID(t *testing.T) {
	type fields struct {
		id       string
		username string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should return name",
			fields: fields{
				id:       "id",
				username: "username",
			},
			want: "id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id:       tt.fields.id,
				username: tt.fields.username,
			}
			if got := u.ID(); got != tt.want {
				t.Errorf("User.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Username(t *testing.T) {
	type fields struct {
		id       string
		username string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should return username",
			fields: fields{
				id:       "id",
				username: "username",
			},
			want: "username",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id:       tt.fields.id,
				username: tt.fields.username,
			}
			if got := u.Username(); got != tt.want {
				t.Errorf("User.Username() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetUsername(t *testing.T) {
	type fields struct {
		id       string
		username string
	}
	type args struct {
		un string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "should change username correctly",
			fields: fields{
				id:       "id",
				username: "username",
			},
			args: args{
				un: "newUsername",
			},
			want: "newUsername",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id:       tt.fields.id,
				username: tt.fields.username,
			}
			u.SetUsername(tt.args.un)
			assert.Equal(t, tt.want, u.username)
		})
	}
}

// type mockUUID struct{}

// func (u *mockUUID) V4() string {
// 	return "test"
// }
// func TestGetUUID(t *testing.T) {
// 	type args struct {
// 		u uuid.UUIDGenerator
// 	}
// 	test := struct {
// 		name string
// 		args args
// 		want interface{}
// 	}{
// 		name: "should return uuid",
// 	}
// 	t.Run(test.name, func(t *testing.T) {
// 		got := GetUUID(&mockUUID{})
// 		if diff := cmp.Diff("test", got); diff != "" {
// 			t.Errorf("GetUUID() = %v", got)
// 		}
// 	})
// }
