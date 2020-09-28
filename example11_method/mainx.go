package main

import (
	"fmt"
)

type text string

type rectangle struct {
	width  int
	length int
}

func area(rec rectangle) int {
	return rec.width * rec.length
}

func (rec rectangle) area() int {
	return rec.width * rec.length
}

/// method getter setter
func (t *text) set(s string) {
	*t = text(s)
}

func (t *text) get() string {
	return string(*t)
}

func main() {
	rec := rectangle{2, 3}
	fmt.Println(area(rec))
	fmt.Println(rec.area())

	// =====================
	var t text
	t.set("settext")
	fmt.Println(t.get())
}
