package tkit

import "context"

type Runtime interface {
	Run(ctx context.Context) error

	setApplication(app *Application) error
	name() string
}

var (
	RuntimeName_HTTP  = "tkit.runtime.http"
	RuntimeName_Event = "tkit.runtime.event"
)
