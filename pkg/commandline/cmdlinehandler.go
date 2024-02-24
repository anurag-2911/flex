package commandline

import (
	"flag"
	"fmt"
)

// This function displays and reads the commandline args
func ReadCommandLine() string {
	var applicationID string
	// applicationid is taken from command line or 374 is used as per the requirement doc.
	flag.StringVar(&applicationID, "appid", "374", "The application ID to calculate the minimum number of copies needed.")

	//help message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", flag.CommandLine.Name())
		flag.PrintDefaults()
	}
	//parse the flags
	flag.Parse()
	return applicationID
}
