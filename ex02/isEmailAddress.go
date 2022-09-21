/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isEmailAddress.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/09/09 23:56:27 by wchen             #+#    #+#             */
/*   Updated: 2022/09/10 15:44:29 by wchen            ###   ########.fr       */
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

	if i == 1 {
		fmt.Printf("Invalid argument")
		return
	}
	re := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)
	for _, v := range os.Args[1:] {
		if re.MatchString(v) {
			fmt.Printf("%v is a valid email address\n", v)
		} else {
			fmt.Printf("%v is NOT a valid email address\n", v)
		}
	}
}
