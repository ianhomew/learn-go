package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var now = time.Now()

	// time 的 Unix and UnixNano
	fmt.Println("Unix = ", now.Unix()) // this is timestamp
	fmt.Println("UnixNano = ", now.UnixNano())
	fmt.Println(math.MaxInt64) // now.UnixNano() 這返回是 int64 超過？

	// 補充拿到一個 timestamp 怎麼轉人類年月日時分秒
	var timestamp int64 = 1592236612
	var timeObj = time.Unix(timestamp, 0)
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))

}
