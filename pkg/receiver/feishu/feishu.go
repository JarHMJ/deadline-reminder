package feishu

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Feishu struct {
	baseUrl string
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(body)
	return nil
}
