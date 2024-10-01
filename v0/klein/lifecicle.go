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

func (l *Lifecycle) Append(r Resource) {
	l.resources = append(l.resources, r)
}

func (l *Lifecycle) run() error {
	for _, r := range l.resources {
		if err := r.run(); err != nil {
			log.Print(err)
		}
		l.wg.Add(1)
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
