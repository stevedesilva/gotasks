package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	passme()
}

type user struct {
	username string
	password string
}

func passme() {

	const (
		errUsr = "Access denied for %q. \n"
		errPw  = "Invalid password for %q. \n"
		grant  = "Access granted to %q. \n"
		help   = "Usage: [username] [password]"
	)

	m := make(map[string]string)
	m["jack"] = "1888"
	m["steve"] = "1234"

	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l != 2 {
		fmt.Println(help)
		return
	}

	u, p := args[1], args[2]

	if i, ok := m[u]; !ok {
		fmt.Println(i, ok)
		fmt.Printf(errUsr, u)
	} else if p != i {
		fmt.Println(i, ok)

		fmt.Printf(errPw, u)
	} else {
		fmt.Println(i, ok)

		fmt.Printf(grant, u)
	}

}


