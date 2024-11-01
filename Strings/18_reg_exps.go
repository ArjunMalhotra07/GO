package strings

import (
	"fmt"
	"regexp"
)

func DoRegularExpressions18() {
	// regexp6()
	fmt.Println(regexp7("1010"))
}
func regexp1() {

	const pattern string = `^[a-z]{1,}$`
	check := []string{"abcAda", "", "123", "abcdea", "a", "z", "Aba"}
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	fmt.Println("All lowercase with same start and end characters")
	for _, stringToCheck := range check {
		var doesMatch bool = re.MatchString(stringToCheck)
		if !doesMatch {
			fmt.Printf("%s %v \n", stringToCheck, doesMatch)
			continue
		}
		doesMatch = stringToCheck[0] == stringToCheck[len(stringToCheck)-1]
		fmt.Printf("%s %v \n", stringToCheck, doesMatch)
	}
}

func regexp2() {
	const pattern string = `^a[a-z]*`
	check := []string{"abcAda", "", "123", "abcdea", "a", "z", "applesauce123"}
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	fmt.Println("Start with 'a' and are followed by zero or more lowercase letters")
	for _, stringToCheck := range check {
		var doesMatch bool = re.MatchString(stringToCheck)
		fmt.Printf("%s %v \n", stringToCheck, doesMatch)
	}
}
func regexp3() {
	const pattern string = `^[a-z]{3}$`
	check := []string{"abcAda", "", "123", "abc", "a", "z", "Abc"}
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	fmt.Println("Only 3 lowercase")
	for _, stringToCheck := range check {
		var doesMatch bool = re.MatchString(stringToCheck)
		fmt.Printf("%s %v \n", stringToCheck, doesMatch)
	}
}

func regexp4() {
	const pattern string = `^[A-Z]{1,5}$`
	check := []string{"ABC", "", "123", "AbD", "ABCDE", "ABCDEF", "Abc"}
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	fmt.Println("strings that contain only uppercase letters (A-Z) and are between 1 and 5 characters long")
	for _, stringToCheck := range check {
		var doesMatch bool = re.MatchString(stringToCheck)
		fmt.Printf("%s %v \n", stringToCheck, doesMatch)
	}
}

func regexp5() {
	const pattern string = `^s.*5$`
	check := []string{"s45", "", "5", "s34@905", "s_5", "sBCDEF", "6b5"}
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	fmt.Println("start with the letter 's' (case-sensitive), followed by any number of characters (including none), and end with the digit '5'")
	for _, stringToCheck := range check {
		var doesMatch bool = re.MatchString(stringToCheck)
		fmt.Printf("%s %v \n", stringToCheck, doesMatch)
	}
}

func regexp6() {

	const pattern string = `^[a-z]{1,}$`
	check := []string{"abcAda", "", "123", "abcdea", "a", "z", "Aba"}
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	fmt.Println("All lowercase with same start and end characters")
	for _, stringToCheck := range check {
		var doesMatch bool = re.MatchString(stringToCheck)
		if !doesMatch {
			fmt.Printf("%s %v \n", stringToCheck, doesMatch)
			continue
		}
		doesMatch = stringToCheck[0] == stringToCheck[len(stringToCheck)-1]
		fmt.Printf("%s %v \n", stringToCheck, doesMatch)
	}
}

func regexp7(s string) bool {
	regex := regexp.MustCompile(`^[01]+$`)
	if !regex.MatchString(s) {
		return false
	}

	// Count zeros
	count := 0
	for _, ch := range s {
		if ch == '0' {
			count++
		}
	}
	return count%2 == 0
}
