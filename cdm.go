package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var dbPostGres bool
var dbPostGresDSN string

var projectDirectory string
var cfgData Configuration
var noDefaultLibraries bool

var debugLogger = log.New(ioutil.Discard, "DEBUG: ", 0)
var infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var fatalLogger = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

type Configuration struct {
	LibraryFiles       []string `yaml:"library_files"`
	NoDefaultLibraries bool     `yaml:"no_default_libraries"`
}

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

	//var dbPG *gorm.DB
	var err error
	// init logging and command line flags
	err = initialization()
	if err != nil {
		fatalLogger.Panicln("Logging configuration and  flag init failed", err)
	}
	// connect to database

	//if dbPostGres {

	//	if dbPG, err = connect("PGSQL", dbPostGresDSN); err != nil {
	//		fatalLogger.Panicln("Failed to initialize PostGreSQL database", err)
	//	} else {
	//		// cannot close database connections created with GORM
	//		// in the normal way
	//		//	defer dbPG.Close()

	//	}
	//}
	// https://zetcode.com/golang/yaml/
	// Parse project config file
	err = readConfig()
	if err != nil {
		fatalLogger.Panicln("Reading configuration errored", err)
	}
	debugLogger.Printf("%#v\n", cfgData)

}

//set up command line options, logging and configuration
func initialization() error {

	//Define and parse commandline flags here
	//Defaults are set in flags as appropriate

	flagDBPostGres := flag.Bool("p", false, "Use PostgreSQL database")
	flagDBPostGresDSN := flag.String("pgDSN", "", "Data Source Name (DSN) for PostGres database connection")
	flagDebugLogging := flag.Bool("V", false, "Show debug logs")
	flagQuiet := flag.Bool("Q", false, "Minimal Output")
	flagNoDefaultLibs := flag.Bool("D", false, "Do not use default libraries")
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
	dbPostGres = *flagDBPostGres
	dbPostGresDSN = *flagDBPostGresDSN
	noDefaultLibraries = *flagNoDefaultLibs

	if len(flag.Args()) > 1 {
		infoLogger.Println("Multiple project directories specified, only using first")
	}

	projectDirectory = flag.Arg(0) // first remaining cmd line arg after flag processing

	// check to see if projectDirectory was actually specified
	if projectDirectory == "" {
		//TODO: switch to custom error here
		return errors.New("No project directory specified")
		//fatalLogger.Fatalln("Please specify a project directory")
	}

	// check to see if projectDirectory is a valid directory

	if !fs.ValidPath(projectDirectory) {
		// projectDirectory is not a valid path
		return &fs.PathError{Op: "check", Path: projectDirectory, Err: fs.ErrInvalid}
	} else {
		fileSystem := os.DirFS(projectDirectory)
		fileInfo, err := fs.Stat(fileSystem, projectDirectory)
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			// projectDirectory is not a directory
			return fmt.Errorf("Specified Project Directory %s is not a directory", projectDirectory)
		}

	}

	return nil

}

// config file can either be in root or src directory,
// and must be named "cdm_config.yaml"
// config file in src directory has priority over one in root directory
func readConfig() error {
	var rootCfgExists bool
	var srcCfgExists bool
	var rootCfgFile []byte
	var srcCfgFile []byte

	srcCfgFile, err := os.ReadFile(filepath.Join(projectDirectory, "src", "cdm_config.yaml"))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			srcCfgExists = false
		}

	} else {
		srcCfgExists = true
	}

	rootCfgFile, err = os.ReadFile(filepath.Join(projectDirectory, "cdm_config.yaml"))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			rootCfgExists = false
		}

	} else {
		rootCfgExists = true
	}

	if srcCfgExists {
		debugLogger.Println("Using src config file")
		err = yaml.Unmarshal(srcCfgFile, &cfgData)
	} else if rootCfgExists {
		debugLogger.Println("Using root config file")
		err = yaml.Unmarshal(rootCfgFile, &cfgData)
	} else if rootCfgExists {
		debugLogger.Println("Using root config file")
	} else {
		infoLogger.Println("No config file detected and parsed")
	}

	if err != nil {
		return err
	}
	if noDefaultLibraries || cfgData.NoDefaultLibraries {
		debugLogger.Println("default libraries disabled")
	}
	return nil

}
