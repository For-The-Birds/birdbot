package tgbot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	//tb "github.com/tucnak/telebot"
	tb "gopkg.in/tucnak/telebot.v2"
)

type TgBot struct {
	bot *tb.Bot
	PhotoPath string
	birds, birds_dump tb.ChatID
}

var TBot TgBot

func (b * TgBot) AddPhoto(pid int) {
	newmsg := fmt.Sprintf("photo id %d", pid)
	msg, _ := b.bot.Send(b.bird1, newmsg, &tb.ReplyMarkup{
		InlineKeyboard: CreateBirdButtons(pid),
	})
	// sotre mid, cid := msg.MessageSig() ?
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func StartBot(ppath string) {

	// move to main
	var (
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
		listena   = os.Getenv("LISTEN_ADDR")
	)

	// poll?
	webhook := &tb.Webhook{
		Listen:   listena,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	checkerr(err)
	TBot.bot = b

	b.Handle("/start", func(m *tb.Message) {
		log.Println("got start from " + strconv.Itoa(m.Sender.ID))
		//id, _ := strconv.Itoa(m.Sender.ID)
		b.Send(m.Sender, "Hi", &tb.ReplyMarkup{
			ReplyKeyboardRemove: true,
		})
	})


	b.Handle(tb.OnCallback, func(c *tb.Callback) {
		log.Println("got inline " + c.Data)

		//switch classid

		//b.Edit(

		b.Respond(c, &tb.CallbackResponse{Text: "clicked " + c.Data})
	})

	log.Println("operating")

	b.Start()
}
