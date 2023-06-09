package common

import (
	"bufio"
	"mime/multipart"
	"os"
)

func SaveMultipartFile(file *multipart.FileHeader, filePath string) (err error) {
	httpFile, err := file.Open()
	reader := bufio.NewReader(httpFile)
	localFile, err := os.Create(filePath)
	defer func(localFile *os.File) {
		err := localFile.Close()
		if err != nil {
			panic(err)
		}
	}(localFile)
	if err != nil {
		return
	}
	_, err = reader.WriteTo(localFile)
	return
}
