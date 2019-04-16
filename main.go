package main

import (
	crand "crypto/rand"

	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/chase0213/reinforcement-learning-go/pkg/bandit"
)

func main() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	ms := []float64{
		0.1, 0.2, 0.3,
	}
	eps := 0.1
	upperLimit := 10.0
	N := 10000

	fmt.Printf("\n=== Epsilon Greedy Algorithm ===\n")
	egbs, _ := bandit.RunEpsilonGreedy(ms, eps, N)
	for i, b := range egbs {
		fmt.Printf("True mean of bandit(No. %d) = %f, actual = %f\n", i+1, b.M(), b.Mean())
	}

	fmt.Printf("\n=== Optimistic Initial Value Algorithm ===\n")
	oivbs, _ := bandit.RunOptimisticInitialValue(ms, upperLimit, N)
	for i, b := range oivbs {
		fmt.Printf("True mean of bandit(No. %d) = %f, actual = %f\n", i+1, b.M(), b.Mean())
	}

}