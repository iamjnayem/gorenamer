package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	prefix    string
	suffix    string
	replace   string
	startWith string
	ext       string
	dryRun    bool
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Batch rename files based on patterns",
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."

		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("❌ Failed to read directory:", err)
			return
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			oldName := file.Name()

			// Filter: extension
			if ext != "" && filepath.Ext(oldName) != ext {
				continue
			}

			// Filter: startWith
			if startWith != "" && !strings.HasPrefix(oldName, startWith) {
				continue
			}

			newName := oldName

			// Replace
			if replace != "" {
				parts := strings.SplitN(replace, ":", 2)
				if len(parts) == 2 {
					newName = strings.ReplaceAll(newName, parts[0], parts[1])
				}
			}

			// Prefix
			if prefix != "" {
				newName = prefix + newName
			}

			// Suffix (before extension)
			if suffix != "" {
				ext := filepath.Ext(newName)
				name := strings.TrimSuffix(newName, ext)
				newName = name + suffix + ext
			}

			// Final Rename
			if newName != oldName {
				fmt.Printf("🔁 %s -> %s\n", oldName, newName)
				if !dryRun {
					err := os.Rename(filepath.Join(dir, oldName), filepath.Join(dir, newName))
					if err != nil {
						fmt.Println("❌ Rename failed:", err)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	renameCmd.Flags().StringVar(&prefix, "prefix", "", "Add prefix to filenames")
	renameCmd.Flags().StringVar(&suffix, "suffix", "", "Add suffix to filenames (before extension)")
	renameCmd.Flags().StringVar(&replace, "replace", "", "Replace pattern in filename (format: from:to)")
	renameCmd.Flags().StringVar(&startWith, "start-with", "", "Only rename files that start with this prefix")
	renameCmd.Flags().StringVar(&ext, "ext", "", "Only rename files with this extension (e.g., .txt)")
	renameCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be renamed without actually renaming")
}
