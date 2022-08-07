package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) MarshalText() ([]byte, error) {
	fullName := fmt.Sprintf("%s %s", p.FirstName, p.LastName)
	return []byte(fullName), nil
}

func (p *Person) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}

	s := string(text)
	parts := strings.Split(s, " ")

	if len(parts) < 2 {
		return errors.New("Please provide first and last name")
	}

	*p = Person{
		FirstName: parts[0],
		LastName:  parts[1],
	}

	return nil
}

func main() {
	var p Person
	defaultValue := Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	flag.TextVar(&p, "person", defaultValue, "Enter first and last name")

	flag.Parse()

	fmt.Println(p)
}
