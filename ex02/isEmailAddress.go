/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isEmailAddress.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/09/09 23:56:27 by wchen             #+#    #+#             */
/*   Updated: 2022/12/04 10:00:16 by wchen            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	i := len(os.Args)
	if i < 2 {
		fmt.Printf("Invalid argument\n")
		return
	}
	re := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)
	for _, v := range os.Args[1:] {
		if len(v) > 256 {
			fmt.Printf("%v is NOT a valid email address.\n", v)
		} else if re.MatchString(v) {
			fmt.Printf("%v is a valid email address.\n", v)
		} else {
			fmt.Printf("%v is NOT a valid email address.\n", v)
		}
	}
}
