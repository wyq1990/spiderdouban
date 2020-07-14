package parse_d

import (
	"douban/fetch"
	"douban/service"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ParseUrls(url string) {
	body := fetch.Fetch(url)
	body = strings.Replace(body, "\n", "", -1)

	rp := regexp.MustCompile(`<div class="item">[\s\S]*</div>`)
	titleRe := regexp.MustCompile(`alt="(.*?)" src`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)

	picRe := regexp.MustCompile(`https://img(\d+).doubanio.com/view/photo/s_ratio_poster/public/p(\d+).[A-Za-z]+`)
	contentRe := regexp.MustCompile(`<span class="inq">(.*?)</span>`)

	items := rp.FindAllStringSubmatch(body, -1)

	itemIds := idRe.FindAllStringSubmatch(items[0][0], -1)
	itemNames := titleRe.FindAllStringSubmatch(items[0][0], -1)
	itemPics := picRe.FindAllStringSubmatch(items[0][0], -1)
	itemContents := contentRe.FindAllStringSubmatch(items[0][0], -1)

	// pic := regexp.MustCompile(`<div class="pic">(.*?)</div>`)
	// for _, item := range items {
	// fmt.Println(idRe.FindStringSubmatch(item[1])[1],
	// 	titleRe.FindStringSubmatch(item[1])[1])

	
	itemIds = killRepetion(itemIds)
	fmt.Println(len(itemNames))
	fmt.Println(len(itemIds))
	fmt.Println(len(itemPics))
	fmt.Println(len(itemContents))

	for i := 0; i < len(itemIds); i++ {
		doubanId, err := strconv.Atoi(itemIds[i][1])
		name := itemNames[i][1]
		image := itemPics[i][0]
		content := itemContents[i][1]
		if err != nil {
			fmt.Println("没转成功")
		}

		service := service.InsertMovieService{
			DoubanId:     doubanId,
			MovieName:    name,
			MovieImg:     image,
			MovieContent: content,
		}
		service.Insert()
	}

	// }
}

func killRepetion(nums [][]string) [][]string {
	newRes := make([][]string, 0)
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if reflect.DeepEqual(nums[i], nums[j]) {
				flag = true
				break
			}
		}
		if !flag {
			newRes = append(newRes, nums[i])
		}
	}
	return newRes
}
