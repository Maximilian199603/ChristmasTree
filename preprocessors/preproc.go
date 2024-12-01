package preprocessors

import (
	"bufio"
	"os"
	"strings"
)

func AsLines(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func AsGrid(filepath string) ([][]rune, error) {
	lines, err := AsLines(filepath)
	if err != nil {
		return nil, err
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid, nil
}

func AsWords(filepath string) ([]string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(content)), nil
}

func SplitByEmptyLines(filepath string) ([][]string, error) {
	lines, err := AsLines(filepath)
	if err != nil {
		return nil, err
	}

	var chunks [][]string
	var currentChunk []string
	for _, line := range lines {
		if line == "" {
			if len(currentChunk) > 0 {
				chunks = append(chunks, currentChunk)
				currentChunk = []string{}
			}
		} else {
			currentChunk = append(currentChunk, line)
		}
	}
	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}
	return chunks, nil
}
