package tkit

import "context"

type Runtime interface {
	Run(ctx context.Context) error

	setApplication(app *Application) error
	name() string
}

var (
	ForgeRuntimeName_HTTP  = "forge.runtime.http"
	ForgeRuntimeName_Event = "forge.runtime.event"
)
