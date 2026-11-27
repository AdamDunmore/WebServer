package classes

import (
	"time"
)

type MediaItem struct {
	path string
	file_path string
	file_name string
	cover string // TODO check if this should be a method
	artist string
	album string
	releaseDate time.Time 
}

func (m MediaItem) new() *MediaItem {
	nm := new(MediaItem);	
	nm.path = ""; // TODO fill out contructor
	return nm;
} 
