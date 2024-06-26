package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/lmittmann/tint"
	"golang.org/x/term"
)

const version = "v0.0.1"

var (
	isTTY        = term.IsTerminal(int(os.Stdout.Fd()))
	printVersion = flag.Bool("version", false, "print version and exit")
	noColor      = flag.Bool("no-color", false, "disable color output")
)

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	p("")
	p("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	p("┃                                        ┃")
	p("┃                                        ┃")
	p("┃              put ascii art             ┃")
	p("┃                  here                  ┃")
	p("┃                                        ┃")
	p("┃                                        ┃")
	p("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ %s", version)
	p("")
	p("\033[32m•\033[0m starting `cli-name`")
	p("\033[32m•\033[0m press ctrl+c to stop")
	p("")

	logger := slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			NoColor:    !isTTY || *noColor,
		}),
	)

	logger.Info("ok", "version", version)
}

func p(format string, a ...any) {
	s := fmt.Sprintf(format, a...)

	if !isTTY || *noColor {
		s = strings.ReplaceAll(s, "\033[32m", "")
		s = strings.ReplaceAll(s, "\033[0m", "")
	}

	fmt.Fprintln(os.Stderr, s)
}
