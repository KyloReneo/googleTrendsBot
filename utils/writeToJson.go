package utils

import (
	"log"
	"os"

)

func WriteToJSON(json string) {
	// 0777 for read, write, & execute for owner, group and others
	err := os.WriteFile("./result.json", []byte(json), 0777)
	if err != nil {
		log.Fatalf("Somthing went wrong when creating json file, original error:\n %v", err)
	}
}
