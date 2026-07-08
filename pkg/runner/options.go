package runner

import (
	"flag"
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type Options struct {
	Verbose      bool
	RootList     string
	Debug        bool
	JsonOutput   bool
	CompactJson  bool
	WatchFile    bool
	OutputDir    string
	RabbitBroker string
	RabbitQueue  string
	ActorPID     *actor.PID
	ActorEngine  *actor.Engine
}

func ParseOptions() (*Options, error) {
	options := &Options{}

	flag.StringVar(&options.RootList, "r", "", "Path to the list of root domains to filter against (use '-' to read from stdin)")
	flag.BoolVar(&options.WatchFile, "f", false, "Monitor the root domain file for updates and restart the scan. requires the -r flag")
	flag.BoolVar(&options.Verbose, "v", false, "Output go logs (500/429 errors) to command line")
	flag.BoolVar(&options.Debug, "debug", false, "Debug CT logs to see if you are keeping up")
	flag.BoolVar(&options.JsonOutput, "j", false, "JSONL output cert info")
	flag.BoolVar(&options.CompactJson, "cj", true, "JSON matching my workflow")
	flag.StringVar(&options.OutputDir, "o", "", "Directory to store output files (one per hostname, requires -r flag)")
	flag.StringVar(&options.RabbitBroker, "rh", "", "RabbitMQ AMQP URI to publish domains to (e.g. amqp://guest:guest@localhost:5672/)")
	flag.StringVar(&options.RabbitQueue, "rq", "gungnir", "RabbitMQ queue to publish domains to")

	flag.Parse()

	// Validate that output directory is only used with root list
	if options.OutputDir != "" && options.RootList == "" {
		return nil, fmt.Errorf("the -o flag requires the -r flag to be set")
	}

	// The file watcher cannot monitor stdin
	if options.WatchFile && options.RootList == "-" {
		return nil, fmt.Errorf("the -f flag cannot be used when reading root domains from stdin")
	}

	return options, nil
}
