package fileTools

import (
	"Strava_Light/cmd/gpx/gpx_info"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseXml(filePath string) gpx_info.GpxFile{
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var file gpx_info.GpxFile

	xml.Unmarshal(byteValue, &file)

	xmlFile.Close()

	return file
}
