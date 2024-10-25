package ports

type Serveable interface {
	Init() error
	Serve()
}
