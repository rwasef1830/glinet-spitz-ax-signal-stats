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

func (s *Svc) HandleRequest(*gserv.Context) gserv.Response {
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

	finalOutput := ""

	var commandOutputs []string
	commandOutputs = append(commandOutputs, execAtCommandAndGetResponse("AT+QCAINFO"))
	commandOutputs = append(commandOutputs, execAtCommandAndGetResponse("AT+QENG=\"servingcell\""))
	commandOutputs = append(commandOutputs, execAtCommandAndGetResponse("AT+QENG=\"neighbourcell\""))
	commandOutputs = append(commandOutputs, execAtCommandAndGetResponse("AT+QNWCFG=\"up/down\""))

	finalOutput = strings.Join(commandOutputs, "\n<br />\n")
	html = strings.Replace(html, "[placeholder]", finalOutput, 1)

	return gserv.PlainResponse("text/html", html)
}

func execAtCommandAndGetResponse(command string) string {
	return execCommandAndGetResponse("gl_modem", "-D", "AT", command)
}

func execCommandAndGetResponse(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		errorHtml := fmt.Sprint("Error:", err)
		return "<pre>" + errorHtml + "</pre>"
	} else {
		output := string(out.Bytes())
		return "<pre>" + output + "</pre>"
	}
}
