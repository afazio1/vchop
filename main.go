package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		panic(e)
	}
}
func main() {
	// init flags
	var input string
	flag.StringVar(&input, "input", "", "the video file you want to chop")
	var d float64
	flag.Float64Var(&d, "dur", 2, "silence duration until notification")
	var noise int
	flag.IntVar(&noise, "noise", -30, "volume threshold for silence in dB")
	var output string
	flag.StringVar(&output, "output", "output.mp4", "the file to output to")

	flag.Parse()

	// fmt.Println("input:", input)
	// fmt.Println("duration:", d)
	// fmt.Println("noise:", noise)
	// fmt.Println("output:", output)
	
	cmd := exec.Command("ffmpeg", "-i", input, "-af", fmt.Sprintf("silencedetect=noise=%ddB:d=%f", noise, d), "-f", "null", "-")
	cmd.Dir = ""

	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(out.String())
		return
	}
	outStr := out.String()
	startTimes, endTimes, duration := parseOutput(outStr)
	
	command := constructCmd(&startTimes, &endTimes, duration, input, output)

	chopCmd := exec.Command(command[0], command[1:]...)
	chopCmd.Dir = ""
	chopCmd.Stdout = os.Stderr
	chopCmd.Stderr = os.Stderr
	chopErr := chopCmd.Run()
	check(chopErr)

}

func parseOutput(out string) ([]float64, []float64, float64) {
	duration := strings.Index(out, "Duration:")
	durationStr := out[duration:]
	durTokens := strings.Split(durationStr, " ")
	dur := strings.TrimSpace(strings.Replace(durTokens[1], ",", "", 1))
	
	times := strings.Split(dur, ":")
	hs, _ := strconv.ParseFloat(times[0], 64)
	ms, _ := strconv.ParseFloat(times[1], 64)
	ss, _ := strconv.ParseFloat(times[2], 64)

	totalSec := (hs * 3600) + (ms * 60) + ss

	firstSilence := strings.Index(out, "silencedetect")
	out = out[firstSilence:]
	tokens := strings.Split(out, " ")
	startTimes := make([]float64, 0)
	endTimes := make([]float64, 0)
	for i, token := range tokens {
		if token == "silence_start:" {
			sTime := strings.Split(tokens[i+1], "\n")[0]
			time, _ := strconv.ParseFloat(sTime, 64)
			startTimes = append(startTimes, time)
		} else if token == "silence_end:" {
			time, _ := strconv.ParseFloat(tokens[i+1], 64)
			endTimes = append(endTimes, time)
		}
	}
	return startTimes, endTimes, totalSec
}

func constructCmd(startTimes *[]float64, endTimes *[]float64, duration float64, inFile string, outFile string) []string {
	command := make([]string, 0)

	command = append(command,"ffmpeg", "-y", "-i", inFile, "-vf")

	var timeStamp float64
	timeStamp = 0
	vf := ""
	af := ""
	for i := 0; i < len(*startTimes); i++ { // first silence
		if i == 0 {
			addV := fmt.Sprintf("select='between(t,%.5f,%.5f)", timeStamp, (*startTimes)[i])
			vf += addV
			addA := fmt.Sprintf("aselect='between(t,%.5f,%.5f)", timeStamp, (*startTimes)[i])
			af += addA
		} else if i == len(*startTimes) - 1 { // last silence
			addV := fmt.Sprintf("+between(t,%.5f,%.5f)+between(t, %.5f, %.5f)',setpts=N/FRAME_RATE/TB", timeStamp, (*startTimes)[i], (*endTimes)[i], duration)
			vf += addV
			addA := fmt.Sprintf("+between(t,%.5f,%.5f)+between(t, %.5f, %.5f)',asetpts=N/SR/TB", timeStamp, (*startTimes)[i], (*endTimes)[i], duration)
			af += addA
		} else {
			addV := fmt.Sprintf("+between(t,%.5f,%.5f)", timeStamp, (*startTimes)[i])
			vf += addV
			af += addV
		}
		timeStamp = (*endTimes)[i]
	}
	command = append(command, vf, "-af", af, outFile)

	return command
}