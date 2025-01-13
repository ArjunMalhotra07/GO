package strings

import (
	"fmt"
	"strings"
)

func MainFunction10() {
	p := fmt.Println
	f := fmt.Printf
	p("hlo")
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("f.0o", "0", ".", -1))
	p("Replace:   ", strings.Replace("fooo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

	name := "Arjun"
	for idx, s := range name {
		f("index of %c is %d\n", s, idx)
	}

	langs := []string{"F#", "Go", "Python", "Perl", "Erlang"}

	s := strings.Join(langs, ", ")
	fmt.Println(s)

	data := strings.Split(s, ",")
	fmt.Println(data)

	w1 := "Hello"
	w2 := "hello"

	if strings.EqualFold(w1, w2) {
		p("Equal")
	} else {
		p("nott equal")
	}

	str := "sun,mon,tue,wed,thu,fri,sat"
	weekdays := strings.Split(str, ",")
	fmt.Println(weekdays)
	for _, day := range weekdays {
		fmt.Println(day)
	}

}
