package handlers

import (
	"context"

	irc "github.com/gempir/go-twitch-irc/v3"
	"github.com/samber/lo"
	"github.com/satont/tsuwari/libs/grpc/generated/parser"
)

func (c *Handlers) handleCommand(msg irc.PrivateMessage, userBadges []string) {
	requestStruct := &parser.ProcessCommandRequest{
		Sender: &parser.Sender{
			Id:          msg.User.ID,
			Name:        msg.User.Name,
			DisplayName: msg.User.DisplayName,
			Badges:      userBadges,
		},
		Channel: &parser.Channel{
			Id:   msg.RoomID,
			Name: msg.Channel,
		},
		Message: &parser.Message{
			Id:   msg.ID,
			Text: msg.Message,
		},
	}

	res, err := c.parserGrpc.ProcessCommand(context.Background(), requestStruct)
	if err != nil {
		return
	}

	if res.KeepOrder != nil && !*res.KeepOrder {
		for _, v := range res.Responses {
			r := v
			go c.BotClient.SayWithRateLimiting(
				msg.Channel,
				r,
				lo.If(res.IsReply, lo.ToPtr(msg.ID)).Else(nil),
			)
		}
	} else {
		for _, r := range res.Responses {
			validateResposeErr := ValidateResponseSlashes(r)
			if validateResposeErr != nil {
				c.BotClient.SayWithRateLimiting(
					msg.Channel,
					validateResposeErr.Error(),
					nil,
				)
			} else {
				c.BotClient.SayWithRateLimiting(
					msg.Channel,
					r,
					lo.If(res.IsReply, &msg.ID).Else(nil),
				)
			}
		}
	}
}
