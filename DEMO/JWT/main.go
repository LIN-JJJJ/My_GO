package main

import (
	"fmt"
	"time"
)

func main() {
	// ----------------------------
	// HS256

	token, err := generateTokenUsingHs256()
	if err != nil {
		panic(err)
	}
	fmt.Println("Token = ", token)

	time.Sleep(time.Second * 1)

	my_claim, err := parseTokenHs256(token)
	if err != nil {
		panic(err)
	}
	fmt.Println("my claim = ", my_claim)
	// ----------------------------
	// RS256
	tokenRS256, err := generateTokenUsingRS256()
	if err != nil {
		panic(err)
	}
	fmt.Println("Token = ", tokenRS256)

	time.Sleep(time.Second * 1)

	my_claimRS256, err := parseTokenRs256(tokenRS256)
	if err != nil {
		panic(err)
	}
	fmt.Println("my claim = ", my_claimRS256)

}
