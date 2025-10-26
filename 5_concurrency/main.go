package main

import (
	"fmt"
	"time"
)

type Server struct {
	shutdown chan struct{} //empty struct channel takes 0 memory
}

func (s *Server) ListenAndServe() {
free: //this will give are mark to break outer loop from switch on line 18
	for {
		select {
		case <-s.shutdown:
			fmt.Println("attemt to gracefull shutdown")
			break free //if remove 'free', loop will be infinite
		default:
		}
	}
}

func (s *Server) quit() {
	close(s.shutdown)
}

func main() {
	srv := &Server{
		shutdown: make(chan struct{}),
	}

	go func() {
		time.Sleep(1 * time.Second)
		srv.quit()
	}()

	srv.ListenAndServe()
}
