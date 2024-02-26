package commandline

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Gets the applicationID and filepath of the csv file
// applicationID if not given by the user is 374
// filepath is a mandatory parameter
func GetCommandLineArguments() (string, string) {
	var applicationID string
	var filePath string
	flag.StringVar(&applicationID, "appid", "374", "The application ID to calculate the minimum number of copies needed.")
	flag.StringVar(&filePath, "filepath", "", "Provide the path of the csv file")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", flag.CommandLine.Name())
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check if filePath is provided; if not, print usage and exit the program.
	if filePath == "" {
		log.Println("Error: The path to the csv file is required.")
		flag.Usage()
		os.Exit(1) // Exit the program indicating an error.
	}

	return applicationID, filePath
}
