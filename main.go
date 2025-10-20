package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/AGCunningham/mk-gen/selector"
	"github.com/AGCunningham/mk-gen/webserver"
)

const (
	serverPortEnvVar  = "MK_GEN_PORT"
	defaultServerPort = "8080"
)

var (
	serverPort string
)

func init() {
	// check for a track file override
	trackFile := os.Getenv(selector.TrackEnvVar)
	if trackFile != "" {
		selector.TracksYamlFilePath = trackFile
	}
	fmt.Printf("tracks to be loaded from \"%s\"\n", selector.TracksYamlFilePath)
	// check for a characters file override
	characterFile := os.Getenv(selector.CharacterEnvVar)
	if characterFile != "" {
		selector.CharactersCsvFilePath = characterFile
	}
	fmt.Printf("characters to be loaded from \"%s\"\n", selector.CharactersCsvFilePath)
	// check for a gliders file override
	gliderFile := os.Getenv(selector.GliderEnvVar)
	if gliderFile != "" {
		selector.GlidersCsvFilePath = gliderFile
	}
	fmt.Printf("gliders to be loaded from \"%s\"\n", selector.GlidersCsvFilePath)
	// check for a karts file override
	kartFile := os.Getenv(selector.KartEnvVar)
	if kartFile != "" {
		selector.KartsCsvFilePath = kartFile
	}
	fmt.Printf("karts to be loaded from \"%s\"\n", selector.KartsCsvFilePath)
	// check for a tyre file override
	tyreFile := os.Getenv(selector.TyreEnvVar)
	if tyreFile != "" {
		selector.TyresCsvFilePath = tyreFile
	}
	fmt.Printf("tyres to be loaded from \"%s\"\n", selector.TyresCsvFilePath)

	// Load all karts & tracks into memory on initialisation
	err := selector.LoadAll()
	if err != nil {
		// no benefit to catching an error that failed to be written
		_, _ = fmt.Fprintf(os.Stderr, "failed to load data: %v\n", err)
		os.Exit(1)
	}

	// set the template dir
	templateDir := os.Getenv(webserver.TemplateDirEnvVar)
	if templateDir != "" {
		webserver.TemplateDir = templateDir
	}

	// set the port
	serverPort = os.Getenv(serverPortEnvVar)
	if serverPort == "" {
		serverPort = defaultServerPort
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := webserver.LoadRenderAndWrite("root", w, nil)
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}
	})

	http.HandleFunc("/random-tracks", func(w http.ResponseWriter, r *http.Request) {
		tracks, err := selector.SelectTracksAndRemove(4)
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}

		err = webserver.LoadRenderAndWrite("random-tracks", w, struct {
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

	http.HandleFunc("/random-player", func(w http.ResponseWriter, r *http.Request) {
		// fetch the player count from the
		params := r.URL.Query()
		count, err := strconv.Atoi(params.Get(selector.PlayerCountQueryParam))
		if err != nil {
			// default count to 4 if there are any errors
			count = 4
		}

		players, err := selector.SelectPlayers(count)
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}

		err = webserver.LoadRenderAndWrite("random-player", w, struct {
			Players []*selector.Player
		}{
			Players: players,
		})
		if err != nil {
			webserver.PrintAndReturnError(err, w)
		}
	})

	http.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		err := selector.LoadAll()
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

	port := fmt.Sprintf(":%s", serverPort)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		// no benefit to catching an error that failed to be written
		_, _ = fmt.Fprintf(os.Stderr, "failed to start webserver: %v\n", err)
		os.Exit(1)
	}
}
