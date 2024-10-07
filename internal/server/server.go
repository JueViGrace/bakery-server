package server

import (
	"github.com/JueViGrace/bakery-go/internal/api"
)

type FiberServer struct {
	api api.Api
}

func New() *FiberServer {
	return &FiberServer{
		api: api.New(),
	}
}

func (s *FiberServer) Init() (err error) {
	err = s.api.Init()
	if err != nil {
		return err
	}
	return
}
