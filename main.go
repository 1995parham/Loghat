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
)

func main() {
	flag.Parse()

	for _, w := range flag.Args() {
		fmt.Println(w)
	}
}
