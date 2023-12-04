package logger

import (
	"context"

	"farms-planning/pkg/runapp/fiber/middleware"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerFiber interface {
	Info(ctx context.Context, msg string)
}

type fiberLog struct{}

func NewLogger() LoggerFiber {
	return &fiberLog{}
}

type zeroLogFormat struct {
	RequestID string `json:"request_id"`
}

func (lf *zeroLogFormat) MarshalZerologObject(e *zerolog.Event) {
	e.Str("request_id", lf.RequestID)
}

func (*fiberLog) Info(ctx context.Context, msg string) {
	value := ctx.Value(middleware.CONTEXT_REQUEST_ID)
	if str, ok := value.(string); ok {
		var mylog = &zeroLogFormat{
			RequestID: string(str),
		}
		log.Info().EmbedObject(mylog).Msg(msg)
	}
}
