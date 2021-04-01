package sever

import (
	"fmt"
	"github.com/wang1137095129/go-tcp/pkg/handle"
	"net"
)

const tcpServerNetwork = "tcp"

type ErrorHandler handle.Handler

type tcpErrorHandler struct {

}

func (t *tcpErrorHandler) Handle(obj interface{}) error {
	return nil
}

type tcpserver struct {
	port         int
	listener     net.Listener
	errorHandler ErrorHandler
}

func (s *tcpserver) Start() (bool, error) {
	listen, err := net.Listen(tcpServerNetwork, fmt.Sprintf(":%d", s.port))
	if err != nil {
		return false, err
	}
	s.listener = listen
	return true, err
}

func (s *tcpserver) Bind(port int) (bool, error) {
	s.port = port
	return true, nil
}

//func (s *tcpserver) listen() error {
//	conn, err := s.listener.Accept()
//	if err != nil {
//		s.errorHandler.Handle(err)
//	}
//
//	//conn.
//	return nil
//}
