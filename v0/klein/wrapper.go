package klein

import "context"

type Wrapper struct {
	ctx     context.Context
	OnStart func() error
	OnStop  func() error
}

func (w *Wrapper) run() error {
	return w.OnStart()
}

func (w *Wrapper) stop() error {
	return w.OnStop()
}

func (w *Wrapper) setContext(ctx context.Context) {
	w.ctx = ctx
}
