package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"os"

	"github.com/dhowden/tag"
)

/*
TrackInfo contains details about the track
*/
type TrackInfo struct {
	Filename string
	Title    string
	Artist   string
}

func hashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

/*
UniqueID returns a unique identifier for this track.
*/
func (trackInfo *TrackInfo) UniqueID() uint32 {
	return hashString(
		trackInfo.Filename +
			":" + trackInfo.Artist +
			":" + trackInfo.Title)
}

/*
TrackInfoFrom generates a track record from a file.
*/
func TrackInfoFrom(path string) (trackInfo TrackInfo, err error) {
	trackInfo = TrackInfo{path, "", ""}
	err = nil

	// get file details
	stat, err := os.Stat(path)
	if err != nil {
		return
	}

	// reject if it's a dir
	if stat.IsDir() {
		return
	}

	// open the file
	file, err := os.Open(path)
	if err != nil {
		return
	}

	// open the file to get details
	meta, err := tag.ReadFrom(file)
	if err != nil {
		if verboseLevel > 1 {
			fmt.Fprintf(
				os.Stderr,
				"%s: failed to read metadata (%s)\n",
				path, err.Error())
		}

		trackInfo.Title = stat.Name()
		err = nil
	} else {
		trackInfo.Title = meta.Title()
		trackInfo.Artist = meta.Artist()
	}

	return
}

type isAcceptable func(path string) bool

func walkFolders(
	library *[]TrackInfo,
	path string,
	isAcceptable isAcceptable) error {
	// TODO: rewrite this function to follow an iterative approch, that way we
	// won't run the risk of

	// get the file info for this file
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// check whether the file is a folder
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return errors.New("the path given must be a folder/directory")
	}

	// get the files in the folder
	files, err := file.Readdir(0)
	if err != nil {
		return err
	}

	for _, thisFile := range files {
		thisPath := path + "/" + thisFile.Name()

		// if the file is a folder, recurse
		if thisFile.IsDir() {
			err := walkFolders(library, thisPath, isAcceptable)
			if err != nil {
				return err
			}
		} else {
			// check if the file is acceptable and skip if not
			if !isAcceptable(thisPath) {
				if verboseLevel > 1 {
					fmt.Fprintf(os.Stderr, "\"%s\" is unacceptable.\n", thisPath)
				}
				continue
			}

			// get the metadata for this file and add it to the library
			thisTrack, err := TrackInfoFrom(thisPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: major error: %s\n", thisPath, err.Error())
				continue
			}

			*library = append(*library, thisTrack)
		}
	}

	return nil
}
