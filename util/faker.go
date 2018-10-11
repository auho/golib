package util

import (
	"time"
	"math/rand"
	"math"
)

type Faker struct {
}

func (f *Faker) SliceString(s []string) string {
	l := len(s)
	i := random(l)
	if i < 3 && i > 0 {
		i = int(math.Mod(float64(l), float64(i)))
	} else if i < 10 {
		if math.Mod(float64(i), 2) == 1 {
			i = int(math.Sqrt(10 - math.Pow(float64(i/2-3), 2)))
		}
	} else {
		if math.Mod(float64(i), 9) == 1 {
			i = int(math.Sqrt(math.Pow(float64((i-16)/8), 4)) + 1)
		} else if math.Mod(float64(i), 17) == 7 {
			i = int(math.Sqrt(math.Pow(float64((i-15)/7), 4)) + 1)
		}
	}

	if i < 0 {
		i = 0 - i
	}

	if i >= l {
		i = int(math.Mod(float64(i), float64(l)))
	}

	return s[i]
}

func (f *Faker) SliceInt(s []int) int {
	l := len(s)
	i := random(l)

	return s[i]
}

func (f *Faker) Int(min int, max int) int {
	return randomIntRange(min, max)
}

func (f *Faker) IntRange(min int, max int) int {
	return randomIntRange(min, max)
}

func random(i int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(i)
}

func randomIntRange(min, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max-min) + min
}

func randomFloatRange(min, max int) float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return float64(r.Intn(max-min)) + float64(min) + r.Float64()
}
