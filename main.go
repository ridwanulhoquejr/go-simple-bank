package main

import (
	"fmt"
	"log"
	"strconv"
)

// ? to implement this interface, any GO types must have a method String() which returns a string
type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Sold   Stringer
}

type Count int

// now we can say that Book implements the Stringer interface
func (b Book) String() string {
	return fmt.Sprintf("Book: %s by %s and Sold %s", b.Title, b.Author, b.Sold)
}

// now we can say that Count implements the Stringer interface
func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// ? so both Book and Count implement the Stringer interface,
// ? now we can pass both Book and Count objects to WriteLog as arguments
func WriteLog(s Stringer) {
	// at this moment, we don't know if s is a Book or a Count
	// but we know that it implements the Stringer interface
	// So in runtime, it will be either a Book or a Count and it will call the respective String() method
	log.Print(s.String())
}

func main() {

	// ? now we can pass Book and Count to WriteLog
	b := Book{
		Title:  "The Art of Computer Programming",
		Author: "Donald Knuth",
		// Sold:   Book{},
		Sold: Count(10),
	}
	WriteLog(b)

	c := Count(3)
	WriteLog(c)

	//! if we try to pass a type that does not implement the Stringer interface, we will get a compile-time error
	// i := 3
	// WriteLog(i)
}
