package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
)

type Particle struct {
	position   Point
	Xv, Yv, Zv int
	Xa, Ya, Za int
}

type Point struct {
	X, Y, Z int
}

func main() {
	file, err := os.Open("20.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	particles := map[int]Particle{}
	particles2 := map[int]Particle{}
	particleNum := 0
	for scanner.Scan() {
		particles[particleNum] = parse(scanner.Text())
		particles2[particleNum] = parse(scanner.Text())
		particleNum++
	}

	println(solve(particles))
	println(solve2(particles2))
}

func solve2(particles map[int]Particle) (particleCount int) {
	i := 0
	for i < 5000 {
		positions := map[Point][]int{}
		for k, v := range particles {
			v.Xv += v.Xa
			v.Yv += v.Ya
			v.Zv += v.Za
			v.position.X += v.Xv
			v.position.Y += v.Yv
			v.position.Z += v.Zv
			particles[k] = v

			positions[v.position] = append(positions[v.position], k)
		}
		removeCollisions(particles, positions)
		i++
	}

	return len(particles)
}

func removeCollisions(particles map[int]Particle, positions map[Point][]int) {
	for _, ps := range positions {
		if len(ps) > 1 {
			for i := range ps {
				//fmt.Printf("DELETING: %+v AT POINT: %+v\n", ps, p)
				delete(particles, ps[i])
			}
		}
	}
}

// Part one
func solve(particles map[int]Particle) (particleNum int) {
	distances := map[int]int{}

	particleNum = -1
	i := 0
	for i < 5000 {
		for k, v := range particles {
			v.Xv += v.Xa
			v.Yv += v.Ya
			v.Zv += v.Za
			v.position.X += v.Xv
			v.position.Y += v.Yv
			v.position.Z += v.Zv
			particles[k] = v

			distances[k] = manhattanDistance(v.position)
		}
		particleNum = minDistance(distances)
		//fmt.Printf("least distance particleNum: %d\n", particleNum)
		i++
	}

	return particleNum
}

func minDistance(distances map[int]int) int {
	leastDistParticle := 0
	min := distances[0]
	for k, v := range distances {
		if v < min {
			min = v
			leastDistParticle = k
		}
	}
	return leastDistParticle
}

func manhattanDistance(p Point) int {
	return util.Abs(p.X) + util.Abs(p.Y) + util.Abs(p.Z)
}

func parse(s string) Particle {
	r := regexp.MustCompile(`<[^>]*>`)
	matches := r.FindAllString(s, -1)

	m0splitNstripped := strings.Split(util.Stripchars(matches[0], "<> "), ",")
	m1splitNstripped := strings.Split(util.Stripchars(matches[1], "<> "), ",")
	m2splitNstripped := strings.Split(util.Stripchars(matches[2], "<> "), ",")

	p := util.ConvStringArrayToIntArray(m0splitNstripped)
	v := util.ConvStringArrayToIntArray(m1splitNstripped)
	a := util.ConvStringArrayToIntArray(m2splitNstripped)

	return Particle{position: Point{X: p[0], Y: p[1], Z: p[2]},
		Xv: v[0], Yv: v[1], Zv: v[2],
		Xa: a[0], Ya: a[1], Za: a[2],
	}
}
