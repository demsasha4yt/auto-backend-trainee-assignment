package base62_test

import (
	"testing"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/base62"
)

func TestBase62_Decode(t *testing.T) {
	testData := []struct {
		key      string
		err      bool
		expected int64
	}{
		{"0", false, 0},
		{"a", false, 10},
		{"aa", false, 630},
		{"abc123EFG", false, 2222821365901088},
		{"hjNv8tS3K", false, 3781504209452600},
		{"hjNv8tS3K-", true, 0},
	}
	for _, tc := range testData {
		n, err := base62.Decode(tc.key)
		if n != tc.expected {
			t.Fatalf("expected %v, but got %v", tc.expected, n)
		}
		if err != nil && !tc.err {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}

func TestBase62_Encode(t *testing.T) {
	testData := []struct {
		n        int64
		expected string
	}{
		{0, "0"},
		{10, "a"},
		{630, "aa"},
		{2222821365901088, "abc123EFG"},
		{3781504209452600, "hjNv8tS3K"},
	}
	for _, tc := range testData {
		r := base62.Encode(tc.n)
		if r != tc.expected {
			t.Fatalf("encode expected '%v', but got '%v'", tc.expected, r)
		}
	}
}
