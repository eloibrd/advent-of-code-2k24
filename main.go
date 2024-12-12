package main

import (
	day1 "advent-of-code-2k24/day-1"
	"fmt"
	"log/slog"
)

func main() {
	// day 1
	day1TotalDistance, err := day1.CalculateTotalDistance()
	if err != nil {
		slog.Error(fmt.Sprintf("Day 1 total distances generated an error : %v", err))
	} else {
		slog.Info(fmt.Sprintf("Day 1 total distance : %d", day1TotalDistance))
	}
	day1SimilarityScore, err := day1.CalculateSimilarityScore()
	if err != nil {
		slog.Error(fmt.Sprintf("Day 1 similarity score an error : %v", err))
	} else {
		slog.Info(fmt.Sprintf("Day 1 similarity score : %d", day1SimilarityScore))
	}
}
