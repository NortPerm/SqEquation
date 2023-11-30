package battle

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestRotateCommand_Execute(t *testing.T) {
	tests := []struct {
		name        string
		r           Rotable
		want        int
		wantErr     bool
		expectedErr error
	}{
		{
			name:    "all is good",
			r:       NewMockRotable(21, 24, 10, canReadDirection|canReadAngularVelocity|canWritePosition),
			want:    7, // 33 === 7 (mod 24)
			wantErr: false,
		},
		{
			name:        "can not read direction",
			r:           NewMockRotable(21, 24, 10, canReadAngularVelocity|canWritePosition),
			want:        0, // чтение невозможно вернется зироВэлью
			wantErr:     true,
			expectedErr: ErrInvalidDirection,
		},
		{
			name:        "can not read velocity",
			r:           NewMockRotable(21, 24, 10, canReadDirection|canWritePosition),
			want:        21, // остался там же
			wantErr:     true,
			expectedErr: ErrInvalidAngularVelocity,
		},
		{
			name:        "negative directions",
			r:           NewMockRotable(21, -5, 10, canReadDirection|canReadAngularVelocity|canWritePosition),
			want:        21, // остался там же
			wantErr:     true,
			expectedErr: ErrInvalidDirectionsCount,
		},
		{
			name:        "can not write direction",
			r:           NewMockRotable(21, 24, 10, canReadDirection|canReadAngularVelocity),
			want:        21, // остался там же
			wantErr:     true,
			expectedErr: ErrInvalidWriteDirection,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := NewRotateCommand(tt.r)
			err := rc.Execute()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("RotateCommand.Execute() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("RotateCommand.Execute() error = %v, wantErr %v", err, tt.expectedErr)
				}

			}

			p, _ := tt.r.GetDirection()
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("RotateCommand.Execute() = %v, want %v", p, tt.want)
			}
		})
	}
}
