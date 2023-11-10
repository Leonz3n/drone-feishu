package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Body struct {
	Timestamp   string       `json:"timestamp"`
	Sign        *string      `json:"sign,omitempty"`
	MessageType string       `json:"msg_type,omitempty"`
	Content     *BodyContent `json:"content,omitempty"`
	Card        interface{}  `json:"card,omitempty"`
}

type BodyContent struct {
	Text     *string          `json:"text,omitempty"`
	Post     *BodyContentPost `json:"post,omitempty"`
	ImageKey *string          `json:"image_key,omitempty"`
}

type BodyContentPost struct {
	ZhCN *BodyContentPostLang `json:"zh_cn,omitempty"`
	EnUS *BodyContentPostLang `json:"en_us,omitempty"`
}

type BodyContentPostLang struct {
	Title   *string                                 `json:"title,omitempty"`
	Content [][]BodyContentPostLangContentParagraph `json:"content"`
}

type BodyContentPostLangContentParagraph struct {
	Tag      string  `json:"tag"`
	UnEscape *string `json:"un_escape,omitempty"`
	Text     *string `json:"text,omitempty"`
	Href     *string `json:"href,omitempty"`
	UserId   *string `json:"user_id,omitempty"`
	Username *string `json:"user_name,omitempty"`
	ImageKey *string `json:"image_key,omitempty"`
}

func NewBodyBuffer(timestamp int64, sign *string, arg Args) (*bytes.Buffer, error) {
	line1 := fmt.Sprintf("Repository: %s", arg.Repo.Name)
	line2 := fmt.Sprintf("Branch: %s", arg.Commit.Branch)
	line3 := fmt.Sprintf("Author: %s", arg.Commit.Author.Name)
	line4 := fmt.Sprintf("Commit: %s", arg.Commit.Rev)
	line5 := fmt.Sprintf("Message: %s", arg.Commit.Message)
	line6 := fmt.Sprintf("Status: %s", arg.Build.Status)

	body := &Body{
		Timestamp:   fmt.Sprintf("%d", timestamp),
		Sign:        sign,
		MessageType: "post",
		Content: &BodyContent{
			Post: &BodyContentPost{
				ZhCN: &BodyContentPostLang{
					Title: nil,
					Content: [][]BodyContentPostLangContentParagraph{
						{
							BodyContentPostLangContentParagraph{
								Tag:  "text",
								Text: &line1,
							},
						},
						{
							BodyContentPostLangContentParagraph{
								Tag:  "text",
								Text: &line2,
							},
						},
						{
							BodyContentPostLangContentParagraph{
								Tag:  "text",
								Text: &line3,
							},
						},
						{
							BodyContentPostLangContentParagraph{
								Tag:  "text",
								Text: &line4,
							},
						},
						{
							BodyContentPostLangContentParagraph{
								Tag:  "text",
								Text: &line5,
							},
						},
						{
							BodyContentPostLangContentParagraph{
								Tag:  "text",
								Text: &line6,
							},
						},
					},
				},
			},
		},
	}

	if buffer, err := json.Marshal(body); err == nil {
		return bytes.NewBuffer(buffer), nil
	} else {
		return nil, err
	}
}
