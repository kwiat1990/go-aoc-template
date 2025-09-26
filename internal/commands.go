package internal

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func HandleCommands() {
	command := os.Args[1]

	switch command {
	case "generate", "gen", "new":
		handleGenerate()
	case "help", "-h", "--help":
		PrintHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		PrintHelp()
		os.Exit(1)
	}
}

func handleGenerate() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . generate <day> [year]")
		fmt.Println("Example: go run . generate 5")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid day number: %v", err)
	}

	year := time.Now().Year()
	if len(os.Args) > 3 {
		year, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("Invalid year: %v", err)
		}
	}

	if err := GenerateDay(day, year); err != nil {
		log.Fatalf("Failed to generate day %d: %v", day, err)
	}
}

func PrintHelp() {
	fmt.Println("🎄 Advent of Code Template Generator")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  go run .                    - Auto-discover and run all solutions")
	fmt.Println("  go run . generate <day>     - Generate files for a specific day")
	fmt.Println("  go run . gen <day> [year]   - Short form of generate")
	fmt.Println("  go run . help               - Show this help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run . generate 1         - Generate day 1 for current year")
	fmt.Println("  go run . gen 5 2023         - Generate day 5 for year 2023")
	fmt.Println()
	fmt.Println("Workflow:")
	fmt.Println("  1. go run . generate 1      - Generate day from templates")
	fmt.Println("  2. Edit solutions/day01/ files and implement solutions")
	fmt.Println("  3. go run .                 - Automatically discovers and runs ALL days!")
	fmt.Println("  4. No manual imports needed - completely automatic!")
	fmt.Println()
	fmt.Println("Testing:")
	fmt.Println("  go test ./solutions/day01   - Test specific day")
	fmt.Println("  go test ./solutions/...     - Test all days")
	fmt.Println("  go test -bench=. ./solutions/day01 - Benchmark specific day")
}
