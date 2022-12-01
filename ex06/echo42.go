/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   echo42.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/01 22:39:15 by wchen             #+#    #+#             */
/*   Updated: 2022/12/01 22:39:16 by wchen            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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
	} else if p_chr_s != "" {
		var i int
		for i = 0; i < (len(flag.Args()) - 1); i++ {
			fmt.Printf("%v%v", flag.Arg(i), p_chr_s)
		}
		fmt.Printf("%v", flag.Arg(i))
		fmt.Printf("\n")
	} else {
		var i int
		for i = 0; i < (len(flag.Args()) - 1); i++ {
			fmt.Printf("%v ", flag.Arg(i))
		}
		fmt.Printf("%v", flag.Arg(i))
		fmt.Printf("\n")
	}
}
