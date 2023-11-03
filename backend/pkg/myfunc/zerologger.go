package myfunc

import (
	"context"
	"farms-planning/pkg/mymiddleware"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type zeroLogFormat struct {
	RequestID string `json:"request_id"`
}

func (lf *zeroLogFormat) MarshalZerologObject(e *zerolog.Event) {
	e.Str("request_id", lf.RequestID)
}

func Logger(ctx context.Context, msg string) {
	value := ctx.Value(mymiddleware.CONTEXT_REQUEST_ID)
	if str, ok := value.(string); ok {
		var mylog = &zeroLogFormat{
			RequestID: string(str),
		}
		log.Info().EmbedObject(mylog).Msg(msg)
	}
}
