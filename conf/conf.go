package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// DNSRecord 用于构造 DNS 记录的 JSON 结构体
type DNSRecord struct {
	ZoneID  string `json:"ZoneID"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Proxied bool   `json:"proxied"`
}

type RecoidRes struct {
	Success bool `json:"success"`
	Errors  []struct {
		Message string `json:"message"`
	} `json:"errors"`
	Result []OneRes `json:"result"`
}

type OneRes struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Mata struct {
	Target string
	Main   DNSRecord
	Then   DNSRecord
}

type cf struct {
	ApiKey   string
	ZoneID   string
	BotToken string
	ChatID   string
	TgApiUrl string
	Corn     int64
	Mata     []Mata
}

var Config cf

func init() {
	// 读取配置文件
	file, err := os.Open("mata.json")
	if err != nil {
		fmt.Println("mata.json配置文件不存在")
		os.Exit(0)
	}
	defer file.Close()

	// 解码 JSON
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("mata.json配置文件错误")
		os.Exit(0)
	}
	Config.TgApiUrl = strings.TrimRight(Config.TgApiUrl, "/")
}
