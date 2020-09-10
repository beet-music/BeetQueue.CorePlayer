package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

var options struct {
	// program options
	Verbose   []bool   `short:"v" long:"verbose" description:"Show verbose information"`
	Player    string   `short:"p" long:"player" description:"The program to use for playback" default:"ffplay"`
	Libraries []string `short:"l" long:"library" description:"The location of the music library" default:"."`
	GoodExt   []string `short:"e" long:"ext" descruption:"Known good extentions" default:".mp3" default:".m4a" default:".ogg" default:".wav" default:".flac"`

	// cloud options
	Server string `short:"s" long:"server" description:"The location of the BeetrootCloud gateway" default:"cloud.beetroot.app"`
	Name   string `short:"n" long:"name" description:"The name of the player publicly" default:"Beetroot CorePlayer"`
}

var verboseLevel int
var library []string = make([]string, 0)

func main() {
	// parse the command-line arguments
	_, err := flags.Parse(&options)
	if err != nil {
		return
	}
	verboseLevel = len(options.Verbose)

	// discover the library
	fmt.Print("Discovering tracks... ")
	for _, folder := range options.Libraries {
		walkFolders(&library, folder, &options.GoodExt)
	}
	fmt.Printf("Found %d tracks in %d libraries.\n", len(library), len(options.Libraries))

	//
}
