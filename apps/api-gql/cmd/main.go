package main

import (
	cfg "github.com/satont/twir/libs/config"
	"github.com/twirapp/twir/apps/api-gql/internal/gql"
	"github.com/twirapp/twir/apps/api-gql/internal/gql/directives"
	"github.com/twirapp/twir/apps/api-gql/internal/gql/resolvers"
	subscriptions_store "github.com/twirapp/twir/apps/api-gql/internal/gql/subscriptions-store"
	"github.com/twirapp/twir/apps/api-gql/internal/httpserver"
	"github.com/twirapp/twir/apps/api-gql/internal/minio"
	pubclicroutes "github.com/twirapp/twir/apps/api-gql/internal/routes/public"
	"github.com/twirapp/twir/apps/api-gql/internal/routes/webhooks"
	"github.com/twirapp/twir/apps/api-gql/internal/sessions"
	"github.com/twirapp/twir/apps/api-gql/internal/wsrouter"
	"github.com/twirapp/twir/libs/baseapp"
	buscore "github.com/twirapp/twir/libs/bus-core"
	commandscache "github.com/twirapp/twir/libs/cache/commands"
	keywordscacher "github.com/twirapp/twir/libs/cache/keywords"
	twitchcache "github.com/twirapp/twir/libs/cache/twitch"
	"github.com/twirapp/twir/libs/grpc/clients"
	"github.com/twirapp/twir/libs/grpc/events"
	"github.com/twirapp/twir/libs/grpc/tokens"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		baseapp.CreateBaseApp("api-gql"),
		fx.Provide(
			sessions.New,
			func(config cfg.Config) tokens.TokensClient {
				return clients.NewTokens(config.AppEnv)
			},
			func(config cfg.Config) events.EventsClient {
				return clients.NewEvents(config.AppEnv)
			},
			minio.New,
			twitchcache.New,
			commandscache.New,
			keywordscacher.New,
			buscore.NewNatsBusFx("api-gql"),
			fx.Annotate(
				wsrouter.NewNatsSubscription,
				fx.As(new(wsrouter.WsRouter)),
			),
			subscriptions_store.New,
			resolvers.New,
			directives.New,
			httpserver.New,
			gql.New,
		),
		fx.Invoke(
			pubclicroutes.New,
			webhooks.New,
		),
	).Run()
}
