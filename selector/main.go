package selector

import (
	"fmt"
	"math/rand"
	"time"
)

func SelectTracksAndRemove(numberOfTracks int) ([]Track, error) {
	// Create and seed the generator.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// if there are not sufficient unloaded tracks reload all tracks again
	if val := len(AllTracks); val < numberOfTracks {
		fmt.Printf("require \"%d\" unselected tracks, only \"%d\" remain so reloading...\n", numberOfTracks, val)
		err := LoadTracks()
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
