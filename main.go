package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"go.oneofone.dev/gserv"
)

func main() {
	// Define a variable to hold the port number
	var port int

	// Use flag.IntVar to define a command-line flag
	flag.IntVar(&port, "p", 8080, "Port number to listen on (default: 8080)")

	// Parse command-line arguments
	flag.Parse()

	if port > 65535 || port <= 0 {
		panic("Port must be between 1 and 65535")
	}

	start(port)
}

func start(port int) {
	srv := gserv.New()

	svc := Svc{}
	srv.GET("/", svc.HandleRequest)

	fmt.Printf("Listening on port %d. Press CTLR+C to exit...\n", port)
	log.Panic(srv.Run(context.Background(), "0.0.0.0:"+fmt.Sprintf("%d", port)))
}

type Svc struct {
	// some stuff here
}

func (s *Svc) HandleRequest(ctx *gserv.Context) gserv.Response {
	html := `
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="refresh" content="1" />
<title>Signal Statistics</title>
<body>
	[placeholder]
</body>
</html>`

	cmd := exec.Command("gl_modem", "-D", "AT", "AT+QCAINFO")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		errorHtml := fmt.Sprint("Error:", err)
		html = strings.Replace(html, "[placeholder]", "<pre>"+errorHtml+"</pre>", 1)
	} else {
		output := string(out.Bytes())
		html = strings.Replace(html, "[placeholder]", "<pre>"+output+"</pre>", 1)
	}

	return gserv.PlainResponse("text/html", html)
}
