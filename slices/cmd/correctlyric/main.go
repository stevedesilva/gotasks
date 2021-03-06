package main

import (
	"fmt"
	"strings"
	s "github.com/inancgumus/prettyslice"

)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice that contains the words of Beatles'
//  legendary song: Yesterday. However, the order of the
//  words are incorrect.
//
// CURRENT OUTPUT
//
//  [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
// STEPS
//
//  INITIAL SLICE:
//    [all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay]
//
//
//  2. Put the words to the correct positions in the `lyric` slice.
//
//     RESULT SHOULD BE:
//     [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
//  3. Print the `lyric` slice.
//
//
// BONUS
//
//   + Think about when does the append allocate a new backing array.
//
//   + Check whether your conclusions are correct.
//
//
// HINTS
//
//   If you get stuck, check out the hints.md file.
//
// ---------------------------------------------------------

func main() {
	// DON'T TOUCH THIS:
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)
	s.Show("lyric",lyric)

	yesterday := lyric[11:12:12]
	s.Show("yesterday",yesterday)

	allmytbls := lyric[:7:7]
	s.Show("allmytbls",allmytbls)

	update := append(yesterday,allmytbls...)

	now := append(lyric[12:cap(lyric)])
	s.Show("now",now)

	oh := append(lyric[7:12:12])
	s.Show("now",oh)

	update = append(update,now...)
	lyric = append(update,oh...)
	s.Show("lyric",lyric)
	

	// ADD YOUR CODE BELOW:
	// ...
	fmt.Printf("%s\n", lyric)


}
