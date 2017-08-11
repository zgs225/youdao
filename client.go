package youdao

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Youdao翻译的客户端

const (
	HOST        = "https://openapi.youdao.com/api"
	LAuto       = "auto"
	LChinese    = "zh-CHS"
	LJapanese   = "ja"
	LEnglish    = "en"
	LKorean     = "ko"
	LFrance     = "fr"
	LRussian    = "ru"
	LPortuguese = "pt" // 葡萄牙语
	LEspanol    = "es" // 西班牙语
)

var (
	ErrUnsupportLanguage = errors.New("不支持的语言类型")
	AllLanguages         = [9]string{LAuto, LChinese, LJapanese, LEnglish, LKorean, LFrance, LRussian, LPortuguese, LEspanol}
)

type Client struct {
	AppID     string
	AppSecret string

	from string
	to   string
}

func (c *Client) SetFrom(p string) error {
	if !checkLanguage(p) {
		return ErrUnsupportLanguage
	}
	c.from = p
	return nil
}

func (c *Client) GetFrom() string {
	if len(c.from) == 0 {
		return LAuto
	}
	return c.from
}

func (c *Client) SetTo(p string) error {
	if !checkLanguage(p) {
		return ErrUnsupportLanguage
	}
	c.to = p
	return nil
}

func (c *Client) GetTo() string {
	if len(c.to) == 0 {
		return LAuto
	}
	return c.to
}

func (c *Client) sign(q, salt string) string {
	h := md5.New()
	io.WriteString(h, c.AppID)
	io.WriteString(h, q)
	io.WriteString(h, salt)
	io.WriteString(h, c.AppSecret)
	b := h.Sum(nil)
	return hex.EncodeToString(b)
}

func (c *Client) Query(q string) (*Result, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resource, err := url.Parse(HOST)
	if err != nil {
		return nil, err
	}
	salt := randString(6)
	query := url.Values{}
	query.Set("q", q)

	from := "auto"
	if len(c.from) > 0 {
		from = c.from
	}
	to := "auto"
	if len(c.to) > 0 {
		to = c.to
	}
	query.Set("from", from)
	query.Set("to", to)
	query.Set("appKey", c.AppID)
	query.Set("salt", salt)
	query.Set("sign", c.sign(q, salt))
	resource.RawQuery = query.Encode()

	response, err := client.Get(resource.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var r Result
	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func checkLanguage(p string) bool {
	for _, v := range AllLanguages {
		if len(v) != len(p) {
			continue
		}
		if v == p {
			return true
		}
	}
	return false
}
