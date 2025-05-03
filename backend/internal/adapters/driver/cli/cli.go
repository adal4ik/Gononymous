package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	port = flag.String("port", "8080", "Port number.")
	help = flag.Bool("help", false, "Show this screen.")
	Port string
)

func init() {
	flag.Parse()

	if *help {
		fmt.Println(`gononymous - anonymous imageboard

Usage:
	gononymous [--port <N>]
	gononymous --help

Options:
	--help       Show this screen.
	--port N     Port number.`)
		os.Exit(0)
	}

	n, err := strconv.Atoi(*port)
	if err != nil || n <= 0 || n > 65535 {
		fmt.Fprintf(os.Stderr, "Error: invalid port number '%s'\n", *port)
		os.Exit(1)
	}

	Port = fmt.Sprintf(":%d", n)
}
