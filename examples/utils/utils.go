package utils

import (
	"encoding/json"
	"fmt"
	lotrsdk "lotrsdk/sdk"
	"math/rand"
)

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandIndex(indexMap map[int]bool, indexLen int) int {
	randValue := RandInt(0, indexLen-1)
	_, in := indexMap[randValue]
	if !in {
		indexMap[randValue] = true
		return randValue
	}
	RandIndex(indexMap, indexLen)
	return -1
}

func GetRandomDocs(docs lotrsdk.Documenter, maxCount int) ([]interface{}, error) {
	indexes := make(map[int]bool)
	var selectedDocs []interface{}
	docArray := docs.GetDocs()
	docsLen := len(docArray)
	if docsLen >= maxCount {
		for i := 0; i <= maxCount-1; i++ {
			randValue := RandIndex(indexes, docsLen)
			if randValue != -1 {
				selectedDocs = append(selectedDocs, docArray[randValue])
			} else {
				return nil, fmt.Errorf("a random index could not be created")
			}
		}
	} else {
		selectedDocs = docArray
	}
	return docArray, nil
}

func ToJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
