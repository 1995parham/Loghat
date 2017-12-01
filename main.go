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

	for _, w := range flag.Args() {
		m := mw.New(w)

		err := m.Fetch()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, d := range m.Definitions {
			fmt.Println(d)
		}
	}
}
