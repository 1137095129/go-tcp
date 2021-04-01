package handle

type Handler interface {
	Handle(interface{}) error
}

type name struct {
	
}

