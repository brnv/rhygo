package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/moriyoshi/pulsego/src/github.com/moriyoshi/pulsego"
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

	tempoFloat, _ := strconv.ParseFloat(tempo, 64)
	runMetronome(tempoFloat)
}

func runMetronome(tempo float64) {
	pulseAudio, pulseAudioContext := initPulseAudio()
	defer pulseAudio.Dispose()
	defer pulseAudioContext.Dispose()

	for {
		performTickSound(pulseAudioContext)

		sleepDuration, _ := time.ParseDuration(fmt.Sprintf(
			"%fs", 60/tempo,
		))
		time.Sleep(sleepDuration)
	}
}

func exitOnTimeout(timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	os.Exit(0)
}

func performTickSound(context *pulsego.PulseContext) {
	stream := context.NewStream(
		"default",
		&pulsego.PulseSampleSpec{
			// SAMPLE_S16LE works fine with int8 samples
			// 48000 is specific value
			Format: pulsego.SAMPLE_S16LE, Rate: 48000, Channels: 1,
		},
	)

	if stream == nil {
		log.Fatal("Failed to create pulseaudio stream")
	}
	defer stream.Dispose()
	stream.ConnectToSink()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		count := 0
		for {
			// magic comes here
			stream.Write(tickWeakSample, pulsego.SEEK_RELATIVE)
			// i dunno why
			count++
			if count == 200 { // same with 200
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
}

func initPulseAudio() (*pulsego.PulseMainLoop, *pulsego.PulseContext) {
	pulseaudio := pulsego.NewPulseMainLoop()
	if pulseaudio == nil {
		log.Fatal("Failed to create pulseaudio main loop")
	}

	pulseaudio.Start()

	context := pulseaudio.NewContext("default", 0)
	if context == nil {
		log.Fatal("Failed to create pulseaudio context")
	}

	return pulseaudio, context
}
