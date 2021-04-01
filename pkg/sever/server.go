package sever

type Server interface {
	Start() (bool, error)
	Bind(int) (bool, error)
}

type serverListener interface {
	listen() error
}
