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

func IFlag() {
	text := "Apples are delicious but so are apples, apPleS and APPLE."

	pattern := regexp.MustCompile(`(?i)apple`)
	appleMatches := pattern.FindAllString(text, -1)
	fmt.Println(appleMatches)
}

// G Flag does not exist in GO. However
// FindAllString is the substitute for that

// M Flag -> Multiline match.
// Probs not as useful

func MFlag() {
	paragraph := "Hello this is hello \n Hello to you hello. \n Hello hello"

	pattern := regexp.MustCompile(`(?mi)hello`)
	match := pattern.FindAllString(paragraph, -1)
	fmt.Println(match)
}

// S Flag treats the input as 1-line-string

func SFlag() {
	p := "Start of line \n 2nd Line \n Another line."

	pattern := regexp.MustCompile(`(?s)Start.*line`)
	match := pattern.FindAllString(p, -1)
	fmt.Println(match)
}

// Character Classes
// \d -> digit 	\w -> word 	\s -> whitespace
// \D -> non-digit.... \W	\S

func CharacterClasses() {
	text := "The price is $20, the code is ABC123"

	pattern := regexp.MustCompile(`\w+\d\w*`)
	match := pattern.FindAllString(text, -1)
	fmt.Println(match)
}

// Find all characters that exclude these
// ^ -> Negate Characters

func NegatedChar() {
	text := "Hello, how are you"

	pattern := regexp.MustCompile(`[^aeiouAEIOU\s]`)
	matches := pattern.FindAllString(text, -1)

	fmt.Println(matches)
}

func Range() {
	text := "The numbers are 5, 25, 48, 60"

	pattern := regexp.MustCompile(`[1-4][0-9]`)
	match := pattern.FindAllString(text, -1)
	fmt.Println(match)
}

// Quantifiers

func Asterisk() {
	pattern := regexp.MustCompile(`o*k`)
	fmt.Println(pattern.FindString("oogenesis"))
	fmt.Println(pattern.FindString("cook"))
	fmt.Println(pattern.FindString("oocyst"))
	fmt.Println(pattern.FindString("book"))
	fmt.Println()
}

// + := 1 or more
// i.e `go+` -> Has 1 g, has 1..* o

func Plus() {
	pattern := regexp.MustCompile(`go+`)
	fmt.Println(pattern.FindAllString("goo", -1))
	fmt.Println(pattern.FindAllString("goooo", -1))
	fmt.Println(pattern.FindAllString("go", -1))
	fmt.Println(pattern.FindAllString("g", -1))
	fmt.Println()
}

// {min, max}

func MinMax() {
	pattern := regexp.MustCompile(`go{2,4}`)
	fmt.Println(pattern.FindAllString("goo", -1))
	fmt.Println(pattern.FindAllString("goooooo", -1))
	fmt.Println(pattern.FindAllString("go", -1))
	fmt.Println(pattern.FindAllString("g", -1))
	fmt.Println()
}

// ? := matches 0..1
// Lazy Quantifer means to match as little as possible

func Lazy() {
	pattern := regexp.MustCompile(`".*?"`)
	fmt.Println(pattern.FindString(`"This hello" "This is"`))
}

// Repitition
// () -> Looks for repetitiion
// []byte(string), []byte(replacement) -> used to replace string

func Parantheses() {
	text := `The detective carefully examined the crime scene. It was a gruesome sight—the victim had been brutally killed. The evidence pointed towards a professional hitman as the perpetrator. The investigators were determined to solve the case and bring the killer to justice. The motive behind the killing remained unclear, but the police were determined to uncover the truth.`

	pattern := regexp.MustCompile(`(kill)+`)
	modStr := pattern.ReplaceAll([]byte(text), []byte("---"))
	fmt.Println(string(modStr))
}

func Dot() {
	pattern := regexp.MustCompile(`h.t`)

	fmt.Println(pattern.MatchString("hat"))
	fmt.Println(pattern.MatchString("hot"))
	fmt.Println()
}

// Match 1 or more pattern
// OR Operator

func Pipe() {
	pattern := regexp.MustCompile(`Dog|Cat`)

	fmt.Println(pattern.MatchString("Dog"))
	fmt.Println(pattern.MatchString("Cat"))
	fmt.Println()
}

// Use \ to register special characters

func SpecialChar() {
	pattern := regexp.MustCompile(`\d+.\d+`)

	fmt.Println(pattern.MatchString("3.14"))
	fmt.Println(pattern.MatchString("23"))
	fmt.Println()
}

func Anchors() {
	pattern := regexp.MustCompile(`^start.*end$`)

	fmt.Println(pattern.MatchString("start at the end"))
	fmt.Println(pattern.MatchString("start at the bottom"))
	fmt.Println()
}

// Word boundaries are used to match positions between  word characters and non-word characters

func WordBoundaries() {
	// bet at start of string
	pattern := regexp.MustCompile(`(?i)\bbet`)
	fmt.Println(pattern.FindAllString("Betty's better bet was to buy the blue blouse.", -1))

	// sion at the end of string
	re := regexp.MustCompile(`(?i)sion\b`)
	fmt.Println(re.FindAllString("After much discussion, the team came to a consensus on the vision for the project.", -1))

	// start and end of word.  i.e -> whole word
	re2 := regexp.MustCompile(`(?i)\bcat\b`)
	fmt.Println(re2.FindAllString("cat is in the feline category of animals", -1))
}

func Grouping() {
	pattern := regexp.MustCompile(`(quick|lazy) (brown|gray) (fox|dog)`)
	text := "the lazy gray dog jumps over the lazy fox"
	match := pattern.FindAllString(text, -1)
	fmt.Println(match)
}

func CapturingGroups() {
	pattern := regexp.MustCompile(`(\d{2})-(\d{2})-(\d{4})`)
	modDate := pattern.ReplaceAllString("daTE: 12-31-2024", "$2/$1/$3")
	fmt.Println(modDate)
}
