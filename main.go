package main

import (
	"context"
	"log"
	"os"
	"regexp"
	"sarcastick/pkg/news"
	"sarcastick/pkg/openai"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var detailsRegex = regexp.MustCompile("^/details [0-9]*$")
var sarcasmRegex = regexp.MustCompile("^/sarcasm [0-9]*$")

func main() {
	b, err := bot.New(os.Getenv("TELEGRAM_BOT_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	client := openai.NewClient()
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, helpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, helpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/news", bot.MatchTypeExact, news.SummaryHandler)
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, detailsRegex, news.DetailsHandler)
	b.RegisterHandlerRegexp(bot.HandlerTypeMessageText, sarcasmRegex, client.SarcasmHandler)
	b.Start(context.Background())
}

func helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var sb strings.Builder
	sb.WriteString("- use /news to get the main titles of the day (given with their ID)\n")
	sb.WriteString("- use /details with the ID of the news to know more\n")
	sb.WriteString("- use /sarcasm with the ID of the news to generate a sarcastic comment on the news\n")
	sb.WriteString("- use /help to show this message \n")
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   sb.String(),
	})
}
