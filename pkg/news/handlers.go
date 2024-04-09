package news

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SummaryHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	t, err := getSummary()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   t,
	})
}

func DetailsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	id, err := ParseID(update)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	n, err := GetNewsById(id)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	var qb strings.Builder
	qb.WriteString(n.Title + "\n\n")
	qb.WriteString(n.Content)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   qb.String(),
	})
}

func ParseID(update *models.Update) (int, error) {
	text := update.Message.Text
	id := strings.Split(text, " ")[1]
	return strconv.Atoi(id)
}
