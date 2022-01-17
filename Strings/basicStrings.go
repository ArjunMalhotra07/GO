package main

import (
	"fmt"
	"strings"
	s "strings"
)

func main() {
	p := fmt.Println
	f := fmt.Printf
	p("hlo")
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("fooo", "o", "0", -1))
	p("Replace:   ", s.Replace("fooo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

	name := "Arjun"
	for idx, s := range name {
		f("index of %c is %d\n", s, idx)
	}

	langs := []string{"F#", "Go", "Python", "Perl", "Erlang"}

	s := strings.Join(langs, "+ ")
	fmt.Println(s)

	data := strings.Split(s, "+")
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
	for _, day := range weekdays {
		fmt.Println(day)
	}

}
