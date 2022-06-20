package feishu

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Feishu struct {
	baseUrl string
}

type ErrResp struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewFeishu(baseUrl string) *Feishu {
	return &Feishu{baseUrl: baseUrl}
}

func (f *Feishu) Notify(msg string) error {
	tpl := `{
    "msg_type": "text",
    "content": {
        "text": "%s"
    }
}`
	feishuMsg := fmt.Sprintf(tpl, msg)
	resp, err := http.Post(f.baseUrl, "application/json", strings.NewReader(feishuMsg))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	if code <= 399 {
		log.Println("send success")
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var errResp ErrResp
	json.Unmarshal(body, &errResp)
	log.Printf("resp: %+v \n", errResp)

	return nil
}
