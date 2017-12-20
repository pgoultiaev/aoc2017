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
	Xpos, Ypos, Zpos int
	Xv, Yv, Zv       int
	Xa, Ya, Za       int
}

func main() {
	file, err := os.Open("20.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	particles := map[int]Particle{}
	particleNum := 0
	for scanner.Scan() {
		particles[particleNum] = parse(scanner.Text())
		particleNum++
	}

	println(solve(particles))
}

func solve(particles map[int]Particle) (particleNum int) {
	distances := map[int]int{}

	particleNum = -1
	i := 0
	for i < 5000 {
		for k, v := range particles {
			v.Xv += v.Xa
			v.Yv += v.Ya
			v.Zv += v.Za
			v.Xpos += v.Xv
			v.Ypos += v.Yv
			v.Zpos += v.Zv
			particles[k] = v

			distances[k] = manhattanDistance(v)
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

func manhattanDistance(p Particle) int {
	return util.Abs(p.Xpos) + util.Abs(p.Ypos) + util.Abs(p.Zpos)
}

func parse(s string) Particle {
	r := regexp.MustCompile(`<[^>]*>`)
	matches := r.FindAllString(s, -1)

	m0splitNstripped := strings.Split(util.Stripchars(matches[0], "<> "), ",")
	m1splitNstripped := strings.Split(util.Stripchars(matches[1], "<> "), ",")
	m2splitNstripped := strings.Split(util.Stripchars(matches[2], "<> "), ",")

	position := util.ConvStringArrayToIntArray(m0splitNstripped)
	velocity := util.ConvStringArrayToIntArray(m1splitNstripped)
	acceleration := util.ConvStringArrayToIntArray(m2splitNstripped)

	return Particle{Xpos: position[0], Ypos: position[1], Zpos: position[2],
		Xv: velocity[0], Yv: velocity[1], Zv: velocity[2],
		Xa: acceleration[0], Ya: acceleration[1], Za: acceleration[2],
	}
}
