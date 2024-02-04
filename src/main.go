package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/DTunnel0/CheckUser-Go/src/data/repository"
	"github.com/DTunnel0/CheckUser-Go/src/infra/factory"
	"github.com/DTunnel0/CheckUser-Go/src/infra/http"
)

var (
	version     = "0.0.1"
	version_str = fmt.Sprintf("checkuser %s", version)
	author      = "Glemison C. Dutra"
	email       = "glemyson20@gmail.com"
	description = fmt.Sprintf(
		"DTChecker - CHECKUSER | BY %s <%s> | VERSION: %s",
		author, email, version,
	)
)

type Args struct {
	Host           string
	Port           int
	Start          bool
	ListAllDevices bool
	ListDevices    string
	DeleteDevices  string
	DeleteDB       bool
	ShowVersion    bool
}

func initializeArgs() *Args {
	args := &Args{}

	flag.StringVar(&args.Host, "host", "0.0.0.0", "Host to listen")
	flag.IntVar(&args.Port, "port", 5000, "Port")
	flag.BoolVar(&args.Start, "start", false, "Start the daemon")

	flag.BoolVar(&args.ListAllDevices, "list-all-devices", false, "List all devices")
	flag.StringVar(&args.ListDevices, "list-devices", "", "List devices from a user")
	flag.StringVar(&args.DeleteDevices, "delete-devices", "", "Delete devices from a user")
	flag.BoolVar(&args.DeleteDB, "delete-db", false, "Delete database of devices")
	flag.BoolVar(&args.ShowVersion, "version", false, "Show version information")

	flag.Parse()
	return args
}

func printVersion() {
	fmt.Println(version_str)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n%s\n\n", description)
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage
	args := initializeArgs()

	if args.ShowVersion {
		printVersion()
		return
	}

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if args.Start {
		http.Start(args.Host, args.Port)
		return
	}

	if args.DeleteDB {
		repository.DeleteDB()
		return
	}

	if args.ListAllDevices {
		presenter := factory.MakeListDevicesPresenter()
		presenter.Present(context.Background())
		return
	}

	if args.ListDevices != "" {
		presenter := factory.MakeListDevicesByUsernamePresenter()
		presenter.Present(context.Background(), args.ListDevices)
		return
	}
}
