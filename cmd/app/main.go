package main

import (
	"fmt"
	"github.com/gardasar-code/go-storage/v2/internal/file"
	"github.com/gardasar-code/go-storage/v2/internal/storage"
	"log"
)

func main() {
	fmt.Println("hello")

	st := storage.NewStorage()
	fmt.Printf("%+m\n", st)

	file1, err := file.NewFile("file.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%+m\n", file)
	fmt.Println(file1)

	newfile, err1 := st.UploadFile(file1)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(newfile)

	restoreFile, err2 := st.GetFile(file1.Id)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(restoreFile)

}
