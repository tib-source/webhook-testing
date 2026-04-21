package converter

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/yuin/goldmark"
)

// Result holds the output of a markdown conversion.
type Result struct {
	HTML        string
	Title       string
	WordCount   int
	ReadingTime int // minutes
}

// Convert takes raw markdown and returns the rendered HTML along with metadata.
func Convert(markdown string) (Result, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return Result{}, fmt.Errorf("converting markdown: %w", err)
	}

	title := ExtractTitle(markdown)
	words := CountWords(markdown)

	return Result{
		HTML:        buf.String(),
		Title:       title,
		WordCount:   words,
		ReadingTime: EstimateReadingTime(words),
	}, nil
}

// ExtractTitle returns the first H1 heading from the markdown, or empty string.
func ExtractTitle(markdown string) string {
	for _, line := range strings.Split(markdown, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
		}
	}
	return ""
}

// CountWords returns the number of words in the given text.
func CountWords(text string) int {
	return len(strings.Fields(text))
}

// EstimateReadingTime returns estimated minutes to read based on word count.
// Uses an average reading speed of 200 words per minute, minimum 1 minute.
func EstimateReadingTime(wordCount int) int {
	if wordCount == 0 {
		return 0
	}
	minutes := wordCount / 200
	if minutes < 1 {
		minutes = 1
	}
	return minutes
}

// CharCount returns the number of characters (Unicode code points) in the text.
func CharCount(text string) int {
	return utf8.RuneCountInString(text)
}
