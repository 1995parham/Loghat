/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 01-12-2017
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"flag"
	"fmt"

	"github.com/1995parham/Loghat/mw"
)

func main() {
	flag.Parse()
	m := mw.New("8159d5d5-7a05-4e4d-a1d5-001b81ceb618", "collegiate")

	for _, w := range flag.Args() {

		w, ss, err := m.Fetch(w)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, s := range ss {
			fmt.Printf("Did you mean %q ?\n", s)
		}

		for _, d := range w.Definitions {
			fmt.Printf("> %s\n", d.EWord)
			fmt.Printf("\t %s\n", d.Type)
			fmt.Printf("\t %s:\n", d.Date)
			for i, df := range d.Def {
				fmt.Printf("\t %d. %s\n", i, df)
			}
		}
	}
}
