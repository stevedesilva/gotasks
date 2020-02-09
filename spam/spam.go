package spam

import (
	"errors"
)

// Spammer base type
type Spammer struct {
	Input string
}

// New constructor
func New() *Spammer {
	return &Spammer{}
}

// ErrSpammer erro
var ErrSpammer = errors.New("Spam error")

// Process entry
func (s *Spammer) Process(args []string) (string, error) {
	input := args[1:]
	if len(input) != 1 {
		return "", ErrSpammer
	}
	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = input[0]
		size = len(text)
		buf  = make([]byte, 0, size)
		in   bool
	)
	for i := 0; i < size; i++ {
		// check there is more room for a link
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			buf = append(buf, link...)
			i += nlink
		}
		// use byte buf instead of string to allocate once then change; since strings are immutatble string += text[i] will create a new string each time
		c := text[i]
		switch c {
		case ' ', '\t', '\n':
			in = false
		}
		if in {
			c = mask
		}
		buf = append(buf, c)
	}

	// set input
	s.Input = string(buf)
	return s.process()
}

func (s *Spammer) process() (string, error) {
	result := s.Input
	return result, nil
}
