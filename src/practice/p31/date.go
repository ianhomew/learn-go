package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var now = time.Now()

	// time 的 Unix and UnixNano
	fmt.Println("Unix = ", now.Unix())
	fmt.Println("UnixNano = ", now.UnixNano())
	fmt.Println(math.MaxInt64) // now.UnixNano() 這返回是 int64 超過？

}
