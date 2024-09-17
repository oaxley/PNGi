package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ----- constants

// Program Version
const ProgramVersion = "1.0.0"

// default randomizer seed (prime number)
const DefaultRandomSeed = 7261046627

const kiloBytes = 1024
const megaBytes = kiloBytes * 1024
const gigaBytes = megaBytes * 1024

// ----- type
type Args struct {
	Randomize   bool
	Seed        int
	Secret      string
	BlockSize   int
	BlockUnit   int
	InPath      string
	OutPath     string
	PayloadPath string
	Command     string
}

var AppArgs Args
var blockStr string

// ----- functions

// specific usage functions
func usage() {
	programName := filepath.Base(os.Args[0])

	fmt.Printf("%s - PNG Injector - v%s\n", programName, ProgramVersion)
	fmt.Println("Syntax:")
	fmt.Printf("\t%s [options] <command>\n", programName)

	fmt.Println("\nOptions:")
	flag.PrintDefaults()

	fmt.Println("\nCommands:")
	fmt.Println("  insert\n\tInsert a payload inside a PNG or a list of PNGs")
	fmt.Println("  extract\n\tExtract a payload from a PNG or a list of PNGs")
	fmt.Println("  remove\n\tRemove a payload from a PNG or a list of PNGs")
	fmt.Println("  test\n\tTest if a payload is present in a PNG file")
	fmt.Println("  check\n\tCheck the integrity of the payload")
	fmt.Println("  show\n\tShow payload information")
}

func SetupFlags() {
	flag.BoolVar(&AppArgs.Randomize, "randomize", false, "Randomize payload blocks inside the PNG victim")
	flag.IntVar(&AppArgs.Seed, "seed", DefaultRandomSeed, "Seed for the payload randomizer")
	flag.StringVar(&AppArgs.Secret, "secret", "", "Shared secret to encrypt the payload")
	flag.StringVar(&blockStr, "block", "64K", "Size for each Payload chunks in the PNG victim")
	flag.StringVar(&AppArgs.InPath, "inPath", "./", "Path to a directory or file of source PNG files")
	flag.StringVar(&AppArgs.OutPath, "outPath", "./", "Path to a directory or file of target PNG files")
	flag.StringVar(&AppArgs.PayloadPath, "payload", "", "Path to the payload file to inject")

	flag.Usage = usage
}

func Parse() {
	// parse the command line
	flag.Parse()

	// ensure we have at least a command
	if flag.NArg() == 0 {
		log.Fatalln("Please specify a command.")
	}
	AppArgs.Command = strings.ToLower(flag.Args()[0])

	// setup the block unit parameters
	blockStr = strings.ToLower(blockStr)
	if strings.Contains(blockStr, "k") {
		AppArgs.BlockUnit = kiloBytes
	} else {
		if strings.Contains(blockStr, "m") {
			AppArgs.BlockUnit = megaBytes
		} else {
			if strings.Contains(blockStr, "g") {
				AppArgs.BlockUnit = gigaBytes
			} else {
				// check for byte only
				_, err := strconv.Atoi(blockStr)
				if err != nil {
					log.Fatalf("Unit '%c' unknown. Please use k, m or g.", blockStr[len(blockStr)-1])
				} else {
					AppArgs.BlockUnit = 1
				}
			}
		}
	}

	// setup the block size
	var blockSize string
	if AppArgs.BlockUnit > 1 {
		blockSize = blockStr[:len(blockStr)-1]
	} else {
		blockSize = blockStr
	}

	i, err := strconv.Atoi(blockSize)
	if err != nil {
		log.Fatal("Unable to convert ", blockSize, " in integer.")
	}
	AppArgs.BlockSize = i
}

func (args *Args) Debug() {
	fmt.Println("Randomize   : ", args.Randomize)
	fmt.Println("Seed        : ", args.Seed)
	fmt.Println("Secret      : ", args.Secret)
	fmt.Println("BlockSize   : ", args.BlockSize)
	fmt.Println("BlockUnit   : ", args.BlockUnit)
	fmt.Println("InPath      : ", args.InPath)
	fmt.Println("OutPath     : ", args.OutPath)
	fmt.Println("PayloadPath : ", args.PayloadPath)
	fmt.Println("Command     : ", args.Command)
}
