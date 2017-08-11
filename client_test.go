package youdao

import (
	"encoding/json"
	"testing"
)

func BenchmarkSign(b *testing.B) {
	c := &Client{AppID: "hello", AppSecret: "world"}
	for i := 0; i < b.N; i++ {
		c.sign("你好", "3xhg3b")
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randString(6)
	}
}

func TestCH2JAClient(t *testing.T) {
	c := &Client{AppID: "2f871f8481e49b4c", AppSecret: "CQFItxl9hPXuQuVcQa5F2iPmZSbN0hYS"}
	if err := c.SetFrom(LChinese); err != nil {
		t.Error("SetFrom error: ", err)
		return
	}
	if err := c.SetTo(LJapanese); err != nil {
		t.Error("SetTo error: ", err)
		return
	}
	if v, err := c.Query("你好"); err != nil {
		t.Error("Query error: ", err)
		return
	} else {
		b, _ := json.Marshal(v)
		t.Logf("%s", b)
	}
}

func TestQuery(t *testing.T) {
	c := &Client{AppID: "2f871f8481e49b4c", AppSecret: "CQFItxl9hPXuQuVcQa5F2iPmZSbN0hYS"}
	if v, err := c.Query("hello"); err != nil {
		t.Error("Query error: ", err)
		return
	} else {
		b, _ := json.Marshal(v)
		t.Logf("%s", b)
	}
}
