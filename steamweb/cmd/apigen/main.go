package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"

	"github.com/13k/valve.go/steamweb/schema"
)

const (
	fmtUsage = `
Usage: apigen [options] <command> [command_options]

Commands:

filenames  Print a list of interfaces and respective filenames
generate   Generate files
clean      Remove all generated files
sandbox    Create a temporary "sandbox" to test generated files

Options:

%s

If any of the cached schema option is used and the file doesn't exist, the
schema will be fetched remotely and the file will be created. If the file
already exists, the schema will be loaded from it.

Filenames options:

%s

Sandbox

Creates a sandbox by linking all non-generated Go files and the go.mod file from the package into
the destination directory.

If the destination directory option is not given, it will create a temporary directory.

NOTE: must be run from the main package directory!

Sandbox options:

%s

`
)

var (
	cwd        string
	apiKey     string
	schemaFile string
	outputDir  string

	filenamesFlags       *flag.FlagSet
	filenamesOnlyMissing bool

	sandboxFlags *flag.FlagSet
	sandboxDir   string
)

func init() {
	var err error

	cwd, err = os.Getwd()

	if err != nil {
		panic(err)
	}

	flag.StringVar(&apiKey, "key", "", "API key to fetch API schema")
	flag.StringVar(&schemaFile, "cache", "", "Use file as cached API schema")
	flag.StringVar(&outputDir, "output", cwd, "Output directory")
	flag.StringVar(&outputDir, "o", cwd, "Output directory")

	filenamesFlags = flag.NewFlagSet("filenames", flag.ExitOnError)
	filenamesFlags.BoolVar(&filenamesOnlyMissing, "m", false, "Print only missing filenames")

	sandboxFlags = flag.NewFlagSet("sandbox", flag.ExitOnError)
	sandboxFlags.StringVar(&sandboxDir, "d", "", "Sandbox directory (defaults to new tempdir)")

	flag.Usage = usage
}

func usage() {
	var optionsBuf bytes.Buffer

	flag.CommandLine.SetOutput(&optionsBuf)
	flag.PrintDefaults()
	flag.CommandLine.SetOutput(os.Stderr)

	var filenamesBuf bytes.Buffer

	filenamesFlags.SetOutput(&filenamesBuf)
	filenamesFlags.PrintDefaults()
	filenamesFlags.SetOutput(os.Stderr)

	var sandboxBuf bytes.Buffer

	sandboxFlags.SetOutput(&sandboxBuf)
	sandboxFlags.PrintDefaults()
	sandboxFlags.SetOutput(os.Stderr)

	fmt.Fprintf(
		os.Stderr,
		fmtUsage,
		optionsBuf.String(),
		filenamesBuf.String(),
		sandboxBuf.String(),
	)
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(1)
	}

	log := logrus.New()
	log.SetOutput(os.Stderr)

	var cmd Command

	cmdStr := flag.Arg(0)

	switch cmdStr {
	case "filenames":
		if err := filenamesFlags.Parse(flag.Args()[1:]); err != nil {
			fatal(log, err)
		}

		cmd = &FilenamesCommand{
			OnlyMissing: filenamesOnlyMissing,
		}
	case "sandbox":
		if err := sandboxFlags.Parse(flag.Args()[1:]); err != nil {
			fatal(log, err)
		}

		sandboxCmd := &SandboxCommand{
			Directory: sandboxDir,
			Log:       log,
		}

		if err := sandboxCmd.Run(); err != nil {
			fatal(log, err)
		}

		return
	case "clean":
		cmd = &CleanCommand{
			OutputDir: outputDir,
			Log:       log,
		}
	case "generate":
		cmd = &GenerateCommand{
			OutputDir: outputDir,
			Log:       log,
		}
	case "help":
		usage()
		os.Exit(0)
	default:
		fmt.Fprintf(flag.CommandLine.Output(), "Invalid command %q\n", cmdStr)
		os.Exit(1)
	}

	schema, err := GetSchema(schemaFile, apiKey)

	if err != nil {
		fatal(log, err)
	}

	if err := cmd.Run(schema); err != nil {
		fatal(log, err)
	}
}

func fatal(log logrus.FieldLogger, err error) {
	l := log.WithError(err)

	{
		var e *errors.Error

		if errors.As(err, &e) {
			l.Errorln(e.ErrorStack())
		}
	}

	{
		var e *schema.InterfaceNotFoundError

		if errors.As(err, &e) {
			l.WithField("key", e.Key).Fatal()
		}
	}

	{
		var e *schema.InvalidInterfaceNameError

		if errors.As(err, &e) {
			l.WithField("interface", e.Name).Fatal()
		}
	}

	{
		var e *schema.MethodNotFoundError

		if errors.As(err, &e) {
			l.WithField("key", e.Key).Fatal()
		}
	}

	{
		var e *schema.InvalidMethodNameError

		if errors.As(err, &e) {
			l.WithFields(logrus.Fields{
				"method":  e.Name,
				"version": e.Version,
			}).Fatal()
		}
	}

	{
		var e *schema.InvalidMethodVersionError

		if errors.As(err, &e) {
			l.WithFields(logrus.Fields{
				"method":  e.Name,
				"version": e.Version,
			}).Fatal()
		}
	}

	{
		var e *schema.InvalidMethodHTTPMethodError

		if errors.As(err, &e) {
			l.WithFields(logrus.Fields{
				"method":      e.Name,
				"version":     e.Version,
				"http_method": e.HTTPMethod,
			}).Fatal()
		}
	}

	log.Fatal(err)
}
