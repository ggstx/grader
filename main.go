package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var scores []float64
	fmt.Print("Student Name: ")
	var studentName string
	fmt.Scan(&studentName)

	for range 6 {
		sc := getScore()
		scores = append(scores, sc)
	}

	// instance of Student struct.
	var student1 = Student{
		Scores: map[string]float64{
			"Math & Statistics": scores[0],
			"English":           scores[1],
			"Geography":         scores[2],
			"Systems Design":    scores[3],
			"General Studies":   scores[4],
			"Hardware Design":   scores[5],
		},
		Name:  studentName,
		Major: "Computer Science",
	}
	
	keys := getKeys(student1.Scores)
	var ps, fs int

	fmt.Printf("\nDear %s,\nYour performance across the assessed subjects in %s department is summarized below:\n", student1.Name, student1.Major)
	fmt.Printf("\nPassed Subjects:\n")
	for _, k := range keys {
		if isPass(student1.Scores[k]) {
			fmt.Printf("%s: %.1f\n", k, student1.Scores[k])
			ps++
		}
	}
	fmt.Printf("\nFailed Subjects:\n")
	for _, k := range keys {
		if !isPass(student1.Scores[k]) {
			fmt.Printf("%s: %.1f\n", k, student1.Scores[k])
			fs++
		}
	}
	fmt.Printf("\nTotal passed: %d\nTotal failed: %d\n", ps, fs)
}

type Student struct {
	Scores map[string]float64
	Name   string
	Major  string
}

func getScore() float64 {
	time.Now().UnixNano()
	score := 30 + rand.Float64()*50   // 30 ≤ value < 80
	score = math.Round(score*10) / 10 // 1 decimal place
	return score
}
func getKeys(m map[string]float64) []string { 
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
func isPass(sc float64) bool {
	if sc >= 50.0 {
		return true
	} else {
		return false
	}
}
