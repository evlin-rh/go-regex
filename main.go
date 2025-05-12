package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	simpleRegex()

	LiteralPatterns()
}

func simpleRegex() {
	//Regular Expression with regexp
	match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	fmt.Println(match)

	// Regular Expression Rules
	r, _ := regexp.Compile("p([a-z]+)ch")
	peach := "peac"
	fmt.Println(r.MatchString(peach))

	// Makes you find the string to match
	fmt.Println(r.FindString("peach punch"))

	// Returns start and end index for match instead of String
	fmt.Println("idx:", r.FindStringIndex(("peach punch")))

	//Identical but only finds submatches
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Matches all string
	fmt.Println(r.FindAllString("peach punch pinch", -1)) //-1 means to print out all

	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2)) // Non-neg number acts as a limit

	// Can just use Match if we have []byte
	fmt.Println(r.Match([]byte("peach")))

	// Use of Panic if program cannot run safely bc of Error with MustCompile
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	//	Replacing Strings
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// ReplaceAllFunc allows you to tranform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
