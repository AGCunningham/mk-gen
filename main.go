package main

import (
	"fmt"
	"github.com/AGCunningham/mk-gen/selector"
	"github.com/AGCunningham/mk-gen/webserver"
	"net/http"
	"os"
)

func init() {
	// Load all tracks into memory on initialisation
	err := selector.LoadTracks()
	if err != nil {
		// no benefit to catching an error that failed to be written
		_, _ = fmt.Fprintf(os.Stderr, "failed to load tracks: %v", err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := webserver.LoadRenderAndWrite("root", "./templates/root.html", w, nil)
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}
	})

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		tracks, err := selector.SelectTracksAndRemove(4)
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}

		err = webserver.LoadRenderAndWrite("random", "./templates/random.html", w, struct {
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

		err = webserver.LoadRenderAndWrite("reload", "./templates/reload.html", w, struct {
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
