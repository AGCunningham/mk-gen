package selector

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	TrackEnvVar = "MK_GEN_TRACK_FILE"
)

type Track struct {
	Name       string `yaml:"name"`
	Cup        string `yaml:"cup"`
	TrackImage string `yaml:"img"`
}

var (
	// AllTracks is an array containing ALL tracks that can be selected, this variable should not be mutated
	AllTracks []Track

	// SelectedTracks contains tracks which have been previously selected
	SelectedTracks []Track

	// TracksYamlFilePath is the path to the YAML file containing the track data
	// can be overridden by the `TrackEnvVar` environment variable
	TracksYamlFilePath = "./static/tracks.yaml"
)

func LoadTracks() error {
	yamlFile, err := os.ReadFile(TracksYamlFilePath)
	if err != nil {
		return err
	}

	var tracks struct {
		Tracks []Track `yaml:"tracks"`
	}

	err = yaml.Unmarshal(yamlFile, &tracks)
	if err != nil {
		return err
	}

	// reset variables
	AllTracks = tracks.Tracks
	SelectedTracks = []Track{}

	if len(AllTracks) < 1 {
		return fmt.Errorf("zero tracks stored in memory")
	}

	fmt.Println("tracks successfully loaded")

	return nil
}
