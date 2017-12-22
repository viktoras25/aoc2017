package main

import "fmt"
import "strings"
import "sort"
import "viktoras.de/aoc2017/utils"

type vector3 struct {
	x, y, z int
}

type particle struct {
	position     vector3
	velocity     vector3
	acceleration vector3
}

type particles []particle

func (p particles) Len() int      { return len(p) }
func (p particles) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p particles) Less(i, j int) bool {
	if p[i].position.x != p[j].position.x {
		return p[i].position.x < p[j].position.x
	}

	if p[i].position.y != p[j].position.y {
		return p[i].position.y < p[j].position.y
	}

	if p[i].position.z != p[j].position.z {
		return p[i].position.z < p[j].position.z
	}

	return true
}

func NewVector(str string) vector3 {
	str = strings.Trim(str, "<>")
	values := utils.StringsToInts(strings.Split(str, ","))
	return vector3{values[0], values[1], values[2]}
}

func NewParticle(str string) particle {
	parts := strings.Split(str, ", ")
	return particle{
		NewVector(parts[0][2:]),
		NewVector(parts[1][2:]),
		NewVector(parts[2][2:]),
	}
}

func (p *particle) Move(steps int) {
	p.position.x += p.velocity.x*steps + (steps*p.acceleration.x*(1+steps))/2
	p.position.y += p.velocity.y*steps + (steps*p.acceleration.y*(1+steps))/2
	p.position.z += p.velocity.z*steps + (steps*p.acceleration.z*(1+steps))/2

	p.velocity.x += steps * p.acceleration.x
	p.velocity.y += steps * p.acceleration.y
	p.velocity.z += steps * p.acceleration.z
}

func (p *particle) Distance() int {
	return abs(p.position.x) + abs(p.position.y) + abs(p.position.z)
}

func (p *particle) DistanceAfter(steps int) int {
	return abs(p.position.x+p.velocity.x*steps+(steps*p.acceleration.x*(1+steps))/2) +
		abs(p.position.y+p.velocity.y*steps+(steps*p.acceleration.y*(1+steps))/2) +
		abs(p.position.z+p.velocity.z*steps+(steps*p.acceleration.z*(1+steps))/2)
}

func (p *particle) Collide(s particle) bool {
	return p.position.x == s.position.x && p.position.y == s.position.y && p.position.z == s.position.z
}

func main() {

	particles := particles{}
	distances := []int{}

	for _, line := range utils.FileGetLines("input20.txt") {
		newParticle := NewParticle(line)
		particles = append(particles, newParticle)
		distances = append(distances, newParticle.DistanceAfter(1000))
	}

	for i := 0; i < 100; i++ {
		sort.Stable(particles)

		for j := 0; j < len(particles)-1; j++ {
			startedAt := j
			for j < len(particles)-1 && particles[j].Collide(particles[j+1]) {
				j++
			}

			// We have some collisions - remove and restart
			if startedAt != j {
				particles = append(particles[:startedAt], particles[j+1:]...)
				j = 0
			}
		}

		// Move a step
		for i := 0; i < len(particles); i++ {
			particles[i].Move(1)
		}
	}

	fmt.Println("Part 1", utils.IndexOf(distances, utils.MinInt(distances)))
	fmt.Println("Part 2", len(particles))
}

func abs(i int) int {
	if i > 0 {
		return i
	}

	return -i
}
