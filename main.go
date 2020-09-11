package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jessevdk/go-flags"
)

var options struct {
	// program options
	Verbose    []bool   `short:"v" long:"verbose" description:"Show verbose information"`
	Player     string   `short:"p" long:"player" description:"The program to use for playback" default:"ffplay"`
	PlayerArgs []string `long:"player-args" description:"The arguments to pass to the player" default:"-autoexit" default:"-nodisp"`
	Libraries  []string `short:"l" long:"library" description:"The location of the music library" default:"."`
	GoodExt    []string `short:"e" long:"ext" descruption:"Known good extentions" default:".mp3" default:".m4a" default:".ogg" default:".wav" default:".flac" default:".wma"`

	// cloud options
	Server string `short:"s" long:"server" description:"The location of the BeetrootCloud gateway" default:"cloud.beetroot.app"`
	Name   string `short:"n" long:"name" description:"The name of the player publicly" default:"Beetroot CorePlayer"`
}

var verboseLevel int
var library []TrackInfo = make([]TrackInfo, 0)

func main() {
	// parse the command-line arguments
	_, err := flags.Parse(&options)
	if err != nil {
		return
	}
	verboseLevel = len(options.Verbose)

	// show intro message
	intro()

	// discover the library and play
	discoverLibrary()
	playLoop()
}

func pathAcceptor(path string) bool {
	for _, ext := range options.GoodExt {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}

	return false
}

func intro() {
	fmt.Println("!! BeetQueue CorePlayer")
	fmt.Println("!! Written by Raresh Nistor")
	fmt.Println("!!")
	fmt.Println("!! Using the router: " + options.Server)
	fmt.Println("!! Run -h for command line arguments")
	fmt.Println("!!")
	fmt.Println()
}

func discoverLibrary() {
	// discover the library
	for _, folder_ := range options.Libraries {
		folder := strings.Trim(folder_, "\" ")

		fmt.Print("Discovering tracks in " + folder + "...")
		walkFolders(&library, folder, pathAcceptor)
		fmt.Println(" Done.")
	}
	fmt.Printf("Found %d tracks in %d libraries.\n\n", len(library), len(options.Libraries))
}

func playLoop() {
	for true {
		for _, track := range library {
			fmt.Println("NOW PLAYING | " + track.Title + " - " + track.Artist)

			// play all the songs
			cmd := exec.Command(options.Player, append(options.PlayerArgs, track.Filename)...)
			err := cmd.Run()
			if err != nil {
				fmt.Fprintf(
					os.Stderr, "%s: playback didn't finish successfully: %s",
					track.Filename, err.Error())
			}
		}
	}
}
