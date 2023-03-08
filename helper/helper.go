package helper

import "os"

func GetPathImg(fileName string) string {
	host := os.Getenv("HOST")
	return host + "/file/" + fileName
}