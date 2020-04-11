package discord

import "haracho-go/internal/logger"

type CommandContext struct {
	client  *Client
	channel string
	log     *logger.BasicLogger
}

func (c *CommandContext) SendMessage(message string) {
	_, err := c.client.session.ChannelMessageSend(c.channel, message)
	if err != nil {
		c.log.Infof("メッセージ送信失敗 \nChannelID: %v\nMessage: %v", c.channel, message)
	}
}
