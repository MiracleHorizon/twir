package internal

import (
	cfg "github.com/satont/tsuwari/libs/config"
	"github.com/satont/tsuwari/libs/grpc/generated/bots"
	"github.com/satont/tsuwari/libs/grpc/generated/tokens"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Services struct {
	DB     *gorm.DB
	Logger *zap.Logger
	Cfg    *cfg.Config

	BotsGrpc   bots.BotsClient
	TokensGrpc tokens.TokensClient
}

type Data struct {
	UserName        string `json:"userName,omitempty"`
	UserDisplayName string `json:"userDisplayName,omitempty"`

	RaidViewers int64 `json:"raidViewers,omitempty"`

	ResubMonths  int64  `json:"resubMonths"`
	ResubStreak  int64  `json:"resubStreak"`
	ResubMessage string `json:"resubMessage"`
	SubLevel     string `json:"subLevel"`

	OldStreamTitle    string `json:"oldStreamTitle"`
	NewStreamTitle    string `json:"newStreamTitle"`
	OldStreamCategory string `json:"oldStreamCategory"`
	NewStreamCategory string `json:"newStreamCategory"`

	StreamTitle    string `json:"streamTitle"`
	StreamCategory string `json:"streamCategory"`

	RewardID    string  `json:"-"'`
	RewardName  string  `json:"rewardName"`
	RewardCost  string  `json:"rewardCost"`
	RewardInput *string `json:"rewardInput"`

	CommandName string `json:"commandName"`
	CommandID   string `json:"-"`

	TargetUserName        string `json:"targetUserName"`
	TargetUserDisplayName string `json:"targetUserDisplayName"`

	DonateAmount   string `json:"donateAmount"`
	DonateMessage  string `json:"donateMessage"`
	DonateCurrency string `json:"donateCurrency"`

	PrevOperation *DataFromPrevOperation `json:"prevOperation"`
}

type DataFromPrevOperation struct {
	UnmodedUserName string `json:"unmodedUserName"`
	UnvipedUserName string `json:"unvipedUserName"`
	BannedUserName  string `json:"bannedUserName"`
}