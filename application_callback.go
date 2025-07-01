package tkit

import "context"

type CallbackPoint int32

const (
	CALL_LOAD_CONFIG CallbackPoint = iota + 1
	CALL_SETUP_VARS
	CALL_WITH_END
)

func (app *Application) loadConfigCallback(ctx context.Context) error {
	if f, ok := app.Callbacks[CALL_LOAD_CONFIG]; ok {
		return f(ctx)
	}

	return nil
}

func (app *Application) setupVarsCallback(ctx context.Context) error {
	if f, ok := app.Callbacks[CALL_SETUP_VARS]; ok {
		return f(ctx)
	}

	return nil
}

func (app *Application) endCallback(ctx context.Context) error {
	if f, ok := app.Callbacks[CALL_WITH_END]; ok {
		return f(ctx)
	}

	return nil
}
