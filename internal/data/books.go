package data

import "time"

type Book struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Published int
	Pages     int
	Geners    []string
	Rating    float32
	Verstion  int32
}
