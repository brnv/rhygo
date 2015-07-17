package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/moriyoshi/pulsego/src/github.com/moriyoshi/pulsego"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("")

const usage = `Rhygo,

simple metronome with good clicking sound and timeout feature.

Usage:
    rhygo [-t <tempo>] [-s <strong_beat>] [-o <timeout>]

Options:
    -t <tempo>       Metronome tempo.
                     [default: 120]
    -s <strong_beat> Strong beat index.
                     [default: 0]
    -o <timeout>     Program stop timeout in seconds.
                     [default: 0]
`

func main() {
	args, _ := docopt.Parse(usage, nil, true, "1.0", false)

	var (
		tempo           = args["-t"].(string)
		timeout         = args["-o"].(string)
		strongBeatIndex = args["-s"].(string)
	)

	timeoutInt, _ := strconv.Atoi(timeout)
	if timeoutInt != 0 {
		go exitOnTimeout(timeoutInt)
	}

	strongBeatIndexInt, _ := strconv.Atoi(strongBeatIndex)

	tempoFloat, _ := strconv.ParseFloat(tempo, 64)

	runMetronome(tempoFloat, strongBeatIndexInt)
}

func runMetronome(tempo float64, strongBeatIndex int) {
	pulseAudio, pulseAudioContext := initPulseAudio()
	defer pulseAudio.Dispose()
	defer pulseAudioContext.Dispose()

	tick := 0

	for {
		sample := tickWeakSample

		if strongBeatIndex != 0 {
			if tick%strongBeatIndex == 0 {
				sample = tickStrongSample
				tick = 0
			}
		}

		go performTickSound(sample, pulseAudioContext)
		if strongBeatIndex != 0 {
			tick++
		}

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

func performTickSound(sample []int8, context *pulsego.PulseContext) {
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

	count := 0
	for {
		stream.Write(sample, pulsego.SEEK_RELATIVE)

		count++
		if count == 200 {
			break
		}
	}
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
