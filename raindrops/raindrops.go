package raindrops

import (
	"fmt"
)

// Convert number to rain talk
func Convert(input int) string  {
	response := ""

	if input % 3 == 0 {
		response += "Pling"
	} 
	if input % 5 == 0 {
		response += "Plang"
	} 
	if input % 7 == 0 {
		response += "Plong"
	}

	if response == "" {
		return fmt.Sprintf("%d", input)
	}

	return response
	
}