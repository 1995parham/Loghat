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

		w, err := m.Fetch(w)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, d := range w.Definitions {
			fmt.Println(d)
		}
	}
}
