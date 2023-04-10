package youtube

import (
	"context"
	"encoding/json"
	"github.com/olahol/melody"
	"github.com/satont/tsuwari/apps/websockets/types"
	model "github.com/satont/tsuwari/libs/gomodels"
	"github.com/satont/tsuwari/libs/grpc/generated/websockets"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *YouTube) AddSongToQueue(_ context.Context, msg *websockets.YoutubeAddSongToQueueRequest) (*emptypb.Empty, error) {
	song := &model.RequestedSong{}
	err := c.services.Gorm.Where("id = ?", msg.EntityId).First(song).Error

	if err != nil {
		c.services.Logger.Error(err)
		return nil, err
	}

	message := &types.WebSocketMessage{
		EventName: "newTrack",
		Data:      song,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		c.services.Logger.Error(err)
		return nil, err
	}

	err = c.manager.BroadcastFilter(bytes, func(session *melody.Session) bool {
		userId, ok := session.Get("userId")
		return ok && userId.(string) == msg.ChannelId
	})

	if err != nil {
		c.services.Logger.Error(err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}