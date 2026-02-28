package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ProcessDirectory(rootDir, patternKey, outFile string) error {
	extensions, ok := FilePatterns[patternKey]
	if !ok {
		return fmt.Errorf("unknown pattern key: %s", patternKey)
	}

	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if IgnoreDirs[d.Name()] {
				return filepath.SkipDir
			}
		}

		if !hasExtension(path, extensions) {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Errorf("warning: could not read %s: %w", path, err)
			return nil
		}

		headerText := fmt.Sprintf("# %s", path)
		separator := strings.Repeat("=", len(headerText))

		fmt.Fprintln(f, separator)
		fmt.Fprintln(f, headerText)
		fmt.Fprintln(f, separator)

		f.Write(content)
		if !bytesEndWithNewLine(content) {
			f.Write([]byte("\n"))
		}
		fmt.Fprintln(f)

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}

	return nil
}

func hasExtension(path string, allowed []string) bool {
	for _, ext := range allowed {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

func bytesEndWithNewLine(b []byte) bool {
	return len(b) > 0 && b[len(b)-1] == '\n'
}
