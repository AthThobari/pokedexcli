package main

import "strings"

// cleanInput cleans and breaks up user input
func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}
