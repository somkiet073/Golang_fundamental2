package model

type Book struct {
	Title    string
	Author   string
	Subtitle string
	Price    float64
	Book     BookList
}

type BookList struct {
	BookCode    int
	BookName    string
	BookCreated string
}
