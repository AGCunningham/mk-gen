package selector

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

const (
	GliderEnvVar = "MK_GEN_GLIDER_FILE"
)

type Glider struct {
	Name string `csv:"Name"`
}

var (
	// AllGliders is an array containing ALL gliders that can be selected, this variable should not be mutated
	AllGliders []*Glider

	// GlidersCsvFilePath is the path to the CSV file containing the glider data
	// can be overridden by the `GliderEnvVar` environment variable
	GlidersCsvFilePath = "./static/gliders.csv"
)

func loadGliders() error {
	file, err := os.OpenFile(GlidersCsvFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.UnmarshalFile(file, &AllGliders)
	if err != nil {
		return err
	}

	if len(AllGliders) < 1 {
		return fmt.Errorf("zero gliders stored in memory")
	}

	fmt.Println("gliders successfully loaded")

	return nil
}
