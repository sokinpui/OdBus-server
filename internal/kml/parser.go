package kml

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"strconv"
)

// LatLong represents a latitude-longitude coordinate pair.
type LatLong struct {
	Latitude  float64
	Longitude float64
}

// kml is the root element of a KML file.
type kml struct {
	Document Document `xml:"Document"`
}

// Document contains a list of Placemarks.
type Document struct {
	Placemarks []Placemark `xml:"Placemark"`
	Folders    []Folder    `xml:"Folder"`
}

// Folder contains a list of Placemarks and other Folders.
type Folder struct {
	Placemarks []Placemark `xml:"Placemark"`
	Folders    []Folder    `xml:"Folder"`
}

// Placemark contains a Point.
type Placemark struct {
	Point Point `xml:"Point"`
}

// Point contains the coordinates.
type Point struct {
	Coordinates string `xml:"coordinates"`
}

// ParseKMZ reads a KMZ file, extracts the KML file, and parses the coordinates.
// It assumes the KML file has Placemarks with Point coordinates in "longitude,latitude,altitude" format.
func ParseKMZ(kmzPath string) ([]LatLong, error) {
	reader, err := zip.OpenReader(kmzPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open KMZ file: %w", err)
	}
	defer reader.Close()

	var kmlFile *zip.File
	for _, file := range reader.File {
		if strings.HasSuffix(strings.ToLower(file.Name), ".kml") {
			kmlFile = file
			break
		}
	}

	if kmlFile == nil {
		return nil, fmt.Errorf("no KML file found in KMZ archive")
	}

	rc, err := kmlFile.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open KML file from archive: %w", err)
	}
	defer rc.Close()

	return parseKML(rc)
}

func parseKML(reader io.Reader) ([]LatLong, error) {
	var kmlData kml
	if err := xml.NewDecoder(reader).Decode(&kmlData); err != nil {
		return nil, fmt.Errorf("failed to decode KML: %w", err)
	}

	var placemarks []Placemark
	placemarks = append(placemarks, kmlData.Document.Placemarks...)

	var collectPlacemarks func(folders []Folder)
	collectPlacemarks = func(folders []Folder) {
		for _, folder := range folders {
			placemarks = append(placemarks, folder.Placemarks...)
			collectPlacemarks(folder.Folders)
		}
	}
	collectPlacemarks(kmlData.Document.Folders)

	var latLongs []LatLong
	for _, placemark := range placemarks {
		coordsStr := strings.TrimSpace(placemark.Point.Coordinates)
		if coordsStr == "" {
			continue
		}

		parts := strings.Split(coordsStr, ",")
		if len(parts) < 2 {
			continue
		}

		longitude, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			continue
		}

		latitude, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			continue
		}

		latLongs = append(latLongs, LatLong{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	return latLongs, nil
}
