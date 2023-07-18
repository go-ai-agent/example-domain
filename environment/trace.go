package environment

import (
	"context"
	"github.com/go-ai-agent/core/runtime"
	"net/url"
)

func GetTrace[E runtime.ErrorHandler](ctx context.Context, url *url.URL) ([]byte, *runtime.Status) {
	return nil, runtime.NewStatusOK()
}
