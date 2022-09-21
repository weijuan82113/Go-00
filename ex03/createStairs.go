/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   createStairs.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/09/11 11:57:59 by wchen             #+#    #+#             */
/*   Updated: 2022/09/12 07:38:59 by wchen            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input_str := os.Args[1]
	input_int, _ := strconv.Atoi(input_str)
	sum := 0

	for i := 1; sum+i <= input_int; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
		sum = sum + i
	}
}
