package basicsagain

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func TestDataTypes() {
	// integerDataTypes()
	// stringDataTypes()
	pointerTypes()
}
func integerDataTypes() {
	var a int8 = 127
	var b int16 = -32768
	var c int32 = -2147483648
	var d int64 = 9223372036854775807
	var e int = 78
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	var f float32 = 12345678.9
	var g float64 = 12345678.9
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", g, g)

	var myVal, err = divide(5, 0)
	fmt.Println(myVal, err.Error())
}

func divide(n, d int) (int, error) {
	if d != 0 {
		return n / d, nil
	} else {
		return -1, errors.New("can't divide by zero")
	}
}

func stringDataTypes() {
	var myRune rune = 'a'
	fmt.Println(myRune)
	me := "Arjun"
	length := utf8.RuneCountInString(me)
	fmt.Println(length)
	name2 := "résumé"
	fmt.Printf("Value: %v, Type: %T, Length: %v\n", name2[2], name2[2], len(name2))
	for i, v := range name2 {
		fmt.Println(i, v)
	}
	name3 := "é"
	fmt.Println("Length of String using len method", name3, "is", len(name3))
	//! Runes are alias for int32
	name4 := []rune("résumé")
	fmt.Printf("Value: %v, Type: %T, Length: %v\n", name4[2], name4[2], len(name4))

	for i, v := range name4 {
		fmt.Println(i, v)
	}
	fmt.Println(createStringMethod1())
	fmt.Println(createStringMethod2())

}

func createStringMethod1() time.Duration {
	t0 := time.Now()
	var strSlice = []string{"A", "r", "j", "u", "n", " ", "M", "a", "l", "h", "o", "t", "r", "a", "A", "r", "j", "u", "n", " ", "M", "a", "l", "h", "o", "t", "r", "a", "A", "r", "j", "u", "n", " ", "M", "a", "l", "h", "o", "t", "r", "a"}
	var stringBuilder strings.Builder
	for i := range strSlice {
		stringBuilder.WriteString(strSlice[i])
	}
	var newString = stringBuilder.String()
	fmt.Println(newString)
	return time.Since(t0)
}
func createStringMethod2() time.Duration {
	t0 := time.Now()
	var strSlice = []string{"A", "r", "j", "u", "n", " ", "M", "a", "l", "h", "o", "t", "r", "a", "A", "r", "j", "u", "n", " ", "M", "a", "l", "h", "o", "t", "r", "a", "A", "r", "j", "u", "n", " ", "M", "a", "l", "h", "o", "t", "r", "a"}
	newString := ""
	for i := range strSlice {
		newString += strSlice[i]
	}
	fmt.Println(newString)
	return time.Since(t0)
}

func pointerTypes() {
	var p *int
	var i int
	p = &i
	*p = 43
	fmt.Println(p)
	fmt.Println(i)
	var myVar *int32 = new(int32)
	fmt.Println(myVar)
}

/*
	2^8 , 2^8-1
	int8 [-128, 127]

	2^15, 2^15-1
	int16 [-32768, 32767]

	2^31 , 2^31-1
	int32 [-2147483648, 2147483647]

	2^63 , 2^63-1
	int64 [-9223372036854775808 , 9223372036854775807 ]

	uint8 [0, 255]

	uint16 [0, 65535]

	uint32 [0, 4294967295]

	uint64 [0, 18446744073709551615]

*/
