package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func RunAllSolutions() {
	fmt.Println("🎄 Auto-discovering implemented days...")

	results := DiscoverAndRunDays()

	if len(results) == 0 {
		fmt.Println("📂 No days found in solutions/ directory")
		fmt.Println("   Run: go run . generate 1")
		return
	}

	// Display console summary
	DisplayConsoleSummary(results)

	// Generate markdown summary
	if err := GenerateMarkdownSummary(results); err != nil {
		log.Printf("Error generating markdown summary: %v", err)
	} else {
		fmt.Println("\n📝 Markdown summary generated: README.md")
	}
}

func DiscoverAndRunDays() []DayResult {
	var results []DayResult
	solutionsDir := "solutions"

	// Check if solutions directory exists
	if _, err := os.Stat(solutionsDir); os.IsNotExist(err) {
		return results
	}

	// Scan for day directories
	entries, err := os.ReadDir(solutionsDir)
	if err != nil {
		log.Printf("Could not read solutions directory: %v", err)
		return results
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasPrefix(name, "day") {
			continue
		}

		// Extract day number
		dayStr := strings.TrimPrefix(name, "day")
		dayNum, err := strconv.Atoi(dayStr)
		if err != nil {
			continue
		}

		dayDir := filepath.Join(solutionsDir, name)

		// Check if both part files exist
		part1Path := filepath.Join(dayDir, "part1.go")
		part2Path := filepath.Join(dayDir, "part2.go")
		inputPath := filepath.Join(dayDir, "input.txt")
		if _, err := os.Stat(part1Path); os.IsNotExist(err) {
			continue
		}
		if _, err := os.Stat(part2Path); os.IsNotExist(err) {
			continue
		}

		fmt.Printf("🔍 Discovered day %02d\n", dayNum)

		result := DayResult{Day: dayNum}

		// Run Part 1
		start := time.Now()
		result.Part1 = runDayPart(dayDir, "part1", inputPath)
		result.Part1Time = time.Since(start)

		fmt.Println("oko", result.Part1)

		// Run Part 2
		start = time.Now()
		result.Part2 = runDayPart(dayDir, "part2", inputPath)
		result.Part2Time = time.Since(start)

		results = append(results, result)
	}

	// Sort by day number
	sort.Slice(results, func(i, j int) bool {
		return results[i].Day < results[j].Day
	})

	return results
}

func runDayPart(dayDir, part, inputPath string) string {
	partFile := part + ".go"

	// Check if part file exists
	fullPath := filepath.Join(dayDir, partFile)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return fmt.Sprintf("Error: %s not found", partFile)
	}

	// Check if input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return "No input file"
	}

	// Check if input file is empty
	if info, err := os.Stat(inputPath); err != nil {
		return fmt.Sprintf("Error: cannot stat input: %v", err)
	} else if info.Size() == 0 {
		return "Empty input"
	}

	// Get absolute path to day directory
	absDayDir, err := filepath.Abs(dayDir)
	if err != nil {
		return fmt.Sprintf("Error: cannot get absolute path: %v", err)
	}

	// Run the part using go run with proper working directory
	cmd := exec.Command("go", "run", partFile)
	cmd.Dir = absDayDir // Use absolute path

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error: %s", strings.TrimSpace(string(output)))
	}

	result := strings.TrimSpace(string(output))
	if result == "" {
		return "No output"
	}

	return result
}
