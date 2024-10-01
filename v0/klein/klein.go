package klein

import (
	"context"
	"log"
)

type Klein struct {
	ctx       context.Context
	lifecycle *Lifecycle
}

func NewKlein() *Klein {
	ctx := context.Background()
	return &Klein{
		ctx:       ctx,
		lifecycle: NewLifecycle(ctx),
	}
}

func (k *Klein) Run() error {
	return k.run()
}

func (k *Klein) Provide(wrapper func(lc *Lifecycle) *Wrapper) {
	w := wrapper(k.lifecycle)
	w.ctx = k.ctx
	k.lifecycle.Append(w)
}

func (k *Klein) run() error {
	log.Print("Running Klein")
	return k.lifecycle.run()
}

func (k *Klein) stop() error {
	log.Print("Stopping Klein")
	return k.lifecycle.stop()
}

func (k *Klein) Stop() error {
	return k.stop()
}
