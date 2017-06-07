package main

import (
	"fmt"
	"github.com/zgs225/youdao"
)

const (
	APPID     = "2f871f8481e49b4c"
	APPSECRET = "CQFItxl9hPXuQuVcQa5F2iPmZSbN0hYS"
)

func main() {
	c := &youdao.Client{APPID, APPSECRET}
	r, _ := c.Query("你好")
	fmt.Println(*r)
}
