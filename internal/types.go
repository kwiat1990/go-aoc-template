package internal

import "time"

type DayResult struct {
	Day       int
	Part1     string
	Part2     string
	Part1Time time.Duration
	Part2Time time.Duration
}

type DaySolver interface {
	SolvePart1(input string) (string, error)
	SolvePart2(input string) (string, error)
}

type DayData struct {
	Day       int
	DayPadded string
	Year      int
	Package   string
	Title     string
	Date      string
}
