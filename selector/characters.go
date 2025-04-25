package selector

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

const (
	CharacterEnvVar = "MK_GEN_CHARACTER_FILE"
)

type Character struct {
	Name string `csv:"Name"`
}

var (
	// AllCharacters is an array containing ALL characters that can be selected, this variable should not be mutated
	AllCharacters []*Character

	// CharactersCsvFilePath is the path to the CSV file containing the character data
	// can be overridden by the `CharacterEnvVar` environment variable
	CharactersCsvFilePath = "./static/characters.csv"
)

func loadCharacters() error {
	file, err := os.OpenFile(CharactersCsvFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.UnmarshalFile(file, &AllCharacters)
	if err != nil {
		return err
	}

	if len(AllCharacters) < 1 {
		return fmt.Errorf("zero characters stored in memory")
	}

	fmt.Println("characters successfully loaded")

	return nil
}
