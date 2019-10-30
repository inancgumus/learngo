// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.

Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.

Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	r, size := utf8.DecodeRuneInString("öykü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ykü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("kü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	// for range loop automatically decodes the runes
	//   but it gives you the position of the current rune
	//   instead of its size.

	// for _, r := range text {}
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)

		i += size
	}
	fmt.Println()
}
