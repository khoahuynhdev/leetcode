package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Migration script to reorganize LeetCode solutions into standardized structure
// Usage: go run migrate.go

type Solution struct {
	Number      int
	Name        string
	SourcePath  string
	SourceFiles []string
}

func main() {
	fmt.Println("🔄 LeetCode Repository Migration")
	fmt.Println("================================\n")

	// Collect all solutions from different locations
	solutions := []Solution{}

	// 1. Find root-level numbered directories (e.g., 150/, 232/)
	rootSolutions, err := findRootLevelSolutions()
	if err != nil {
		fmt.Printf("❌ Error finding root-level solutions: %v\n", err)
		return
	}
	solutions = append(solutions, rootSolutions...)
	fmt.Printf("✓ Found %d solutions in root-level directories\n", len(rootSolutions))

	// 2. Find solution/* subdirectories (e.g., solution/104/)
	solutionDirSolutions, err := findSolutionDirectories()
	if err != nil {
		fmt.Printf("❌ Error finding solution/* directories: %v\n", err)
		return
	}
	solutions = append(solutions, solutionDirSolutions...)
	fmt.Printf("✓ Found %d solutions in solution/* directories\n", len(solutionDirSolutions))

	// 3. Find standalone files in solution/ (e.g., solution/136__single-number.go)
	standaloneSolutions, err := findStandaloneSolutions()
	if err != nil {
		fmt.Printf("❌ Error finding standalone solutions: %v\n", err)
		return
	}
	solutions = append(solutions, standaloneSolutions...)
	fmt.Printf("✓ Found %d standalone solution files\n", len(standaloneSolutions))

	fmt.Printf("\n📊 Total solutions to migrate: %d\n\n", len(solutions))

	// Ask for confirmation
	fmt.Print("⚠️  This will create a 'problems/' directory and copy all solutions.\n")
	fmt.Print("   Original files will NOT be deleted (you can review and delete manually).\n")
	fmt.Print("\nProceed? (yes/no): ")

	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) != "yes" {
		fmt.Println("❌ Migration cancelled")
		return
	}

	// Create problems/ directory
	if err := os.MkdirAll("problems", 0755); err != nil {
		fmt.Printf("❌ Error creating problems/ directory: %v\n", err)
		return
	}

	// Migrate each solution
	migrated := 0
	for _, sol := range solutions {
		if err := migrateSolution(sol); err != nil {
			fmt.Printf("❌ Error migrating %s: %v\n", sol.SourcePath, err)
		} else {
			migrated++
		}
	}

	fmt.Printf("\n✅ Migration complete! %d/%d solutions migrated\n", migrated, len(solutions))
	fmt.Println("\n📝 Next steps:")
	fmt.Println("   1. Review the new problems/ directory")
	fmt.Println("   2. Run: go test ./problems/...")
	fmt.Println("   3. If everything looks good, delete old directories:")
	fmt.Println("      - Root-level numbered directories (150/, 232/, etc.)")
	fmt.Println("      - solution/ directory")
	fmt.Println("   4. Update CLAUDE.md with new structure")
}

func findRootLevelSolutions() ([]Solution, error) {
	var solutions []Solution
	numRegex := regexp.MustCompile(`^(\d+)$`)

	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		matches := numRegex.FindStringSubmatch(entry.Name())
		if len(matches) == 2 {
			num, _ := strconv.Atoi(matches[1])
			files, _ := findGoFiles(entry.Name())

			// Try to extract problem name from files
			name := extractProblemName(files)

			solutions = append(solutions, Solution{
				Number:      num,
				Name:        name,
				SourcePath:  entry.Name(),
				SourceFiles: files,
			})
		}
	}

	return solutions, nil
}

func findSolutionDirectories() ([]Solution, error) {
	var solutions []Solution
	solutionDir := "solution"

	if _, err := os.Stat(solutionDir); os.IsNotExist(err) {
		return solutions, nil
	}

	entries, err := os.ReadDir(solutionDir)
	if err != nil {
		return nil, err
	}

	numRegex := regexp.MustCompile(`^(\d+)$`)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		matches := numRegex.FindStringSubmatch(entry.Name())
		if len(matches) == 2 {
			num, _ := strconv.Atoi(matches[1])
			sourcePath := filepath.Join(solutionDir, entry.Name())
			files, _ := findGoFiles(sourcePath)

			name := extractProblemName(files)

			solutions = append(solutions, Solution{
				Number:      num,
				Name:        name,
				SourcePath:  sourcePath,
				SourceFiles: files,
			})
		}
	}

	return solutions, nil
}

func findStandaloneSolutions() ([]Solution, error) {
	var solutions []Solution
	solutionDir := "solution"

	if _, err := os.Stat(solutionDir); os.IsNotExist(err) {
		return solutions, nil
	}

	// Pattern: 136__single-number.go
	fileRegex := regexp.MustCompile(`^(\d+)__(.+)\.go$`)

	entries, err := os.ReadDir(solutionDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		matches := fileRegex.FindStringSubmatch(entry.Name())
		if len(matches) == 3 {
			num, _ := strconv.Atoi(matches[1])
			name := strings.ReplaceAll(matches[2], "_", "-")

			solutions = append(solutions, Solution{
				Number:      num,
				Name:        name,
				SourcePath:  filepath.Join(solutionDir, entry.Name()),
				SourceFiles: []string{entry.Name()},
			})
		}
	}

	return solutions, nil
}

func migrateSolution(sol Solution) error {
	// Create target directory: problems/0001-problem-name/
	targetDir := fmt.Sprintf("problems/%04d-%s", sol.Number, sol.Name)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	// Copy all source files
	for _, file := range sol.SourceFiles {
		sourcePath := filepath.Join(sol.SourcePath, file)

		// For standalone files, sourcePath is already the full path
		if strings.HasSuffix(sol.SourcePath, ".go") {
			sourcePath = sol.SourcePath
		}

		// Determine target filename
		var targetFile string
		if strings.HasSuffix(file, "_test.go") {
			targetFile = "solution_test.go"
		} else if strings.HasSuffix(file, ".go") {
			// Check if we already have a solution.go
			solutionPath := filepath.Join(targetDir, "solution.go")
			if _, err := os.Stat(solutionPath); err == nil {
				// File exists, use alternative name
				baseName := strings.TrimSuffix(file, ".go")
				targetFile = baseName + ".go"
			} else {
				targetFile = "solution.go"
			}
		} else {
			targetFile = file
		}

		targetPath := filepath.Join(targetDir, targetFile)

		if err := copyFile(sourcePath, targetPath); err != nil {
			return fmt.Errorf("copy %s: %w", file, err)
		}
	}

	fmt.Printf("✓ Migrated %04d-%s (%s)\n", sol.Number, sol.Name, sol.SourcePath)
	return nil
}

func findGoFiles(dir string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func extractProblemName(files []string) string {
	// Try to extract a meaningful name from filenames
	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}

		name := strings.TrimSuffix(file, ".go")

		// Convert various formats to kebab-case
		// e.g., "maximumDepthOfBinaryTree" -> "maximum-depth-of-binary-tree"
		name = strings.ReplaceAll(name, "_", "-")

		// If name looks meaningful (not just "main" or "solution"), use it
		if name != "main" && name != "solution" && len(name) > 3 {
			return name
		}
	}

	// Fallback to generic name
	return "problem"
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	return destFile.Sync()
}
