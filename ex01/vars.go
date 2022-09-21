/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   vars.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: wchen <wchen@42studen>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/09/03 14:29:00 by wchen             #+#    #+#             */
/*   Updated: 2022/09/04 00:35:32 by wchen            ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

type FortyTwo struct {
}

var (
	s1  string         = "42"
	s2  uint           = 42
	s3  int            = 42
	s4  uint8          = 42
	s5  uint16         = 42
	s6  uint32         = 42
	s7  uint64         = 42
	s8  bool           = false
	s9  float32        = 42
	s10 float64        = 42
	s11 complex64      = complex(42, 0)
	s12 complex128     = complex(42, 21)
	s13 FortyTwo       = FortyTwo{}
	s14 [1]int         = [1]int{42}
	s15 map[string]int = map[string]int{"42": 42}
	s16 *int
	s17 []int = []int{}
	s18 chan int
)

func main() {
	fmt.Printf("%v is %T\n", s1, s1)
	fmt.Printf("%v is %T\n", s2, s2)
	fmt.Printf("%v is %T\n", s3, s3)
	fmt.Printf("%v is %T\n", s4, s4)
	fmt.Printf("%v is %T\n", s5, s5)
	fmt.Printf("%v is %T\n", s6, s6)
	fmt.Printf("%v is %T\n", s7, s7)
	fmt.Printf("%v is %T\n", s8, s8)
	fmt.Printf("%v is %T\n", s9, s9)
	fmt.Printf("%v is %T\n", s10, s10)
	fmt.Printf("%v is %T\n", s11, s11)
	fmt.Printf("%v is %T\n", s12, s12)
	fmt.Printf("%v is %T\n", s13, s13)
	fmt.Printf("%v is %T\n", s14, s14)
	fmt.Printf("%v is %T\n", s15, s15)
	fmt.Printf("0x0 is %T\n", s16)
	fmt.Printf("%v is %T\n", s17, s17)
	fmt.Printf("%v is %T\n", s18, s18)
	fmt.Printf("%v is %T\n", nil, nil)
}
