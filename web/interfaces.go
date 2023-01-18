package web

type Servable interface {
	Serve(port int, err chan string)
}
