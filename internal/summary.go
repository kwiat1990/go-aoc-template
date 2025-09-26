package internal

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

func DisplayConsoleSummary(results []DayResult) {
	fmt.Println("\n🎄 Advent of Code Solutions Summary 🎄")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "Day\tPart 1\tTime\tPart 2\tTime")
	fmt.Fprintln(w, "---\t------\t----\t------\t----")

	totalTime := time.Duration(0)

	for _, result := range results {
		totalTime += result.Part1Time + result.Part2Time

		fmt.Fprintf(w, "%d\t%s\t%v\t%s\t%v\n",
			result.Day,
			truncate(result.Part1, 15),
			formatDuration(result.Part1Time),
			truncate(result.Part2, 15),
			formatDuration(result.Part2Time))
	}

	fmt.Fprintln(w, "---\t------\t----\t------\t----")
	fmt.Fprintf(w, "Total\t\t%v\t\t\n", formatDuration(totalTime))
	w.Flush()

	fmt.Println()
	fmt.Println("📁 Navigation:")
	for _, result := range results {
		fmt.Printf("Day %02d: solutions/day%02d/part1.go | solutions/day%02d/part2.go | solutions/day%02d/day%02d_test.go\n",
			result.Day, result.Day, result.Day, result.Day, result.Day)
	}
}

func GenerateMarkdownSummary(results []DayResult) error {
	var md strings.Builder

	md.WriteString("# 🎄 Advent of Code - Solutions Summary\n\n")
	md.WriteString(fmt.Sprintf("**Last updated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	// Overview stats
	totalTime := time.Duration(0)
	solvedDays := len(results)
	totalParts := 0

	for _, result := range results {
		totalTime += result.Part1Time + result.Part2Time
		if !strings.Contains(result.Part1, "Error") && !strings.Contains(result.Part1, "No") && !strings.Contains(result.Part1, "Empty") {
			totalParts++
		}
		if !strings.Contains(result.Part2, "Error") && !strings.Contains(result.Part2, "No") && !strings.Contains(result.Part2, "Empty") {
			totalParts++
		}
	}

	md.WriteString("## 📊 Overview\n\n")
	md.WriteString(fmt.Sprintf("- **Days completed:** %d/25\n", solvedDays))
	md.WriteString(fmt.Sprintf("- **Parts completed:** %d/50\n", totalParts))
	md.WriteString(fmt.Sprintf("- **Total execution time:** %s\n", formatDuration(totalTime)))
	if totalParts > 0 {
		md.WriteString(fmt.Sprintf("- **Average time per part:** %s\n\n", formatDuration(totalTime/time.Duration(totalParts))))
	}

	// Solutions table
	md.WriteString("## 🏆 Solutions\n\n")
	md.WriteString("| Day | Part 1 | Time | Part 2 | Time | Files |\n")
	md.WriteString("|-----|--------|------|--------|------|-------|\n")

	for _, result := range results {
		status1 := getStatusIcon(result.Part1)
		status2 := getStatusIcon(result.Part2)

		md.WriteString(fmt.Sprintf("| %02d | %s `%s` | %s | %s `%s` | %s | [Part1](solutions/day%02d/part1.go) \\| [Part2](solutions/day%02d/part2.go) \\| [Tests](solutions/day%02d/day%02d_test.go) |\n",
			result.Day,
			status1, truncateForMD(result.Part1, 15),
			formatDuration(result.Part1Time),
			status2, truncateForMD(result.Part2, 15),
			formatDuration(result.Part2Time),
			result.Day, result.Day, result.Day, result.Day))
	}

	// Add placeholder rows for remaining days
	for dayNum := len(results) + 1; dayNum <= 25; dayNum++ {
		md.WriteString(fmt.Sprintf("| %02d | ❌ `Not started` | - | ❌ `Not started` | - | - |\n", dayNum))
	}

	md.WriteString("\n## 🚀 Quick Start\n\n")
	md.WriteString("```bash\n")
	md.WriteString("# Generate a new day\n")
	md.WriteString("go run . generate 5\n\n")
	md.WriteString("# Run all solutions (auto-discovery)\n")
	md.WriteString("go run .\n\n")
	md.WriteString("# Test specific day\n")
	md.WriteString("go test ./solutions/day01\n\n")
	md.WriteString("# Run individual parts\n")
	md.WriteString("cd solutions/day01 && go run part1.go\n")
	md.WriteString("```\n\n")

	// Quick navigation
	md.WriteString("## 🧭 Navigation\n\n")
	for _, result := range results {
		md.WriteString(fmt.Sprintf("- **Day %02d**: [Part1](solutions/day%02d/part1.go) | [Part2](solutions/day%02d/part2.go) | [Tests](solutions/day%02d/day%02d_test.go) | [Input](solutions/day%02d/input.txt)\n",
			result.Day, result.Day, result.Day, result.Day, result.Day, result.Day))
	}

	return os.WriteFile("SOLUTIONS.md", []byte(md.String()), 0644)
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func truncateForMD(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%dns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.1fμs", float64(d.Nanoseconds())/1000)
	} else if d < time.Second {
		return fmt.Sprintf("%.1fms", float64(d.Nanoseconds())/1000000)
	}
	return fmt.Sprintf("%.2fs", d.Seconds())
}

func getStatusIcon(result string) string {
	if strings.Contains(result, "Error") || strings.Contains(result, "Failed") {
		return "❌"
	}
	if strings.Contains(result, "No") || strings.Contains(result, "Empty") || result == "0" {
		return "⏳"
	}
	return "✅"
}
