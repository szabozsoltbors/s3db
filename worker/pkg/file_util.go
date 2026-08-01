package pkg

import (
	"log"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
