package main

import "github.com/daniel-kiluange/goKlein/v0/klein"

func newHttpServer(lc *klein.Lifecycle) *klein.Wrapper {
	return &klein.Wrapper{
		OnStart: func() error {
			println("Starting HTTP server")
			return nil
		},
		OnStop: func() error {
			println("Stopping HTTP server")
			return nil
		},
	}

}

func main() {
	k := klein.NewKlein()
	defer k.Stop()

	k.Provide(newHttpServer)

	k.Run()

}
