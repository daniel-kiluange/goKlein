package klein

import (
	"context"
	"log"
	"slices"
	"sync"
)

type Lifecycle struct {
	resources []Resource
	context   context.Context
	wg        sync.WaitGroup
}

func NewLifecycle(context context.Context) *Lifecycle {
	return &Lifecycle{
		resources: make([]Resource, 0),
		context:   context,
	}
}

func (l *Lifecycle) AppendToLifecycle(w *Wrapper) *Wrapper {
	w.setContext(context.WithValue(l.context, "resource", "wrapper"))
	l.resources = append(l.resources, w)
	return w
}

func (l *Lifecycle) run() error {
	for _, r := range l.resources {
		l.wg.Add(1)
		if err := r.run(); err != nil {
			log.Print(err)
		}
	}
	l.wg.Wait()
	return nil
}

func (l *Lifecycle) stop() error {
	log.Print("Trying to stop gracefully")
	slices.Reverse(l.resources)
	for _, r := range l.resources {
		if err := r.stop(); err != nil {
			log.Print(err)
		}
		l.wg.Done()
	}
	return nil
}
