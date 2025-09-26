package internal

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func GenerateDay(day int, year int) error {
	if day < 1 || day > 25 {
		return fmt.Errorf("day must be between 1 and 25")
	}

	dayData := DayData{
		Day:       day,
		DayPadded: fmt.Sprintf("%02d", day),
		Year:      year,
		Package:   fmt.Sprintf("day%02d", day),
		Title:     fmt.Sprintf("Day %d", day),
		Date:      time.Now().Format("2006-01-02"),
	}

	// Create day directory
	dayDir := fmt.Sprintf("solutions/day%02d", day)
	if err := os.MkdirAll(dayDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", dayDir, err)
	}

	// Generate files from templates
	files := []string{"part1.go", "part2.go", fmt.Sprintf("day%02d_test.go", day)}
	templates := []string{"part1.go.tmpl", "part2.go.tmpl", "test.go.tmpl"}

	for i, filename := range files {
		templatePath := filepath.Join("templates", templates[i])
		outputPath := filepath.Join(dayDir, filename)

		if err := generateFromTemplate(templatePath, outputPath, dayData); err != nil {
			return fmt.Errorf("failed to generate %s: %v", filename, err)
		}
	}

	// Handle input file
	inputPath := filepath.Join(dayDir, "input.txt")
	sessionCookie := promptForSessionCookie()

	if sessionCookie != "" {
		if err := downloadInput(year, day, sessionCookie, inputPath); err != nil {
			fmt.Printf("Warning: Could not download input for day %d: %v\n", day, err)
			fmt.Printf("Creating empty input file at %s\n", inputPath)
			if err := createEmptyFile(inputPath); err != nil {
				return fmt.Errorf("failed to create empty input file: %v", err)
			}
		} else {
			fmt.Printf("✅ Downloaded input for day %d\n", day)
		}
	} else {
		if err := createEmptyFile(inputPath); err != nil {
			return fmt.Errorf("failed to create empty input file: %v", err)
		}
	}

	fmt.Printf("🎉 Generated day %02d structure successfully!\n", day)
	fmt.Printf("📁 Directory: %s/\n", dayDir)
	fmt.Printf("📄 Files: part1.go, part2.go, day%02d_test.go, input.txt\n", day)
	fmt.Printf("\n🎯 Ready to code! Just edit the files in %s/ and run 'go run .' when done!\n", dayDir)

	return nil
}

func generateFromTemplate(templatePath, outputPath string, data DayData) error {
	if _, err := os.Stat(outputPath); err == nil {
		return fmt.Errorf("file %s already exists", outputPath)
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %v", templatePath, err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", outputPath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	return nil
}

func downloadInput(year, day int, sessionCookie, outputPath string) error {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))
	req.Header.Set("User-Agent", "Go AoC Template Generator")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func createEmptyFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func promptForSessionCookie() string {
	fmt.Println("\n🍪 Optional: Enter your Advent of Code session cookie to auto-download input")
	fmt.Println("   (Find this in your browser's dev tools under Application/Storage > Cookies)")
	fmt.Println("   Leave empty to create empty input files:")
	fmt.Print("Session cookie: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
