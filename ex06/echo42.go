package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var p_chr_n bool
	var p_chr_s string

	flag.BoolVar(&p_chr_n, "n", false, "omit trailing newline")
	flag.StringVar(&p_chr_s, "s", "", "separator (default \" \")")
	flag.Parse()
	if p_chr_n {
		fmt.Printf("%v", strings.Join(flag.Args(), " "))
	}
	if p_chr_s != "" {
		for i := 0; i < len(flag.Args()); i++ {
			fmt.Printf("%v%v", flag.Arg(i), p_chr_s)
		}
		fmt.Printf("\n")
	}
}
