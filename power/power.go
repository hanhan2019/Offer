//给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
package main

import "fmt"

func main() {

}

func pow(base float64, exponent int) (float64, error) {
	var ok bool //指数是否小于0
	var res, curr float64 = 1, base
	if exponent < 0 {
		if base == 0 {
			return 0, fmt.Errorf("wrong")
		}
		ok = true
		exponent = -exponent
	} else if exponent == 0 {
		return 1, nil
	}
	for exponent != 0 {
		if exponent&1 == 1 {
			res *= curr
		}
		curr *= curr
		exponent = exponent >> 1
	}
	if ok {
		return float64(1) / res, nil
	} else {
		return res, nil
	}
}
