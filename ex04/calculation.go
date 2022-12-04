/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   calculation.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/01 22:39:36 by wchen             #+#    #+#             */
/*   Updated: 2022/12/04 10:47:01 by wchen            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(Args []string) (string, bool) {
	if len(Args) != 3 {
		return "", false
	}
	a, err_a := strconv.Atoi(Args[1])
	b, err_b := strconv.Atoi(Args[2])
	if err_a != nil || err_b != nil {
		return "", false
	}
	if b == 0 {
		return "", false
	}
	s := "sum: " + strconv.Itoa(a+b) + "\n" +
		"difference: " + strconv.Itoa(a-b) + "\n" +
		"product: " + strconv.Itoa(a*b) + "\n" +
		"quotient: " + strconv.Itoa(a/b) + "\n"
	return s, true
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
