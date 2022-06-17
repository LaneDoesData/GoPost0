package main

import (
	"CODE/Utils"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type IrisStruct struct {
	ID            int
	SepalLengthCm float64
	SepalWidthCm  float64
	PetalLengthCm float64
	PetalWidthCm  float64
	Species       string
}

func main() {
	
	fmt.Println("Showing dataset in slice of slice")
	sliceData := readDataSimple(`Data\Iris.csv`)
	for idx, v := range sliceData {
		if idx == 5 {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("Showing dataset in struct")
	structData := readDataStruct(`Data\Iris_noHeader.csv`)

	for idx, v := range structData {
		if idx == 5 {
			break
		}
		fmt.Println(v)
	}
}

func HandleErr(err error) bool {
	//this is a simple error handler
	//the only job is to reduce eye junk in the code
	if err != nil {
		return true
	}
	return false
}

func readDataSimple(filePath string) [][]string {

	myCsvFile, err := os.Open(filePath)
	if HandleErr(err) {
		fmt.Println(err)
	}

	defer myCsvFile.Close()

	CSVReader := csv.NewReader(myCsvFile)
	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count. If FieldsPerRecord is negative, no check is
	// made and records may have a variable number of fields.
	CSVReader.FieldsPerRecord = -1

	DataCsv, err := CSVReader.ReadAll()

	if HandleErr(err) {
		fmt.Println(err)
	}

	return DataCsv

}

func readDataStruct(filePath string) []IrisStruct {

	//This function reads in a csv dataset with a known shape and columns.

	MasterStruct := []IrisStruct{}
	myCsvFile, err := os.Open(filePath)
	if HandleErr(err) {
		fmt.Println(err)
	}

	defer myCsvFile.Close()

	CSVReader := csv.NewReader(myCsvFile)
	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count. If FieldsPerRecord is negative, no check is
	// made and records may have a variable number of fields.
	CSVReader.FieldsPerRecord = -1

	if HandleErr(err) {
		fmt.Println(err)
	}
	//A loop to go over all records and append to the master struct
	for {
		record, err := CSVReader.Read()
		if err == io.EOF {
			break
		}

		tempStruct := IrisStruct{
			Utils.String2Int(record[0]),
			Utils.String2float(record[1]),
			Utils.String2float(record[2]),
			Utils.String2float(record[3]),
			Utils.String2float(record[4]),
			record[5],
		}
		MasterStruct = append(MasterStruct, tempStruct)

	}

	return MasterStruct

}
