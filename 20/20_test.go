package main

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Particle
	}{
		{"example-simple", args{"p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>"}, Particle{Point{3, 0, 0}, 2, 0, 0, -1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	type args struct {
		particles map[int]Particle
	}
	tests := []struct {
		name            string
		args            args
		wantParticleNum int
	}{
		{"example", args{map[int]Particle{0: Particle{Point{3, 0, 0}, 2, 0, 0, -1, 0, 0},
			1: Particle{Point{4, 0, 0}, 0, 0, 0, -2, 0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParticleNum := solve(tt.args.particles); gotParticleNum != tt.wantParticleNum {
				t.Errorf("solve() = %v, want %v", gotParticleNum, tt.wantParticleNum)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		particles map[int]Particle
	}
	tests := []struct {
		name              string
		args              args
		wantParticleCount int
	}{
		{"example", args{map[int]Particle{0: Particle{Point{-6, 0, 0}, 3, 0, 0, 0, 0, 0},
			1: Particle{Point{-4, 0, 0}, 2, 0, 0, 0, 0, 0},
			2: Particle{Point{-2, 0, 0}, 1, 0, 0, 0, 0, 0},
			3: Particle{Point{3, 0, 0}, -1, 0, 0, 0, 0, 0},
		}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParticleCount := solve2(tt.args.particles); gotParticleCount != tt.wantParticleCount {
				t.Errorf("solve2() = %v, want %v", gotParticleCount, tt.wantParticleCount)
			}
		})
	}
}
