// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const lineWidth = 40

func main() {
	text := `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.

Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.

Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	var lw int // line width

	for _, r := range text {
		fmt.Printf("%c", r)

		switch lw++; {
		case lw > lineWidth && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lw = 0
		}
	}
	fmt.Println()
}

// call it like: runeHandler(text)
func runeHandler(text string) {
	for i := 0; i < len(text); {
		r := rune(text[i])

		size := 1
		if r > utf8.RuneSelf {
			r, size = utf8.DecodeRuneInString(text[i:])
			// check out the other functions as well, play with them!
			//
			// for example (type these into the command-line):
			//   go doc utf8
			//   go doc utf8 EncodeRune
		}
		i += size

		fmt.Printf("%c", r)
	}
}

// call it like: byteHandler(text)
func byteHandler(text string) {
	for i := 0; i < len(text); i++ {
		fmt.Printf("%c", text[i])
	}
}
