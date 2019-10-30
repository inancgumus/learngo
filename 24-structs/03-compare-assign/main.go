// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// #1b: create the song struct type
type song struct {
	title, artist string
}

// #5: structs can contain other structs
type playlist struct {
	genre string

	// songTitles []string
	// songArtist []string

	// #6: include a slice of song structs
	songs []song
}

func main() {
	// #1: create two struct values with the same type
	song1 := song{title: "wonderwall", artist: "oasis"}
	song2 := song{title: "super sonic", artist: "oasis"}

	fmt.Printf("song1: %+v\nsong2: %+v\n", song1, song2)

	// #4: structs are copied
	// song1 = song2

	// #3: structs can be compared
	if song1 == song2 {
		// #2: struct comparison works like this
		// if song1.title == song2.title &&
		// 	song1.artist == song2.artist {
		fmt.Println("songs are equal.")
	} else {
		fmt.Println("songs are not equal.")
	}

	// #8
	songs := []song{
		// #7b: you don't have to type the element types
		{title: "wonderwall", artist: "oasis"},
		{title: "supersonic", artist: "oasis"},
	}

	// #7: a struct can include another struct
	rock := playlist{
		genre: "indie rock",
		songs: songs,
	}

	// #9: you can't compare struct values that contains incomparable fields
	// you need to compare them manually

	// clone := rock
	// if rock.songs == clone {
	// }
	// if songs == songs {

	// #11: song is a clone, it cannot change the original struct value
	song := rock.songs[0]
	song.title = "live forever"

	// #11c: directly set the original one
	rock.songs[0].title = "live forever"

	// #11b
	fmt.Printf("\n%+v\n%+v\n", song, rock.songs[0])

	// #10: printing
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		// s := rock.songs[i]

		// #12b: s is a copy inside because struct values are copied
		s.title = "destroy"
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}

	// #12
	fmt.Printf("\n%-20s %20s\n", "TITLE", "ARTIST")
	for _, s := range rock.songs {
		fmt.Printf("%-20s %20s\n", s.title, s.artist)
	}
}
