package models

import "testing"

// TestLink returns test link
func TestLink(t *testing.T) *Links {
	t.Helper()
	return &Links{
		URL: "google.com",
	}
}