package main

import (
	"fmt"
	"unicode"
)

// import "fmt"

func main() {

	const text string = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	const maxWidth int = 40

	var lw int = 0

	for _, v := range text {

		fmt.Printf("%c", v)

		switch lw++; {
		case lw > maxWidth && v != '\n' && unicode.IsSpace(v):
			fmt.Println()
			fallthrough
		case v == '\n':
			lw = 0
		}

	}
	fmt.Println()
}
