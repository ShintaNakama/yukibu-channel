package pb

import (
	"context"

	"go.uber.org/zap"
	"local.rpcs.packages/football"
)

type beetleServer struct {
	football.BeetleServer
	logger *zap.Logger
}

// ProvideBeetleServer returns constructed beetleServer
func ProvideBeetleServer(logger *zap.Logger) football.BeetleServer {
	return &beetleServer{logger: logger}
}

func (b *beetleServer) GetTeams(req *football.Empty, stream football.Beetle_GetTeamsServer) error {

	return nil
}

func (b *beetleServer) GetTeam(ctx context.Context, req *football.TeamRequest) (*football.Team, error) {

	t := &football.Team{
		Id:        "",
		Name:      "",
		ShortName: "",
		Tla:       "",
		CrestUrl:  "",
		Founded:   0,
		Venue:     "",
		Squad:     []*football.Player{},
	}

	return t, nil
}
