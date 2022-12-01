/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   subSlice.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/12/01 22:39:28 by wchen             #+#    #+#             */
/*   Updated: 2022/12/01 22:39:29 by wchen            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func subSlice(slice []int, begin int, length int, capacity int) []int {
	ret := []int{}
	switch {
	case capacity < length:
		ret = make([]int, length, length)
	case capacity >= length:
		ret = make([]int, length, capacity)
	}
	for i := 0; i < length; i++ {
		if i < len(slice)-begin {
			ret[i] = slice[begin+i]
		} else {
			ret[i] = 0
		}
	}
	return (ret)
}

func main() {
	var orig = []int{0, 1, 2, 3, 4, 5}
	var ret []int
	ret = subSlice(orig, 0, 3, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}
