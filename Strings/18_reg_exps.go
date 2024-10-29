package strings

import (
	"fmt"
	"regexp"
)

func DoRegularExpressions18() {
	regexp1()
}
func regexp1() {
	pattern := `^(a|b)+(a|b)$`
	var re *regexp.Regexp = regexp.MustCompile(pattern)
	var doesMatch bool = re.MatchString("baba")
	fmt.Println(doesMatch)

}
