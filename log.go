package log

import (
	"os"

	"github.com/fatih/color"
)

const (
	//Succ success
	levelSucc = iota
	//Err error
	levelErro
	//Trac trace
	levelTrac
	//Info infomation
	levelInfo
	//Warn warning
	levelWarn
)

func base(level int, msg ...interface{}) {
	switch level {
	case levelTrac:
		color.New(color.FgHiCyan).Println(msg...)
	case levelInfo:
		color.New(color.FgHiBlue).Println(msg...)
	case levelWarn:
		color.New(color.FgHiYellow).Println(msg...)
	case levelSucc:
		color.New(color.FgHiGreen).Println(msg...)
	case levelErro:
		color.New(color.FgRed).Println(msg...)
		os.Exit(1)
	}
}

//Trac trace
func Trac(msg ...interface{}) {
	base(levelTrac, msg...)
}

//Info infomation
func Info(msg ...interface{}) {
	base(levelInfo, msg...)
}

//Warn warnning
func Warn(msg ...interface{}) {
	base(levelWarn, msg...)
}

//Erro error
func Erro(msg ...interface{}) {
	for _, s := range msg {
		base(levelErro, s)
	}
}

//Succ succwss
func Succ(msg ...interface{}) {
	base(levelSucc, msg...)
}
