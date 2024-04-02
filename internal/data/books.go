package data

import "time"

type Book struct {
	ID        int64     `json:"id`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Geners    []string  `json:"geners,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Verstion  int32     `json:"-"`
}
