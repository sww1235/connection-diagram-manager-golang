package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gorm.io/gorm"
)

var dbSQLite bool
var dbPostGres bool
var dbPostGresDSN string
var openFile string

var debugLogger = log.New(ioutil.Discard, "DEBUG: ", 0)
var infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var fatalLogger = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	// Init logging and command line stuff
	// optionally set up database connections
	// parse project config file
	// unmarshal any referenced library files
	// unmarshal any project yaml files
	// build master datasets from all unmarshalled files
	// launch optional TUI, GUI or command line
	// if command line is selected, perform selected commands
	// Ideas for operations:
	// Print complete connection diagram using graphviz/custom visualizer
	// Calculate conduit / cable tray fill
	// Print out layout diagram of individual or all locations

	var dbSL *gorm.DB
	var dbPG *gorm.DB
	var dbLocalSL *gorm.DB
	var err error
	err = initialization()
	if err != nil {
		fatalLogger.Panicln("Config and flag init failed", err)
	}
	// connect to databases
	if dbSQLite {

		if dbSL, err = connect("SQLite", ""); err != nil {
			fatalLogger.Panicln("Failed to initialize SQLite database", err)
		} else {
			// cannot close database connections created with GORM
			// in the normal way
			//	defer dbSL.Close()
		}
	}
	if dbPostGres {

		if dbPG, err = connect("PGSQL", dbPostGresDSN); err != nil {
			fatalLogger.Panicln("Failed to initialize PostGreSQL database", err)
		} else {
			// cannot close database connections created with GORM
			// in the normal way
			//	defer dbPG.Close()

		}
	}
	if openFile != "" {

		if dbLocalSL, err = connect("SQLite", path.Base(openFile)); err != nil {
			fatalLogger.Panicf("failed to open local project %s. Error: %s", openFile, err)
		} else {
			// cannot close database connections created with GORM
			// in the normal way
			//	defer dbLocalSL.Close()

		}
	}
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

	//project file is a folder containing a sqlLite db
	//and graphviz files, as well as documentation and
	//local config files
	// SQLite file:= {projectfile}/{projectfile}.sqlite
	flagOpenFile := flag.String("o", "", "open project `directory`")
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
	dbSQLite = *flagDBSqlLite
	dbPostGres = *flagDBPostGres
	dbPostGresDSN = *flagDBPostGresDSN
	openFile = *flagOpenFile

	//if dBSqlLite && dBPostGres {
	//	fatalLogger.Panicln("Multiple databases selected. This is not supported. Fix invocation")
	//}

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
