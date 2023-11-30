package battle

import (
	"math"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestMoveCommand_Execute(t *testing.T) {
	tests := []struct {
		name        string
		m           Movable
		want        *Vector
		wantErr     bool
		expectedErr error
	}{
		{
			name:    "all is good",
			m:       NewMockMovable(12, 5, -7, 3, canReadPosition|canReadVelocity|canWritePosition),
			want:    NewVector(5, 8),
			wantErr: false,
		},
		{
			name:        "can not read position",
			m:           NewMockMovable(12, 5, -7, 3, canReadVelocity|canWritePosition),
			want:        nil, // чтение невозможно
			wantErr:     true,
			expectedErr: ErrInvalidPosition,
		},
		{
			name:        "can not read velocity",
			m:           NewMockMovable(12, 5, -7, 3, canReadPosition|canWritePosition),
			want:        NewVector(12, 5), // остался там же
			wantErr:     true,
			expectedErr: ErrInvalidVelocity,
		},
		{
			name:        "invalid velocity",
			m:           NewMockMovable(12, 5, math.NaN(), 3, canReadPosition|canReadVelocity|canWritePosition),
			want:        NewVector(12, 5), // остался там же
			wantErr:     true,
			expectedErr: ErrInvalidVelocity,
		},
		{
			name:        "can not write position",
			m:           NewMockMovable(12, 5, -7, 3, canReadPosition|canReadVelocity),
			want:        NewVector(12, 5), // остался там же
			wantErr:     true,
			expectedErr: ErrInvalidWritePosition,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := NewMoveCommand(tt.m)
			err := mc.Execute()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("MoveCommand.Execute() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("MoveCommand.Execute() error = %v, wantErr %v", err, tt.expectedErr)
				}

			}

			p, _ := tt.m.GetPosition()
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("MoveCommand.Execute() = %v, want %v", p, tt.want)
			}
		})
	}
}
