package main

import (
	"context"
	"fmt"
	"runtime/debug"
)

func main() {
	app, err := createApp()
	if err != nil {
		panic(fmt.Sprintf("Error creating app: %+v\n%s", err, debug.Stack()))
	}
	defer app.Tracer.Shutdown(context.Background())
	addr := app.Config.GetString("http.addr")
	panic(app.HTTPServer.Echo.Start(addr))
}
