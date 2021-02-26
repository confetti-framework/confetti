package commands

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"io"
	net "net/http"
	"strconv"
	"time"
)

type AppServe struct {
	Port int `short:"p" flag:"port"`
}

func (s AppServe) Name() string {
	return "app:serve"
}

func (s AppServe) Description() string {
	return "Start the http server to handle requests."
}

func (s AppServe) Handle(app inter.App, writer io.Writer) inter.ExitCode {
	name := app.Make("config.App.Name").(string)
	handler := app.Make((*net.HandlerFunc)(nil)).(func(net.ResponseWriter, *net.Request))

	_, _ = fmt.Fprintln(writer, "Start "+name+" to handle requests")
	server := &net.Server{
		Addr:         s.getPortAddr(app),
		Handler:      net.HandlerFunc(handler),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != net.ErrServerClosed {
		_, _ = fmt.Fprintln(writer, "Could not ", err)
		return inter.Failure
	}

	_, _ = fmt.Fprintln(writer, "Server stopped")

	return inter.Success
}

func (s AppServe) getPortAddr(app inter.App) string {
	var port int
	if s.Port != 0 {
		port = s.Port
	} else {
		port = app.Make("config.App.Port").(int)
	}
	return ":" + strconv.Itoa(port)
}
