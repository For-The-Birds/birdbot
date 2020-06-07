package tgbot

import (
	"encoding/json"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"strconv"
	"strings"
)

const (
	ButtonPost = iota
	ButtonClass
	ButtonOp
)

type ButtonCallbackData struct {
	PhotoID int `json:"i"`
	Class   int `json:"c"`
}

func GetCallbackData(pid, cls int) string {
	var b []byte
	var err error

	// TODO check Qty arr len
	b, err = json.Marshal(ButtonCallbackData{
		pid, cls,
	})

	log.Fatal(err)

	str := strings.TrimSpace(string(b))
	if len(str) > 64 {
		log.Fatal("too long callback data")
	}
	return str
}

func CreateBirdButtons(pid int) [][]tb.InlineButton {
	return [][]tb.InlineButton{
		[]tb.InlineButton{
			// add post to draft channel?
			tb.InlineButton{
				Unique: GetCallbackData(pid, 0),
				Data:   strconv.Itoa(pid),
				Text:   "Post",
			},
			tb.InlineButton{
				Unique: GetCallbackData(pid, 1),
				Data:   strconv.Itoa(pid),
				Text:   "1",
			},
			tb.InlineButton{
				Unique: GetCallbackData(pid, 2),
				Data:   strconv.Itoa(pid),
				Text:   "0",
			},
			tb.InlineButton{
				Unique: GetCallbackData(pid, 3),
				Data:   strconv.Itoa(pid),
				Text:   "T", // Tail fail
			},
			tb.InlineButton{
				Unique: GetCallbackData(pid, 4),
				Data:   strconv.Itoa(pid),
				Text:   "F", // Focus fail
			},
		},
	}
}

func ParseCallbackData(data string) (ButtonCallbackData, error) {
	d := strings.TrimSpace(data)
	var bcd ButtonCallbackData
	err := json.Unmarshal([]byte(d), &bcd)
	if err != nil {
		return ButtonCallbackData{}, err
	}
	return bcd, nil
}
