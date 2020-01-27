package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	e1()
	e2()
	e3()
	e4()
}

// ---------------------------------------------------------
// EXERCISE: Convert the strings
//
//   1. Loop over the words slice
//
//   2. In the loop:
//      1. Convert each string value to a byte slice
//      2. Print the byte slice
//      3. Append the byte slice to the `bwords`
//
//   3. Print the words using the `bwords`
//
// EXPECTED OUTPUT
//  [103 111 112 104 101 114]
//  [112 114 111 103 114 97 109 109 101 114]
//  [103 111 32 108 97 110 103 117 97 103 101]
//  [103 111 32 115 116 97 110 100 97 114 100 32 108 105 98 114 97 114 121]
//  gopher
//  programmer
//  go language
//  go standard library
// ---------------------------------------------------------

func e1() {
	// Please uncomment the code below

	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	for _, v := range words {
		b := []byte(v)
		fmt.Println(b)
		bwords = append(bwords, b)
	}

	// fmt.Printf("% ", bwords)
	for _, v := range bwords {
		fmt.Println(string(v))
	}
}

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func e2() {
	const word = "console"
	// b := []byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}

	// for _, v := range word {
	// 	fmt.Printf("char %-5c dec %-5[1]d hex %-5[1]x v bin %-5[1]b \n", v)
	// 	b = append(b, byte(v))

	// }

	for _, w := range word {
		fmt.Printf("%c\n", w)
		fmt.Printf("\tdecimal: %[1]d\n", w)
		fmt.Printf("\thex    : 0x%[1]x\n", w)
		fmt.Printf("\tbinary : 0b%08[1]b\n", w)
	}

	fmt.Printf("with runes       : %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print the word manually using decimals
	fmt.Printf("with decimals    : %s\n",
		string([]byte{99, 111, 110, 115, 111, 108, 101}))

	fmt.Printf("with decimal     : %s\n", string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print the word manually using hexadecimals
	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))

	fmt.Printf("with hexadecimals: %s\n", []byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65})
}

var separator = "\n" + strings.Repeat("=", 20) + "\n"

func e4() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	// _ = strings

	// Print the byte and rune length of the strings
	// Hint: Use len and utf8.RuneCountInString
	fmt.Println(separator)
	for i, v := range strings {
		fmt.Printf("%d\t byte %v\t len(%d)\t rune len(%d)\n", i, v, len(v), utf8.RuneCountInString(v))
	}
	fmt.Println(separator)

	// Print the bytes of the strings in hexadecimal
	// Hint: Use % x verb
	for i := 0; i < len(strings); i++ {
		fmt.Printf(" %d\thex(% x)\n", i, strings[i])
	}
	fmt.Println(separator)

	// Print the runes of the strings in hexadecimal
	// Hint: Use % x verb
	for i, v := range strings {
		fmt.Printf("%d hex\n", i)
		for _, y := range v {
			fmt.Printf(" %x ", y)
		}
		fmt.Println("")
	}
	fmt.Println(separator)

	// Print the runes of the strings as rune literals
	// Hint: Use for range
	fmt.Println("Print the runes of the strings as rune literals")
	for _, v := range strings {
		// r := rune(byte(v))
		for _, y := range v {
			fmt.Printf(" %q ", y)
		}
		fmt.Println()

	}
	fmt.Println(separator)

	// Print the first rune and its byte size of the strings
	// Hint: Use utf8.DecodeRuneInString
	fmt.Println("Print the first rune and its byte size of the strings")
	for i, v := range strings {
		last, s := utf8.DecodeRuneInString(v)
		fmt.Printf("%d \t %v\t first %d\t size %d\n", i, v, last, s)

	}
	fmt.Println(separator)

	// Print the last rune of the strings
	// Hint: Use utf8.DecodeLastRuneInString
	fmt.Println("Print the last rune of the strings")
	for i, v := range strings {
		last, _ := utf8.DecodeLastRuneInString(v)
		fmt.Printf("%d \t %v\t last %d\n", i, v, last)

	}
	fmt.Println(separator)

	// Slice and print the first two runes of the strings
	fmt.Println("Slice and print the first two runes of the strings")
	for i, v := range strings {

		var firsttworunes []rune
		for _, r := range v {
			firsttworunes = append(firsttworunes, r)
		}
		firsttworunes = firsttworunes[:2]
		fmt.Printf("%d \t%c\n", i, firsttworunes)

	}
	fmt.Println(separator)

	// Slice and print the last two runes of the strings
	fmt.Println("Slice and print the last two runes of the strings")
	for i, v := range strings {
		var lasttworunes []rune
		for _, r := range v {
			lasttworunes = append(lasttworunes, r)
		}
		lasttworunes = lasttworunes[len(lasttworunes)-2 : len(lasttworunes)]
		fmt.Printf("%d \t%c\n", i, lasttworunes)

	}
	fmt.Println(separator)

	// Convert the string to []rune
	fmt.Println("Convert the string to []rune")
	for i, v := range strings {
		var runes []rune
		for _, r := range v {
			runes = append(runes, r)
		}
		fmt.Printf("%d \t %c\n", i, runes)
	}
	fmt.Println(separator)

	// Print the first and last two runes
	fmt.Println("Print the first and last two runes")
	for i, v := range strings {
		var firstLastTwo, temp []rune
		for _, r := range v {
			temp = append(temp, r)
		}
		firstLastTwo = append(firstLastTwo, temp[0:1]...)
		firstLastTwo = append(firstLastTwo, temp[len(temp)-2:len(temp)]...)
		//firstLastTwo = firstLastTwo[len(firstLastTwo)-2 : len(firstLastTwo)]
		fmt.Printf("%d \t%c\n", i, firstLastTwo)
	}
	fmt.Println(separator)

}

func e3() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range strings {
		fmt.Printf("%q\n", s)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n",
			len(s), utf8.RuneCountInString(s))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes   : % x\n", s)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
		fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(s)
		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))
	}
}
