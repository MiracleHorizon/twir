package clients

import (
	"fmt"
	"log"

	"github.com/satont/tsuwari/libs/grpc/generated/timers"
	"github.com/satont/tsuwari/libs/grpc/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTimers(env string) timers.TimersClient {
	serverAddress := createClientAddr(env, "timers", servers.TIMERS_SERVER_PORT)

	conn, err := grpc.Dial(
		serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingConfig": [{"%s":{}}]}`, roundrobin.Name)),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := timers.NewTimersClient(conn)
	return c
}