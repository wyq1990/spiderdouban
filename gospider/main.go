package main

import (
	"douban/conf"
	"douban/parse_d"
	"fmt"
	"strconv"
	"time"
)

func main() {

	conf.Init()

	start := time.Now()
	for i := 0; i < 9; i++ {
		parse_d.ParseUrls("https://movie.douban.com/top250?start=" + strconv.Itoa(25*i))
	}
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}
