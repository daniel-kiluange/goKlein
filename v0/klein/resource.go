package klein

type Resource interface {
	run() error
	stop() error
}
