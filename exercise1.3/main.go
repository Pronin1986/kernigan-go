package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println()
	fmt.Println("******************************")
	fmt.Println("COMPARE STRINGS CONCAT METHODS")
	fmt.Println("******************************")
	fmt.Println()

	fmt.Println("GENERATE STRINGS")
	t0 := time.Now()
	strSlice := randStrings(10)
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	fmt.Println()

	t0 = time.Now()
	fmt.Println("SIMPLE CONCAT: ", SimpleConcat(strSlice))
	t1 = time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	fmt.Println()

	t0 = time.Now()
	fmt.Println("STRING JOIN CONCAT: ", StringJoinConcat(strSlice))
	t1 = time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	fmt.Println()

	t0 = time.Now()
	fmt.Println("STRING BUILDER CONCAT: ", StringBuilderConcat(strSlice))
	t1 = time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	fmt.Println()

	t0 = time.Now()
	fmt.Println("STRING BUILDER CONCAT 2: ", StringBuilderWithAllocConcat(strSlice))
	t1 = time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	fmt.Println()
}

// randStrings generate slice of randome strings
func randStrings(cnt int64) []string {
	result := []string{}
	var minL int64 = 10
	var maxL int64 = 100
	var i int64
	rand.Seed(time.Now().UnixNano())

	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for i = 0; i < cnt; i++ {
		strL := rand.Int63n(maxL-minL) + minL
		b := make([]byte, strL)
		for i := range b {
			b[i] = letters[rand.Int63()%int64(len(letters))]
		}
		result = append(result, string(b))
	}

	return result
}
