package battle

import (
	"math"
	"reflect"
	"testing"
)

func TestVector_Add(t *testing.T) {
	type args struct {
		v1 *Vector
	}
	tests := []struct {
		name    string
		v       *Vector
		args    args
		want    *Vector
		wantErr bool
	}{
		{
			name:    "simple adding",
			v:       NewVector(3, 4),
			args:    args{v1: NewVector(5, 8)},
			want:    NewVector(8, 12),
			wantErr: false,
		},
		{
			name:    "invalid case",
			v:       NewVector(3, 4),
			args:    args{v1: NewVector(5, math.Log(0))},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Add(tt.args.v1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
