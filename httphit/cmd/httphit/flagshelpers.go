package main

import (
	"errors"
	"strconv"
)

type number int

func toNumber(p *int) *number {

	return (*number)(p)
}

func (n *number) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	switch {
	case err != nil:
		err = custErr("parse error")
	case v <= 0:
		err = custErr("should be positive")
	}
	*n = number(v)
	return err
}

func (n *number) String() string {
	return strconv.Itoa(int(*n))
}

func custErr(s string) error {

	s += "\n"

	return errors.New(s)

}

// making some test
// func test() {
// 	var i int
// 	n1 := toNumber(&i)
// 	n1.Set("20")
// 	fmt.Println(i)
// 	fmt.Println(n1.String())
// 	//dont need to call string method
// 	fmt.Println(n1)
// 	//The types satisfy the Stringer interfaces. Println
// 	//will call the String method whenever it detects an input value is a Stringer type.
//}
