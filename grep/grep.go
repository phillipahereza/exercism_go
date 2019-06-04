package grep

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Search for a string in text file
func Search(pattern string, flags, files []string) []string { 
	var result []string
	for _, file := range files {
		# get file
		
	}
	return result
}

func getMatch(pattern, line string) bool {
	return strings.Contains(line, pattern)
}
