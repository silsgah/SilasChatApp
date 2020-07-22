package main

import "fmt"

//Book some thing here
type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "The Good book of our Lord",
		Author:        "Moses Adams",
		YearPublished: 1989,
	},
	Book{
		ID:            2,
		Title:         "The hobbit",
		Author:        "J.R.R Tokien",
		YearPublished: 1937,
	},
	Book{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	Book{
		ID:            4,
		Title:         "Les Miserables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
	Book{
		ID:            5,
		Title:         "Harry Potter and the Philospher's Stone",
		Author:        "J.k Rowling",
		YearPublished: 1997,
	},
}
