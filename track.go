package main

import "fmt"

type Track struct {
	title  string
	artist string
	album  string
	id     string
	length int
}

var track = Track{}

func (t Track) str() string {
	return fmt.Sprintf("%s|%s|%s|%s|%d", t.title, t.artist, t.album, t.id,
		t.length)
}
