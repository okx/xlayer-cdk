package main

import (
	"os"

	zkevm "github.com/0xPolygon/cdk"
	"github.com/0xPolygon/cdk/config"
	"github.com/0xPolygon/cdk/log"
	"github.com/urfave/cli/v2"
)

const appName = "cdk"

const (
	// SEQUENCE_SENDER name to identify the sequence-sender component
	SEQUENCE_SENDER = "sequence-sender"
	// AGGREGATOR name to identify the aggregator component
	AGGREGATOR = "aggregator"
)

const (
	// NETWORK_CONFIGFILE name to identify the netowk_custom (genesis) config-file
	NETWORK_CONFIGFILE = "custom_network"
)

var (
	configFileFlag = cli.StringFlag{
		Name:     config.FlagCfg,
		Aliases:  []string{"c"},
		Usage:    "Configuration `FILE`",
		Required: true,
	}
	customNetworkFlag = cli.StringFlag{
		Name:     config.FlagCustomNetwork,
		Aliases:  []string{"net-file"},
		Usage:    "Load the network configuration file if --network=custom",
		Required: false,
	}
	yesFlag = cli.BoolFlag{
		Name:     config.FlagYes,
		Aliases:  []string{"y"},
		Usage:    "Automatically accepts any confirmation to execute the command",
		Required: false,
	}
	componentsFlag = cli.StringSliceFlag{
		Name:     config.FlagComponents,
		Aliases:  []string{"co"},
		Usage:    "List of components to run",
		Required: false,
		Value:    cli.NewStringSlice(SEQUENCE_SENDER, AGGREGATOR),
	}
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = zkevm.Version
	flags := []cli.Flag{
		&configFileFlag,
		&yesFlag,
		&componentsFlag,
	}
	app.Commands = []*cli.Command{
		{
			Name:    "version",
			Aliases: []string{},
			Usage:   "Application version and build",
			Action:  versionCmd,
		},
		{
			Name:    "run",
			Aliases: []string{},
			Usage:   "Run the cdk client",
			Action:  start,
			Flags:   append(flags, &customNetworkFlag),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}