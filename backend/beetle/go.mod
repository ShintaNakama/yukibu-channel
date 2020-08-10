module github.com/ShintaNakama/yukibu-channel/backend/beetle

go 1.14

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.31.0

  local.rpcs.packages/football v1.0.0
)

replace local.rpcs.packages/football => ../rpcs/football/v1
