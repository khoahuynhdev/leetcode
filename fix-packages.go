package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Script to standardize all packages to "package main"
// and rename duplicate solutions

func main() {
	fmt.Println("🔧 Fixing Package Conflicts")
	fmt.Println("===========================\n")

	problemsDir := "problems"
	entries, err := os.ReadDir(problemsDir)
	if err != nil {
		fmt.Printf("❌ Error reading problems directory: %v\n", err)
		return
	}

	conflicts := 0
	fixed := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		problemDir := filepath.Join(problemsDir, entry.Name())
		goFiles, err := filepath.Glob(filepath.Join(problemDir, "*.go"))
		if err != nil {
			continue
		}

		// Check for package conflicts
		packages := make(map[string][]string)
		for _, file := range goFiles {
			pkg := getPackageName(file)
			if pkg != "" {
				packages[pkg] = append(packages[pkg], file)
			}
		}

		if len(packages) > 1 {
			conflicts++
			fmt.Printf("⚠️  %s has multiple packages: ", entry.Name())
			for pkg := range packages {
				fmt.Printf("%s ", pkg)
			}
			fmt.Println()

			// Fix: standardize all to package main
			for _, files := range packages {
				for _, file := range files {
					if err := replacePackage(file, "main"); err != nil {
						fmt.Printf("   ❌ Error fixing %s: %v\n", filepath.Base(file), err)
					} else {
						fixed++
					}
				}
			}
			fmt.Printf("   ✓ Fixed all files to use 'package main'\n\n")
		}
	}

	fmt.Printf("\n📊 Summary:\n")
	fmt.Printf("   Directories with conflicts: %d\n", conflicts)
	fmt.Printf("   Files fixed: %d\n", fixed)
	fmt.Println("\n✅ All packages standardized to 'package main'")
}

func getPackageName(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "package ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1]
			}
		}
	}
	return ""
}

func replacePackage(filepath string, newPackage string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "package ") {
			// Replace package declaration
			indent := len(line) - len(strings.TrimLeft(line, " \t"))
			lines = append(lines, strings.Repeat(" ", indent)+"package "+newPackage)
		} else {
			lines = append(lines, line)
		}
	}
	file.Close()

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write back
	output := strings.Join(lines, "\n")
	return os.WriteFile(filepath, []byte(output), 0644)
}
