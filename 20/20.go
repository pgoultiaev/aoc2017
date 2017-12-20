package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/pgoultiaev/aoc2017/util"
)

type Particle struct {
	position    Point
	Xv, Yv, Zv  int
	Xa, Ya, Za  int
	distanceTo0 int
	id          int
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

	particles2 := map[int]Particle{}
	particles := []Particle{}
	particleID := 0
	for scanner.Scan() {
		particles2[particleID] = parse(scanner.Text(), particleID)
		particles = append(particles, parse(scanner.Text(), particleID))
		particleID++
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
func solve(particles []Particle) (particleNum int) {
	sortParticles(particles)
	j := 0
	for {
		for i, p := range particles {
			p.Xv += p.Xa
			p.Yv += p.Ya
			p.Zv += p.Za
			p.position.X += p.Xv
			p.position.Y += p.Yv
			p.position.Z += p.Zv
			p.distanceTo0 = manhattanDistance(p.position)

			particles[i] = p
		}

		newOrderedParticles := make([]Particle, len(particles))
		copy(newOrderedParticles, particles)
		sortParticles(newOrderedParticles)

		// curious how to mathematically solve this
		if j > 10 && reflect.DeepEqual(particles, newOrderedParticles) {
			break
		}
		particles = newOrderedParticles

		j++
	}

	return particles[0].id
}

func sortParticles(ps []Particle) {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].distanceTo0 < ps[j].distanceTo0
	})
}

func manhattanDistance(p Point) int {
	return util.Abs(p.X) + util.Abs(p.Y) + util.Abs(p.Z)
}

func parse(s string, id int) Particle {
	r := regexp.MustCompile(`<[^>]*>`)
	matches := r.FindAllString(s, -1)

	m0splitNstripped := strings.Split(util.Stripchars(matches[0], "<> "), ",")
	m1splitNstripped := strings.Split(util.Stripchars(matches[1], "<> "), ",")
	m2splitNstripped := strings.Split(util.Stripchars(matches[2], "<> "), ",")

	p := util.ConvStringArrayToIntArray(m0splitNstripped)
	v := util.ConvStringArrayToIntArray(m1splitNstripped)
	a := util.ConvStringArrayToIntArray(m2splitNstripped)

	pos := Point{X: p[0], Y: p[1], Z: p[2]}

	return Particle{
		position: pos,
		Xv:       v[0], Yv: v[1], Zv: v[2],
		Xa: a[0], Ya: a[1], Za: a[2],
		distanceTo0: manhattanDistance(pos),
		id:          id,
	}
}
