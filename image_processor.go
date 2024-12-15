package services

import (
	"log"
	"strings"
)

func ProcessImages(imageURLs []string) []string {
	var compressedImages []string
	for _, url := range imageURLs {
		compressedURL := strings.Replace(url, "original", "compressed", 1)
		compressedImages = append(compressedImages, compressedURL)
		log.Println("Processed image:", compressedURL)
	}
	return compressedImages
}
