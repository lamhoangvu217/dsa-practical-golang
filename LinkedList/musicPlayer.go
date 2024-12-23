package LinkedList

import (
	"fmt"
	"strings"
)

type Song struct {
	id       string
	title    string
	artist   string
	next     *Song // Pointer to next song
	previous *Song // Pointer to previous song
}

type MusicPlaylist struct {
	currentSong *Song
}

// NewMusicPlaylist creates a new playlist
func NewMusicPlaylist() *MusicPlaylist {
	return &MusicPlaylist{
		currentSong: nil,
	}
}

// AddSong adds a new song to the end of the playlist
func (p *MusicPlaylist) AddSong(id, title, artist string) {
	//uuid, err := GenerateUUID()
	//if err != nil {
	//	log.Fatalf("Failed to generate UUID: %v", err)
	//}

	newSong := &Song{
		id:     id,
		title:  title,
		artist: artist,
	}

	// If playlist is empty, make this the first song
	if p.currentSong == nil {
		p.currentSong = newSong
		return
	}

	// Find the last song in the playlist
	current := p.currentSong
	for current.next != nil {
		current = current.next
	}

	// Link the new song
	current.next = newSong
	newSong.previous = current
}

func (p *MusicPlaylist) RemoveFirstSong() error {
	// check if the playlist is empty
	if p.currentSong == nil {
		return fmt.Errorf("cannot remove from an empty playlist")
	}
	// if we're currently on the first song, move currentSong to the next one
	// before removing the first song
	if p.currentSong.previous == nil {
		p.currentSong = p.currentSong.next
	}

	// find the first song (the one with no previous song)
	firstSong := p.currentSong
	for firstSong.previous != nil {
		firstSong = firstSong.previous
	}
	// If there's a next song, update its previous pointer to nil
	// since it will become the new first song
	if firstSong.next != nil {
		firstSong.next.previous = nil
	}
	firstSong.next = nil
	firstSong.previous = nil
	return nil
}

func (p *MusicPlaylist) RemoveLastSong() error {
	// check if the playlist is empty
	if p.currentSong == nil {
		return fmt.Errorf("cannot remove from an empty playlist")
	}

	// if we're currently on the last song, move currentSong to the previous one
	// before removing the last song
	if p.currentSong.next == nil {
		p.currentSong = p.currentSong.previous
	}

	// find the last song (the one with no next song)
	lastSong := p.currentSong
	for lastSong.next != nil {
		lastSong = lastSong.next
	}

	// If there's a previous song, update its next pointer to nil
	// since it will become the new last song
	if lastSong.previous != nil {
		lastSong.previous.next = nil
	}

	lastSong.next = nil
	lastSong.previous = nil
	return nil
}

func (p *MusicPlaylist) RemoveSpecificSong(songId string) error {
	// check if the playlist is empty
	if p.currentSong == nil {
		return fmt.Errorf("cannot remove from an empty playlist")
	}
	// find specific song in playlist
	// 1.find the first song in playlist
	firstSong := p.currentSong
	for firstSong.previous != nil {
		firstSong = firstSong.previous
	}
	// 2. Traverse through all songs starting from the first one
	current := firstSong
	for current != nil {
		if current.id == songId {
			if current.previous == nil { // Case 1: If it is first song
				if current.next != nil {
					current.next.previous = nil
				}
				if p.currentSong == current {
					p.currentSong = current.next
				}
			} else if current.next == nil { // Case 2: If it is last song
				if current.previous != nil {
					current.previous.next = nil
				}
				if p.currentSong == current {
					p.currentSong = current.previous
				}
			} else {
				// Case 3: If it is middle song
				current.previous.next = current.next
				current.next.previous = current.previous
				if p.currentSong == current {
					p.currentSong = current.next
				}
			}

			current.next = nil
			current.previous = nil
			return nil
		}
		current = current.next
	}
	return fmt.Errorf("song with ID %s not found", songId)

}

func (p *MusicPlaylist) PlayNext() string {
	if p.currentSong == nil || p.currentSong.next == nil {
		return "End of playlist"
	}

	p.currentSong = p.currentSong.next
	return fmt.Sprintf("Now playing: %s by %s", p.currentSong.title, p.currentSong.artist)
}

func (p *MusicPlaylist) PlayPrevious() string {
	if p.currentSong == nil || p.currentSong.previous == nil {
		return "You are in first song"
	}

	p.currentSong = p.currentSong.previous
	return fmt.Sprintf("Now playing: %s by %s", p.currentSong.title, p.currentSong.artist)
}

func (p *MusicPlaylist) ViewPlaylist() string {
	// Handle empty playlist case
	if p.currentSong == nil {
		return "Playlist is empty"
	}

	// Create a string builder for efficient string concatenation
	var result strings.Builder
	result.WriteString("=== Your Playlist ===\n")

	// First, find the first song in the playlist
	firstSong := p.currentSong
	for firstSong.previous != nil {
		firstSong = firstSong.previous
	}

	// Traverse through all songs starting from the first one
	current := firstSong
	position := 1
	for current != nil {
		// Add currently playing indicator if this is the current song
		nowPlaying := " "
		if current == p.currentSong {
			nowPlaying = "▶ " // Unicode play symbol
		}

		// Format and write this song's information
		songInfo := fmt.Sprintf("%s%d. %s %s - %s\n",
			nowPlaying,
			position,
			current.id,
			current.title,
			current.artist)
		result.WriteString(songInfo)

		current = current.next
		position++
	}

	result.WriteString("==================")
	return result.String()
}

func MusicPlayer() {
	playlist := NewMusicPlaylist()
	playlist.AddSong("1", "Hay trao cho anh", "MTP")
	playlist.AddSong("2", "Muon roi ma sao con", "MTP")
	playlist.AddSong("3", "Thien ly oi", "Jack - J97")
	playlist.AddSong("4", "Mong YU", "AMEE")

	//fmt.Println(playlist.PlayNext())
	//fmt.Println(playlist.PlayPrevious())

	err := playlist.RemoveSpecificSong("1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(playlist.ViewPlaylist())
}
