package selector

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

const (
	KartEnvVar = "MK_GEN_KART_FILE"
)

type Kart struct {
	Name string `csv:"Name"`
	Type string `csv:"Type"`
}

var (
	// AllKarts is an array containing ALL karts that can be selected, this variable should not be mutated
	AllKarts []*Kart

	// KartsCsvFilePath is the path to the CSV file containing the kart data
	// can be overridden by the `KartEnvVar` environment variable
	KartsCsvFilePath = "./static/karts.csv"
)

func loadKarts() error {
	file, err := os.OpenFile(KartsCsvFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.UnmarshalFile(file, &AllKarts)
	if err != nil {
		return err
	}

	if len(AllKarts) < 1 {
		return fmt.Errorf("zero karts stored in memory")
	}

	fmt.Println("karts successfully loaded")

	return nil
}
