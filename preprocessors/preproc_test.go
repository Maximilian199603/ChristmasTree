package preprocessors

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestAsLines(t *testing.T) {
	path := filepath.Join("testdata", "lines.txt")
	expected := []string{
		"line1",
		"line2",
		"line3",
	}
	result, err := AsLines(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAsGrid(t *testing.T) {
	path := filepath.Join("testdata", "grid.txt")
	expected := [][]rune{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	result, err := AsGrid(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAsWords(t *testing.T) {
	path := filepath.Join("testdata", "words.txt")
	expected := []string{
		"word1", "word2", "word3", "word4",
	}
	result, err := AsWords(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSplitByEmptyLines(t *testing.T) {
	path := filepath.Join("testdata", "split.txt")
	expected := [][]string{
		{"chunk1_line1", "chunk1_line2"},
		{"chunk2_line1", "chunk2_line2", "chunk2_line3"},
		{"chunk3_line1"},
	}
	result, err := SplitByEmptyLines(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
