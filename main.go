package main

import (
	"github.com/daniel-kiluange/goKlein/v0/klein"
	"net/http"
	"time"
)

func newHttpServer(lc *klein.Lifecycle) *klein.Wrapper {
	return &klein.Wrapper{
		OnStart: func() error {
			var err error
			httpMux := http.NewServeMux()
			println("Starting HTTP server")
			http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, World!"))
			})

			go func() {
				err = http.ListenAndServe(":8080", httpMux)
			}()
			return err
		},
		OnStop: func() error {
			println("Stopping HTTP server")
			return nil
		},
	}

}

func main() {
	k := klein.NewKlein()

	go func() {
		time.Sleep(5 * time.Second)
		k.Stop()
	}()

	k.Provide(newHttpServer)

	k.Run()

}
