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
	"reflect"

	"gopkg.in/yaml.v3"
)

var dbPostGres bool
var dbPostGresDSN string

var projectDirectory string
var cfgData configuration
var noDefaultLibraries bool

var connectorTypeDict = make(map[string]ConnectorType)

//var equipmentTypeDict EquipmentTypes
//var pathwayTypeDict PathwayTypes
var wireTypeDict = make(map[string]WireType)
var cableTypeDict = make(map[string]CableType)
var termCableTypeDict = make(map[string]TermCableType)
var locationTypeDict = make(map[string]LocationType)

var debugLogger = log.New(ioutil.Discard, "DEBUG: ", 0)
var infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var fatalLogger = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

type configuration struct {
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
	//		fatalLo:gger.Panicln("Failed to initialize PostGreSQL database", err)
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
	debugLogger.Printf("%#+v\n", cfgData)

	err = readData()
	fmt.Println(len(termCableTypeDict))
	fmt.Printf("%+v\n", termCableTypeDict)

	if err != nil {
		fatalLogger.Panicln("Reading data errored", err)
	}
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

	if flag.NArg() > 1 {
		infoLogger.Println("Multiple project directories specified, only using first")
	}

	debugLogger.Println(flag.Arg(0))
	projectDirectory = flag.Arg(0) // first remaining cmd line arg after flag processing

	// check to see if projectDirectory was actually specified
	if projectDirectory == "" {
		//TODO: switch to custom error here
		return errors.New("No project directory specified")
	}

	// check to see if projectDirectory is a directory

	fileInfo, err := os.Stat(projectDirectory)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		// projectDirectory is not a directory
		return fmt.Errorf("Specified Project Directory %s is not a directory", projectDirectory)
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
			err = nil
		}

	} else {
		srcCfgExists = true
	}

	rootCfgFile, err = os.ReadFile(filepath.Join(projectDirectory, "cdm_config.yaml"))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			rootCfgExists = false
			err = nil
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

func readData() error {

	err := filepath.WalkDir(filepath.Join(projectDirectory, "src"), func(path string, d fs.DirEntry, err error) error {
		//reminder: inside walkFunc now
		if err != nil {
			return err
		}

		// walkDir loops through all files in projectDirectory recursively
		debugLogger.Printf("path: %s", path)

		if d.Name() == ".git" {
			return filepath.SkipDir
		}
		if d.Name() == ".github" {
			return filepath.SkipDir
		}
		if !d.IsDir() && filepath.Ext(path) == ".yaml" {
			err := ParseYamlFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil

}

func ParseYamlFile(path string) error {

	fileData := make(map[string]interface{})

	debugLogger.Printf("reading project data from: %s", path)
	fileContents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileContents, &fileData)

	for dictType, dict := range fileData {
		// switch on keys, which are top level dictionaries in yaml files

		debugLogger.Printf("key: %s, value: %v", dictType, dict)
		debugLogger.Printf("key: %s, type of value: %s", dictType, reflect.TypeOf(dict))

		switch dictType {
		// library dicts
		case "connector_type":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParseConnectorTypes(dictValue)
			if err != nil {
				return err
			}
		case "connector_type_mate":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParseConnectorTypeMates(dictValue)
			if err != nil {
				return err
			}
		case "equipment_type":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParseEquipmentTypes(dictValue)
			if err != nil {
				return err
			}
		case "pathway_type":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParsePathwayTypes(dictValue)
			if err != nil {
				return err
			}
		case "wire_type":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParseWireTypes(dictValue)
			if err != nil {
				return err
			}
		case "cable_type":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParseCableTypes(dictValue)
			if err != nil {
				return err
			}
		case "term_cable_type":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParseTermCableTypes(dictValue)
			if err != nil {
				return err
			}
		case "location_type":
			dictValue, ok := dict.(map[string]interface{})
			if !ok {
				debugLogger.Printf("not map. key: %s, type of value: %s", dictType, reflect.TypeOf(dict))
			}
			err := ParseLocationTypes(dictValue)
			if err != nil {
				return err
			}
		// instance dicts
		case "colors":

		case "equipment":

		case "wire_cable":

		case "pathway":

		case "location":

		case "connection":

		default:
			debugLogger.Printf("key:%s, value: %v", dictType, dict)

		}
	}
	return nil

}
