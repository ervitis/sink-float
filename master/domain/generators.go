package domain

import (
	"math/rand"
	"time"
)

func GenerateRandNumber(min, max int) int {
	genRandNumber := func(min, max int) int {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		return rnd.Intn(max-min) + min
	}
	return genRandNumber(min, max)
}

func GenerateDirection() Direction {
	genDirection := func() Direction {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		return Direction(rnd.Intn(2-1) + 1)
	}
	return genDirection()
}
