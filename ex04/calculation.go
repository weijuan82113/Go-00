/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   calculation.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/01 22:39:36 by wchen             #+#    #+#             */
/*   Updated: 2022/12/01 22:39:38 by wchen            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(Args []string) ([]string, bool) {
	a, err_a := strconv.Atoi(Args[1])
	b, err_b := strconv.Atoi(Args[2])
	if len(Args) != 3 || err_a != nil || err_b != nil {
		return nil, false
	}
	s := make([]string, 4)
	s[0] = strconv.Itoa(a + b)
	s[1] = strconv.Itoa(a - b)
	s[2] = strconv.Itoa(a * b)
	s[3] = strconv.Itoa(a / b)
	return s, true
}

func main() {
	s := make([]string, 4)
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Printf("sum: %v\n", s[0])
	fmt.Printf("difference: %v\n", s[1])
	fmt.Printf("pruduct: %v\n", s[2])
	fmt.Printf("quotient: %v\n", s[3])
}
