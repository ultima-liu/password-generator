package test

import (
	"math/rand"
	"time"
)

func RandNumber(max int) int{
	rand.Seed(time.Now().UnixNano());
	i := rand.Int()
	return i%max;
}
