package selector

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

const (
	TyreEnvVar = "MK_GEN_TYRE_FILE"
)

type Tyre struct {
	Name string `csv:"Name"`
}

var (
	// AllTyres is an array containing ALL tyres that can be selected, this variable should not be mutated
	AllTyres []*Tyre

	// TyresCsvFilePath is the path to the CSV file containing the tyre data
	// can be overridden by the `TyreEnvVar` environment variable
	TyresCsvFilePath = "./static/tyres.csv"
)

func loadTyres() error {
	file, err := os.OpenFile(TyresCsvFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.UnmarshalFile(file, &AllTyres)
	if err != nil {
		return err
	}

	if len(AllTyres) < 1 {
		return fmt.Errorf("zero tyres stored in memory")
	}

	fmt.Println("tyres successfully loaded")

	return nil
}
