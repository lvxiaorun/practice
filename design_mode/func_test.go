package design_mode

import (
	"log"
	"testing"
)

func TestNewPeople(t *testing.T) {
	p := NewPeople()
	log.Println(p)
	opts := []Option{WithName("xiaorun"), WithLike("zhangtian")}
	for _, item := range opts {
		item(p)
	}
	log.Println(p)
}
