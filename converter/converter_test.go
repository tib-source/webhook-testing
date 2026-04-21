package converter

import (
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	input := "# Hello World\n\nThis is a **bold** statement."
	result, err := Convert(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(result.HTML, "<h1>Hello World</h1>") {
		t.Errorf("expected H1 tag, got: %s", result.HTML)
	}
	if !strings.Contains(result.HTML, "<strong>bold</strong>") {
		t.Errorf("expected strong tag, got: %s", result.HTML)
	}
	if result.Title != "Hello World" {
		t.Errorf("expected title 'Hello World', got: %s", result.Title)
	}
}

func TestConvertEmptyInput(t *testing.T) {
	result, err := Convert("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.HTML != "" {
		t.Errorf("expected empty HTML, got: %s", result.HTML)
	}
	if result.Title != "" {
		t.Errorf("expected empty title, got: %s", result.Title)
	}
	if result.WordCount != 0 {
		t.Errorf("expected 0 words, got: %d", result.WordCount)
	}
}

func TestExtractTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic title", "# My Title\n\nContent here.", "My Title"},
		{"no title", "Just some text.", ""},
		{"h2 not h1", "## Subtitle\n\nContent.", ""},
		{"title with extra spaces", "#   Spaced Title  \n\nBody.", "Spaced Title"},
		{"title after blank line", "\n\n# Late Title", "Late Title"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractTitle(tt.input)
			if got != tt.expected {
				t.Errorf("ExtractTitle(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"", 0},
		{"one", 1},
		{"  spaced   out  words  ", 3},
		{"line\none\ntwo", 3},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := CountWords(tt.input)
			if got != tt.expected {
				t.Errorf("CountWords(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestEstimateReadingTime(t *testing.T) {
	tests := []struct {
		name      string
		wordCount int
		expected  int
	}{
		{"zero words", 0, 0},
		{"short text", 50, 1},
		{"one minute boundary", 200, 1},
		{"two minutes", 400, 2},
		{"long text", 1000, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EstimateReadingTime(tt.wordCount)
			if got != tt.expected {
				t.Errorf("EstimateReadingTime(%d) = %d, want %d", tt.wordCount, got, tt.expected)
			}
		})
	}
}

func TestCharCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"cafe\u0301", 5}, // café with combining accent
		{"\u4e16\u754c", 2},       // Chinese characters
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := CharCount(tt.input)
			if got != tt.expected {
				t.Errorf("CharCount(%q) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
