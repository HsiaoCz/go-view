package mian

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxCoon int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxCoon: 10,
		id:      "default",
		tls:     true,
	}
}

func withTLS(opts *Opts) {
	opts.tls = false
}

func withMaxConn(n int) OptFunc {
	return func(o *Opts) {
		o.maxCoon = n
	}
}

func withID(id string) OptFunc {
	return func(o *Opts) {
		o.id = id
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(withTLS, withMaxConn(99), withID("hello"))
	fmt.Printf("%+v\n", s)
}
