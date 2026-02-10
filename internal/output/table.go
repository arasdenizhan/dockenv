package output

import (
	"fmt"

	"github.com/fatih/color"
)

type Severity string

const (
	Critical Severity = "CRITICAL"
	High     Severity = "HIGH"
	Medium   Severity = "MEDIUM"
	Low      Severity = "LOW"
)

var (
	criticalColor = color.New(color.FgHiRed, color.Bold)
	highColor     = color.New(color.FgRed)
	mediumColor   = color.New(color.FgYellow)
	lowColor      = color.New(color.FgCyan)
	infoColor     = color.New(color.FgWhite)
)

func PrintFinding(sev Severity, message string, env string) {

	var c *color.Color

	switch sev {
	case Critical:
		c = criticalColor
	case High:
		c = highColor
	case Medium:
		c = mediumColor
	case Low:
		c = lowColor
	default:
		c = infoColor
	}

	colorize(c, "[%s] ", sev)
	fmt.Printf("%s -> %s\n", message, env)
}

func PrintScore(score int) {

	var c *color.Color

	switch {
	case score >= 80:
		c = color.New(color.FgHiGreen, color.Bold)
	case score >= 50:
		c = color.New(color.FgYellow, color.Bold)
	case score >= 20:
		c = color.New(color.FgHiRed)
	default:
		c = color.New(color.BgRed, color.FgWhite, color.Bold)
	}

	fmt.Print("\nSecurity Score: ")
	colorize(c, "%d/100\n", score)
}

func colorize(c *color.Color, format string, a ...interface{}) {
	if NoColor {
		fmt.Printf(format, a...)
		return
	}
	c.Printf(format, a...)
}