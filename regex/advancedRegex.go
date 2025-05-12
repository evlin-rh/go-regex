package regex

import (
	"fmt"
	"regexp"
)

func LiteralPatterns() {
	pattern := regexp.MustCompile("Northen")
	text := "New Yokr is in the Northen part os US."

	if pattern.MatchString(text) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
