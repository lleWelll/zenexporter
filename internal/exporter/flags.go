package exporter

import (
	"flag"
	"fmt"
)

const (
	outDirFlag = "outDir"
	usernameFlag = "u"
	passwordFlag = "p"

	usageHint = "Here should be usage hint"
)

func getParsedFlags() map[string]any {
	outDir := flag.String(outDirFlag, "", "output directory, where .csv file will be saved")
	username := flag.String(usernameFlag, "", "zenmoney username. Equivalent of func ZENMONEY_USERNAME environment variable")
	password := flag.String(passwordFlag, "", "zenmoney password. Equivalent of func ZENMONEY_PASSWORD environment variable")

	flag.Usage = usage
	flag.Parse()

	flags := map[string]any{}
	if outDir != nil {
		flags[outDirFlag] = *outDir
	}
	if username != nil {
		flags[usernameFlag] = *username
	}
	if password != nil {
		flags[passwordFlag] = *password
	}

	return flags
}

func usage() {
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), "%s\n", usageHint)
	flag.PrintDefaults()
}
