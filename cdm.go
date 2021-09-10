package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	//"path"
)

var dBSqlLite bool
var dBPostGres bool
var dBPostGresDSN string
var openFile string

var debugLogger = log.New(ioutil.Discard, "DEBUG: ", 0)
var infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var fatalLogger = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	err := initialization()
	if err != nil {
		fatalLogger.Panicln("Config and flag init failed", err)
	}
	// connect to databases
}

//set up command line options, logging and configuration
//read config file, either from -c flag or default in ~/.config
func initialization() error {

	//if this fails, cannot create default configDir. This is fatal
	//currUserConfigDir, err := os.UserConfigDir()
	//if err != nil {
	//	return err
	//}

	//default paths. Will not be overridden
	//defaultConfigPath := path.Join(currUserConfigDir, "cdm.cfg")

	//Define and parse commandline flags here
	//Defaults are set in flags as appropriate
	//flagConfigPath := flag.String("c", defaultConfigPath, "Path to config file")
	flagOpenFile := flag.String("o", "", "open project file")
	flagDBSqlLite := flag.Bool("s", false, "Use project specific SqlLite database")
	flagDBPostGres := flag.Bool("p", false, "Use PostgreSQL database")
	flagDBPostGresDSN := flag.String("pgDSN", "", "Data Source Name (DSN) for PostGres database connection")
	flagDebugLogging := flag.Bool("V", false, "Show debug logs")
	flagQuiet := flag.Bool("Q", false, "Minimal Output")
	//TODO: provide flags for different editors and viewing modes
	flag.Parse()

	if *flagDebugLogging && *flagQuiet {
		fatalLogger.Panicln("Both Quiet and Debug flags are set and are incompatible. Fix invocation")
	}

	if *flagDebugLogging {
		debugLogger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		debugLogger.SetOutput(os.Stdout)
	}

	if *flagQuiet {
		infoLogger.SetFlags(0)
		infoLogger.SetOutput(ioutil.Discard)
	}

	//retrieve values from flags and set global variables
	dBSqlLite = *flagDBSqlLite
	dBPostGres = *flagDBPostGres
	dBPostGresDSN = *flagDBPostGresDSN
	openFile = *flagOpenFile

	if dBSqlLite && dBPostGres {
		fatalLogger.Panicln("Multiple databases selected. This is not supported. Fix invocation")
	}

	//if *flagConfigPath != defaultConfigPath {
	//	infoLogger.Println("Using config file path from flag", *flagConfigPath)
	//}

	//Attempt to read config
	//tempConfig, cfgErr := readConfig(*flagConfigPath)
	//if cfgErr != nil {
	//	infoLogger.Printf("Config file not openable at path: %s, Err: %s "+
	//		"Using default configuration\n", *flagConfigPath, cfgErr)
	//} else {
	//	config = tempConfig
	//}

	return nil

}
