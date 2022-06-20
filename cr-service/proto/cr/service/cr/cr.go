package cr

import (
	"context"
	"github.com/jonathanmorais/endoenginnering/cr-service/pkg/cr"
	"github.com/jonathanmorais/endoenginnering/cr-service/pkg/logger"
)

type Server struct {
}

func (s *Server) CrHey(ctx context.Context, in *Message) *Response {
	log := logger.Instance()

	log.Info().Float64("CarboRatio", in.RatioCalc).Msg("Carbohydrate Ratio")

	return &Response{Value: int64(cr.RatioCalc())}
}
