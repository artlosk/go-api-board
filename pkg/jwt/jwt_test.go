package jwt

import (
	"testing"

	"github.com/google/uuid"
)

func TestJwt_SigningString(t *testing.T) {
	type fields struct {
		secret string
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "#1",
			fields: fields{
				secret: "qwerty",
			},
			args: args{
				id: uuid.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := Jwt{
				secret: tt.fields.secret,
			}
			got, err := j.SigningString(tt.args.id)
			t.Log(got, err)

		})
	}
}
