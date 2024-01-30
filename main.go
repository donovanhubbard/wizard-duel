package main

import (
  "github.com/charmbracelet/log"
  "github.com/charmbracelet/ssh"
  "github.com/charmbracelet/wish/logging"
  "github.com/charmbracelet/wish"
  "fmt"
)

const (
	host = "localhost"
	port = 23234
)

func main(){
  log.SetLevel(log.DebugLevel)
  log.Infof("Starting program. Listening on %s:%d",host,port)

  s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			func(h ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					wish.Println(s, "Hello, world!")
					h(s)
				}
			},
			logging.Middleware(),
		),
	)

  if err != nil {
		log.Error("could not start server", "error", err)
	}

  log.Debug("Starting to listen")
  err = s.ListenAndServe()

  if err != nil {
    log.Error("Listening failed. Error:")
    log.Error(err)
  } else{
    log.Info("Server terminated")
  }
}
