// Package colorog provides functions to make colorize your logs.
package colorog

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

type Palette struct {
	TimeColor    *color.Color
	SuccessColor *color.Color
	InfoColor    *color.Color
	MessageColor *color.Color
	ErrorColor   *color.Color
	FatalColor   *color.Color
	PanicColor   *color.Color
}

type ColoredLog struct {
	Palette
	ShowTime   bool
	TimeFormat string
}

func NewColoredLog(showTime bool, timeFormat string) *ColoredLog {
	coloredLog := &ColoredLog{
		Palette: Palette{
			TimeColor:    color.New(color.FgMagenta),
			SuccessColor: color.New(color.FgGreen),
			InfoColor:    color.New(color.FgCyan),
			MessageColor: color.New(color.FgWhite),
			ErrorColor:   color.New(color.FgYellow),
			FatalColor:   color.New(color.FgRed),
			PanicColor:   color.New(color.FgWhite, color.BgRed),
		},
		ShowTime:   showTime,
		TimeFormat: timeFormat,
	}

	// Change Log Format
	log.SetFlags(0)
	log.SetOutput(coloredLog)

	return coloredLog
}

func (cl *ColoredLog) Write(bytes []byte) (int, error) {
	timeStr := ""

	if cl.ShowTime {
		timeFormat := "2006-01-02 15:04:05"

		if len(cl.TimeFormat) > 0 {
			timeFormat = cl.TimeFormat
		}

		timeStr = cl.TimeColor.Sprint(
			time.Now().UTC().Format(timeFormat),
		)
	}

	messagePrefix := ""
	if len(timeStr) > 0 {
		messagePrefix = " "
	}

	return fmt.Printf("%v%v%v", timeStr, messagePrefix, string(bytes))
}

func (cl *ColoredLog) Success(v ...interface{}) {
	log.Print(cl.SuccessColor.Sprint(v...))
}

func (cl *ColoredLog) Successf(format string, v ...interface{}) {
	log.Print(cl.SuccessColor.Sprintf(format, v...))
}

func (cl *ColoredLog) Successln(v ...interface{}) {
	log.Println(cl.SuccessColor.Sprint(v...))
}

func (cl *ColoredLog) Error(v ...interface{}) {
	log.Print(cl.ErrorColor.Sprint(v...))
}

func (cl *ColoredLog) Errorf(format string, v ...interface{}) {
	log.Print(cl.ErrorColor.Sprintf(format, v...))
}

func (cl *ColoredLog) Errorln(v ...interface{}) {
	log.Println(cl.ErrorColor.Sprint(v...))
}

func (cl *ColoredLog) Info(v ...interface{}) {
	log.Print(cl.InfoColor.Sprint(v...))
}

func (cl *ColoredLog) Infof(format string, v ...interface{}) {
	log.Print(cl.InfoColor.Sprintf(format, v...))
}

func (cl *ColoredLog) Infoln(v ...interface{}) {
	log.Println(cl.InfoColor.Sprint(v...))
}

func (cl *ColoredLog) Message(v ...interface{}) {
	log.Print(cl.MessageColor.Sprint(v...))
}

func (cl *ColoredLog) Messagef(format string, v ...interface{}) {
	log.Print(cl.MessageColor.Sprintf(format, v...))
}

func (cl *ColoredLog) Messageln(v ...interface{}) {
	log.Println(cl.MessageColor.Sprint(v...))
}

func (cl *ColoredLog) Fatal(v ...interface{}) {
	log.Fatal(cl.FatalColor.Sprint(v...))
}

func (cl *ColoredLog) Fatalf(format string, v ...interface{}) {
	log.Fatal(cl.FatalColor.Sprintf(format, v...))
}

func (cl *ColoredLog) Fatalln(v ...interface{}) {
	log.Fatalln(cl.FatalColor.Sprint(v...))
}

func (cl *ColoredLog) Panic(v ...interface{}) {
	log.Panic(cl.PanicColor.Sprint(v...))
}

func (cl *ColoredLog) Panicf(format string, v ...interface{}) {
	log.Panic(cl.PanicColor.Sprintf(format, v...))
}

func (cl *ColoredLog) Panicln(v ...interface{}) {
	log.Panicln(cl.PanicColor.Sprint(v...))
}
