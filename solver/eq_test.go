package solver

import (
	"math"
	"reflect"
	"sort"
	"testing"
)

func Test_squareEquation_solve(t *testing.T) {
	tests := []struct {
		name    string
		se      *squareEquation
		want    []float64
		wantErr bool
	}{
		{
			name:    "no root",
			se:      NewSquareEquation(1, 0, 1),
			want:    []float64{},
			wantErr: false,
		},
		{
			name:    "two roots",
			se:      NewSquareEquation(1, 0, -1),
			want:    []float64{-1, 1},
			wantErr: false,
		},
		{
			name:    "two roots in desc order",
			se:      NewSquareEquation(1, 0, -1),
			want:    []float64{1, -1},
			wantErr: false,
		},
		{
			name:    "one root",
			se:      NewSquareEquation(1, -2, 1),
			want:    []float64{1},
			wantErr: false,
		},
		{
			name:    "one root with machine epsilon",
			se:      NewSquareEquation(1+eps, -2, 1-eps),
			want:    []float64{1 / (1 + eps)},
			wantErr: false,
		},
		{
			name:    "linear equation",
			se:      NewSquareEquation(1e-14, 2, -4),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "nan test",
			se:      NewSquareEquation(1, 2, math.Log(-1)),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "inf test",
			se:      NewSquareEquation(1, 2, math.Log(0)),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test with float answers",
			se:      NewSquareEquation(121, -121, 28),
			want:    []float64{4.0 / 11.0, 7.0 / 11.0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.se.solve()
			if (err != nil) != tt.wantErr {
				t.Errorf("squareEquation.solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Float64s(got)
			sort.Float64s(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("squareEquation.solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumberList(t *testing.T) {
	type args struct {
		list []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "numbers",
			args: args{list: []float64{1, 2, 3}},
			want: true,
		},
		{
			name: "inf",
			args: args{list: []float64{1, math.Inf(1), 3, 4, math.Inf(-1)}},
			want: false,
		},
		{
			name: "nan",
			args: args{list: []float64{math.NaN(), 2, 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumberList(tt.args.list...); got != tt.want {
				t.Errorf("isNumberList() = %v, want %v", got, tt.want)
			}
		})
	}
}
