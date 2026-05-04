package main

import (
	"errors"
	"flag"
	"fmt"
	"image"

	"strconv"
	"strings"
	"time"

	"go.mau.fi/whatsmeow"
)

const InfoImageURL = "https://pomokit.github.io/pomodoro/information.png"
const WarningImageURL = "https://pomokit.github.io/pomodoro/warning.png"

const defaultDuration = 25 * time.Minute
const defaultBreakDuration = 5 * time.Minute

var WAclient *whatsmeow.Client

var ScreenShotStack []*image.RGBA

// var silence = flag.Bool("silence", false, "Don't ring bell after countdown")

//var simple = flag.Bool("simple", false, "Display simple countdown")

// var simple = flag.Bool("simple", true, "Display simple countdown")

/*
func init() {
	const usage = `Usage of pomodoro %s:

    pomodoro [options] [finish time]

Duration defaults to %d minutes. Finish may be a duration (e.g. "1h2m3s")
or a target time (e.g. "1:00pm" or "13:02:03"). Durations may be expressed
as integer minutes (e.g. "15") or time with units (e.g. "1m30s" or "90s").

Chimes system bell at the end of the timer, unless -silence is set.
`
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, getVersion(), int(defaultDuration/time.Minute))
		flag.PrintDefaults()
	}
	flag.Parse()
} */



func waitDuration(start time.Time) (finish time.Time, err error) {
	if flag.NArg() > 1 {
		err = errors.New("Too many args...")
		return
	}

	arg := flag.Arg(0)

	if arg == "" {
		return start.Add(defaultDuration), nil
	}

	// Do this first so less time passes
	if n, err := strconv.Atoi(arg); err == nil {
		d := time.Duration(n) * time.Minute
		return start.Add(d), nil
	}

	if d, err := time.ParseDuration(arg); err == nil {
		return start.Add(d), nil
	}

	for _, format := range []string{
		time.Kitchen, strings.ToLower(time.Kitchen), "15:04", "15:04:05",
	} {
		finish, err = time.Parse(format, arg)
		if err == nil {
			finish = time.Date(
				start.Year(), start.Month(), start.Day(),
				finish.Hour(), finish.Minute(),
				finish.Second(), finish.Nanosecond(),
				time.Local)
			if !finish.After(start) {
				finish = finish.AddDate(0, 0, 1)
			}
			return finish, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse as time or duration: %q", arg)
}

func waitBreakDuration(start time.Time) (finish time.Time, err error) {
	if flag.NArg() > 1 {
		err = errors.New("Too many args...")
		return
	}

	arg := flag.Arg(0)

	if arg == "" {
		return start.Add(defaultBreakDuration), nil
	}

	// Do this first so less time passes
	if n, err := strconv.Atoi(arg); err == nil {
		d := time.Duration(n) * time.Minute
		return start.Add(d), nil
	}

	if d, err := time.ParseDuration(arg); err == nil {
		return start.Add(d), nil
	}

	for _, format := range []string{
		time.Kitchen, strings.ToLower(time.Kitchen), "15:04", "15:04:05",
	} {
		finish, err = time.Parse(format, arg)
		if err == nil {
			finish = time.Date(
				start.Year(), start.Month(), start.Day(),
				finish.Hour(), finish.Minute(),
				finish.Second(), finish.Nanosecond(),
				time.Local)
			if !finish.After(start) {
				finish = finish.AddDate(0, 0, 1)
			}
			return finish, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse as time or duration: %q", arg)
}
