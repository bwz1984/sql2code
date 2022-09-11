package code2file

import (
	"log"
	"os"
)

func CodeFileWrite(code, fileName string) (string, error) {
	folderPath := "./output/"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0777)
	}
	filePath := folderPath + fileName
	fh, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Printf("openFile fail,err:%v filePath:%v", err, filePath)
		return "", err
	}
	//及时关闭file句柄
	defer fh.Close()

	n, err := fh.Seek(0, os.SEEK_END)
	if err != nil {
		log.Printf("Seek fail,err:%v filePath:%v", err, filePath)
		return "", err
	}
	_, err = fh.WriteAt([]byte(code), n)
	if err != nil {
		log.Printf("WriteAt fail,err:%v filePath:%v", err, filePath)
		return "", err
	}
	return filePath, nil
}
