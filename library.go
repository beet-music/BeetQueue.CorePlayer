package main

import (
	"errors"
	"os"
	"strings"

	"github.com/mikkyang/id3-go"
)

type TrackInfo struct {
	Filename, Title, Artist string
}

func TrackInfoFrom(file string) (trackInfo TrackInfo, err error) {
	trackInfo = TrackInfo{file, "", ""}
	err = nil

	// get file details
	stat, err := os.Stat(file)
	if err != nil {
		return
	}

	// reject if it's a dir
	if stat.IsDir() {
		return
	}

	// open the file to get details
	tags, id3Err := id3.Open(file)
	if id3Err == nil {
		trackInfo.Title = tags.Title()
		trackInfo.Artist = tags.Artist()
	} else {
		trackInfo.Title = stat.Name()
	}

	return
}

func walkFolders(library *[]string, path string, knownExt *[]string) error {
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
		// if the file is a folder, recurse
		if thisFile.IsDir() {
			err := walkFolders(library, path+"/"+thisFile.Name(), knownExt)
			if err != nil {
				return err
			}
		} else {
			// otherwise, check its extention and add it to the library
			for _, ext := range *knownExt {
				if !strings.HasSuffix(thisFile.Name(), ext) {
					continue
				}

				*library = append(*library, path+"/"+thisFile.Name())
				break
			}
		}
	}

	return nil
}
