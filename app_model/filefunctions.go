package app_model

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func appendToFile(iFileName, iContent string) {

	if _, err := os.Stat(iFileName); err == nil {

		f, err := os.OpenFile(iFileName, os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		if _, err = io.WriteString(f, iContent); err != nil {
			//return err
		}

		f.Sync()
	} else {
		f, err := os.Create(iFileName)
		if err != nil {
			//return err
		}
		defer f.Close()

		if _, err = io.WriteString(f, iContent); err != nil {
			//return err
		}

		f.Sync()
	}

}

func renameFile(iFileName, iFileNameNew string) {
	err := os.Rename(iFileName, iFileNameNew)
	if err != nil {
		log.Fatal(err)
	}
}

func copyFile(iFileFrom, iFileTo string) {
	input, err := ioutil.ReadFile(iFileFrom)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile(iFileTo, input, 0644)
	if err != nil {
		log.Println("Error creating", iFileTo)
		log.Println(err)
		return
	}
}
