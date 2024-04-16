package actions

import "flower/engine"

type Action interface {
	Execute(ctx *engine.Context) error
}
