package fun

import (
	"time"
	"sync"
	"context"
	"fmt"
	"errors"
	"net"
)

var (
	ErrServerStarted = errors.New("server already started")
)
type CallBacks interface {
	OnConnect()
	OnClose()
	OnError()
	OnMessage()
}

type DefaultCallBack struct {}

func (df *DefaultCallBack) OnConnect() {
	fmt.Println("OnConnect")
}

func (df *DefaultCallBack) OnClose() {
	fmt.Println("OnClose")
}

func (df *DefaultCallBack) OnError() {
	fmt.Println("OnError")
}

func (df *DefaultCallBack) OnMessage() {
	fmt.Println("OnMessage")
}

type Server struct {
	time time.Time
	callbacks CallBacks
	mu *sync.Mutex
	listeners map[net.Listener]bool
	wg *sync.WaitGroup
	connects *sync.Map
	ctx context.Context
	cancel context.CancelFunc
}

func NewServer() *Server {
	time := time.Now()
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	listeners := make(map[net.Listener]bool)
	connects := &sync.Map{}
	ctx, cancel := context.WithCancel(context.Background())

	return &Server{
		time:time,
		callbacks:new(DefaultCallBack),
		wg:wg,
		mu:mu,
		listeners:listeners,
		connects:connects,
		ctx:ctx,
		cancel:cancel,
	}
}

func (s *Server) Start(l net.Listener) error {

}