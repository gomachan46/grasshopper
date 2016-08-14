package main

import (
	"fmt"

	"os"

	"log"

	contributions "github.com/gomachan46/gh-contributions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "do nothing\n")
		os.Exit(1)
	}

	args := os.Args[1:]
	res, errs := contributions.Get(args)
	fmt.Fprint(os.Stdout, "username,from,to,total,currentStreak\n")
	for _, c := range res {
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%d,%d\n", c.Username, c.From, c.To, c.Total, c.CurrentStreak)
	}
	if len(errs) != 0 {
		log.Fatalf("fail get contributions errs: %v", errs)
	}
}
