package main

import "fmt"

type song struct {
	title    string
	artist   string
	preSong  *song
	nextSong *song
}
type songList struct {
	title   string
	head    *song
	current *song
}

func (s *songList) setNextSong(node *song) {
	if s.head == nil {
		s.head = node
	} else {
		s.current.nextSong = node
	}
}
func (s *songList) setPresong(node *song) {
	if s.head == nil {
		s.head = node
		s.current = node
	} else {
		s.current.preSong = node
	}
}
func (s *songList) setBottomSonglist(node *song) error {
	if s.head == nil {
		s.head = node
		node.preSong = nil
	} else {
		s.getNextEmptyNode().nextSong = node
		node.preSong = s.getNextEmptyNode()
	}
	return nil
}
func (s *songList) playing() *song {
	return s.current
}
func (s *songList) playNext() *song {
	return s.current.nextSong
}
func (s *songList) getNextEmptyNode() *song {
	currentSong := s.head
	for currentSong.nextSong != nil {
		currentSong = currentSong.nextSong
	}
	return currentSong
}
func (p *songList) showAllSongs() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}
	for currentNode.nextSong != nil {
		fmt.Printf("%+v\n", *currentNode)
		currentNode = currentNode.nextSong
	}

	return nil
}
func main() {
	var songs = []song{
		song{title: "Kill This Love", artist: "BlankPink"},
		song{title: "Forever Young", artist: "BlankPink"},
		song{title: "Solo", artist: "Jennie"},
		song{title: "Quite", artist: "CLC"},
		song{title: "Paradox", artist: "CLC"},
		song{title: "One And Only", artist: "CLC"},
	}
	var playlist = songList{title: "Kpopin"}

	playlist.setBottomSonglist(&songs[0])
	playlist.setBottomSonglist(&songs[1])
	playlist.setBottomSonglist(&songs[3])
	playlist.setBottomSonglist(&songs[4])
	playlist.showAllSongs()
	// fmt.Println(playlist.head)
	// fmt.Println(playlist.head.nextSong)

}
