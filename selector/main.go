package selector

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	PlayerCountQueryParam = "player-count"
)

type Player struct {
	Character string
	Kart      string
	Wheels    string
	Glider    string
}

func LoadAll() error {
	sources := map[string]func() error{
		"tracks":     loadTracks,
		"karts":      loadKarts,
		"characters": loadCharacters,
		"tyres":      loadTyres,
		"gliders":    loadGliders,
	}

	for typeName, f := range sources {
		err := f()
		if err != nil {
			return fmt.Errorf("failed to load %s: %v\n", typeName, err)
		}
	}

	return nil
}

func SelectTracksAndRemove(numberOfTracks int) ([]Track, error) {
	// Create and seed the generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// if there are not sufficient unloaded tracks reload all tracks again
	if val := len(AllTracks); val < numberOfTracks {
		fmt.Printf("require \"%d\" unselected tracks, only \"%d\" remain so reloading...\n", numberOfTracks, val)
		err := loadTracks()
		if err != nil {
			return []Track{}, fmt.Errorf("failed to reload tracks: %v", err)
		}
	}

	var selectedTracks []Track

	// select the requested number of tracks
	for range numberOfTracks {
		// select a random track out of remaining tracks
		idx := r.Intn(len(AllTracks))

		// append the selected track to selected tracks
		selectedTracks = append(selectedTracks, AllTracks[idx])

		// remove the selected track - order doesn't matter as it's a new random number each time
		newLength := len(AllTracks) - 1
		AllTracks[idx] = AllTracks[newLength]
		AllTracks = AllTracks[:newLength]
	}

	SelectedTracks = append(SelectedTracks, selectedTracks...)

	return selectedTracks, nil
}

func SelectPlayers(numberOfPlayers int) ([]*Player, error) {
	// Create and seed the generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var players []*Player

	for range numberOfPlayers {
		// FIXME: ought to have some validation to make sure that the variables exist
		players = append(players, &Player{
			Character: AllCharacters[r.Intn(len(AllCharacters))].Name,
			Kart:      AllKarts[r.Intn(len(AllKarts))].Name,
			Wheels:    AllTyres[r.Intn(len(AllTyres))].Name,
			Glider:    AllGliders[r.Intn(len(AllGliders))].Name,
		})
	}

	return players, nil
}
