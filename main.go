package main

import (
	"os"
	"strconv"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("")

const usage = `Rhygo,

Bla-bla-bla...

Usage:
    rhygo [-t <tempo>] [-o <timeout>]

Options:
    -t <tempo>   Metronome tempo.
                 [default: 120]
    -o <timeout> Program stop timeout in seconds.
                 [default: 0]
`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "1.0", false)

	var (
		tempo   = args["-t"].(string)
		timeout = args["-o"].(string)
	)

	timeoutInt, _ := strconv.Atoi(timeout)

	if timeoutInt != 0 {
		go exitOnTimeout(timeoutInt)
	}

	tempoInt, _ := strconv.Atoi(tempo)
	runMetronome(tempoInt)
}

func runMetronome(tempo int) {
	for {
		//@TODO:we need sound here
		log.Notice("%v", "tick")

		time.Sleep(time.Second / (time.Duration(tempo) / 60))
	}
}

func exitOnTimeout(timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	os.Exit(0)
}
