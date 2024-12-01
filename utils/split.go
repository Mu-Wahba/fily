package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
)

type MetaData struct {
	FileName        string `json:"file_name"`
	DistPath        string `json:"dist_path"`
	NumberOfChuncks int64  `json:"number_of_chuncks"`
}

func Split(filePath, distPath string, chunckSize int64) error {
	chunckSize = chunckSize * 1024 * 1024
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to open file  ", err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatal("Unable to get file size  ", err)
	}
	fileSize := fileStat.Size()
	numberOfChuncks := math.Ceil(float64(fileSize) / float64(chunckSize))

	metadata := &MetaData{
		FileName:        file.Name(),
		DistPath:        distPath,
		NumberOfChuncks: int64(numberOfChuncks),
	}
	//Create File For Metadata at dist
	jsonMetadata, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		log.Fatal("Couldn't marshal data ", err)
	}

	metaDataFileName := filepath.Join(distPath, fmt.Sprintf("%s-metadata.json", filepath.Base(filePath)))
	// Create the full path for the metadata file
	metaDataFile, err := os.Create(metaDataFileName)
	if err != nil {
		log.Fatal("couldn't create metadata file", err)
	}
	defer metaDataFile.Close()

	_, err = metaDataFile.Write(jsonMetadata)
	if err != nil {
		log.Fatal("Couldn't write to metadata file ", err)
	}

	for i := int64(1); i <= int64(numberOfChuncks); i++ {
		//Create file in distPath
		chunckName := fmt.Sprintf("%s-chunck-%d", filepath.Base(filePath), i)
		savedChunckPath := filepath.Join(distPath, chunckName)
		savedChunck, err := os.Create(savedChunckPath)
		if err != nil {
			log.Fatal("couldn't create chunk ", err)
		}
		_, err = io.CopyN(savedChunck, file, chunckSize)
		if err != nil && err != io.EOF {
			log.Fatal("couldn't copy chunck ", err)
		}
		if err = savedChunck.Close(); err != nil {
			log.Fatal("couldn't close chunck ", err)
		}
	}

	return nil
}
