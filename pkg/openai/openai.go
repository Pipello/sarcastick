package openai

import (
	"context"
	"fmt"
	"os"
	"strings"

	"sarcastick/pkg/news"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client openai.Client
}

func NewClient() *Client {
	return &Client{
		client: *openai.NewClient(os.Getenv("SARCASTICK_KEY")),
	}
}

func (c *Client) GetComment(n news.News) string {
	return ""
}

const roleSystemPrompt = `You are playing the role of a subversive, sarcastic journalist. Your job is to write an article in a newspaper.
The readers are expecting this article to be:
- in C1 English maximum
- 5 sentences maximum (up to 30 words)
- sarcastic / caricatural
- subversive / humoristic
`

const newsPromptF = `Here is the news of the day you have to write an article about: 
%s
`

func (c *Client) SarcasmHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	id, err := news.ParseID(update)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	n, err := news.GetNewsById(id)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	m, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Loading...",
	})
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4,
		MaxTokens: 150,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: roleSystemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf(newsPromptF, n.Content),
			},
		},
		Temperature: 1.1,
		Stream:      false,
	}
	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	var sb strings.Builder
	sb.WriteString("--- " + n.Title + " ---\n\n")
	sb.WriteString(resp.Choices[0].Message.Content)
	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: m.ID,
		Text:      sb.String(),
	})
}
