package main

import (
	"fmt"
	"github.com/AGCunningham/mk-gen/selector"
	"github.com/AGCunningham/mk-gen/webserver"
	"net/http"
	"os"
)

func init() {
	// check for a track file override
	trackFile := os.Getenv(selector.TrackEnvVar)
	if trackFile != "" {
		selector.TracksYamlFilePath = trackFile
	}
	fmt.Printf("tracks to be loaded from \"%s\"\n", selector.TracksYamlFilePath)

	// Load all tracks into memory on initialisation
	err := selector.LoadTracks()
	if err != nil {
		// no benefit to catching an error that failed to be written
		_, _ = fmt.Fprintf(os.Stderr, "failed to load tracks: %v\n", err)
		os.Exit(1)
	}

	// set the template dir
	templateDir := os.Getenv(webserver.TemplateDirEnvVar)
	if templateDir != "" {
		webserver.TemplateDir = templateDir
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := webserver.LoadRenderAndWrite("root", w, nil)
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}
	})

	// TODO: allow the number of tracks to be input as you generally input 1 track at a time
	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		tracks, err := selector.SelectTracksAndRemove(4)
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}

		err = webserver.LoadRenderAndWrite("random", w, struct {
			Tracks          []selector.Track
			RemainingTracks []selector.Track
			SelectedTracks  []selector.Track
		}{
			Tracks:          tracks,
			RemainingTracks: selector.AllTracks,
			SelectedTracks:  selector.SelectedTracks,
		})
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}
	})

	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		err := selector.LoadTracks()
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}

		// if there isn't a referer page return to the homepage on reload
		redirectTarget := r.Referer()
		if redirectTarget == "" {
			redirectTarget = "/"
		}

		err = webserver.LoadRenderAndWrite("reload", w, struct {
			Referer string
		}{
			Referer: redirectTarget,
		})
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// no benefit to catching an error that failed to be written
		_, _ = fmt.Fprintf(os.Stderr, "failed to start webserver: %v\n", err)
		os.Exit(1)
	}
}
