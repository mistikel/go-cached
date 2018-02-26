package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var bookId int

type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"publisher"`
	Reference string `json:"reference"`
}

func init() {
	Store(Book{
		Title:     "As I Lay Dying",
		Author:    "William Faulkner",
		Reference: "The Odyssey, Homer",
	})
	Store(Book{
		Title:     "All the King's Men",
		Author:    "Robert Penn Warren",
		Reference: "Humpty Dumpty",
	})
}

func Store(b Book) (Book, error) {
	bookId++

	b.Id = bookId

	c := NewRedis()
	defer c.Close()

	d, err := json.Marshal(b)
	HandleError(err)

	rest, err := c.Do("SET", "book:"+strconv.Itoa(b.Id), d)
	HandleError(err)

	fmt.Println("Get ", rest)
	return b, nil
}

func Get(id int) (Book, error) {
	var book Book
	c := NewRedis()
	defer c.Close()

	b, err := c.Do("GET", "book:"+strconv.Itoa(id))
	HandleError(err)

	if err = json.Unmarshal(b.([]byte), &book); err != nil {
		panic(err)
	}

	return book, nil
}

func Gets() ([]Book, error) {
	books := make([]Book, 0)

	c := NewRedis()
	defer c.Close()

	keys, err := c.Do("KEYS", "book:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {
		var b Book

		rest, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err = json.Unmarshal(rest.([]byte), &b); err != nil {
			panic(err)
		}

		books = append(books, b)
	}

	return books, nil
}

func Delete(id int) error {
	c := NewRedis()
	defer c.Close()

	rest, err := c.Do("DEL", "book:"+strconv.Itoa(id))
	HandleError(err)

	if rest.(int) != 1 {
		return fmt.Errorf("No post removed")
	}

	return nil
}

func Put(book Book) (Book, error) {
	c := NewRedis()
	defer c.Close()

	Delete(book.Id)
	Store(book)

	return book, nil
}
