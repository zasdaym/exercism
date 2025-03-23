package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	opts := searchOptions{
		pattern: pattern,
		files:   files,
	}

	for _, flag := range flags {
		switch flag {
		case "-n":
			opts.showLineNumber = true
		case "-v":
			opts.invert = true
		case "-i":
			opts.ignoreCase = true
		case "-l":
			opts.printNameOnly = true
		case "-x":
			opts.matchEntireLine = true
		}
	}
	return search(opts)
}

type searchOptions struct {
	pattern         string
	files           []string
	showLineNumber  bool
	ignoreCase      bool
	invert          bool
	printNameOnly   bool
	matchEntireLine bool
}

func search(opts searchOptions) []string {
	var result []string
	showFilename := len(opts.files) > 1
	matchedFiles := make(map[string]bool)

	pattern := opts.pattern
	if opts.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	for _, filename := range opts.files {
		file, err := os.Open(filename)
		if err != nil {
			return nil
		}

		scanner := bufio.NewScanner(file)
		var lineNumber int

		for scanner.Scan() {
			line := scanner.Text()
			curr := line
			lineNumber += 1

			if opts.ignoreCase {
				curr = strings.ToLower(curr)
			}

			var matched bool
			if opts.matchEntireLine {
				if curr == pattern {
					matched = true
				}
			} else {
				if strings.Contains(curr, pattern) {
					matched = true
				}
			}
			if opts.invert {
				matched = !matched
			}
			if !matched {
				continue
			}

			if opts.showLineNumber {
				line = fmt.Sprintf("%d:%s", lineNumber, line)
			}

			if showFilename {
				line = fmt.Sprintf("%s:%s", filename, line)
			}

			if opts.printNameOnly {
				if matchedFiles[filename] {
					continue
				}
				matchedFiles[filename] = true
				line = filename
			}

			result = append(result, line)
		}

		file.Close()

		if err := scanner.Err(); err != nil {
			continue
		}
	}

	return result
}
