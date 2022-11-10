/*
* Should read files from Garanti SFTP
* Then transfer on netopia-automation/statements/garanti_s/
* https://storage.cloud.google.com/netopia-automation/statements/garanti_s/
* gs://netopia-automation/statements/garanti_s/
* Then Delete files from SFTP
 */
package main

import (
	"finance-utility/configs"
	"finance-utility/pkg/garantiFileTransfer/controllers"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	// Run bash file to Download
	_, err := exec.Command("/bin/sh", "scripts/shell/garantiFileTransfer.sh").Output()
	if err != nil {
		fmt.Printf("Exec shell error! %v", err)
		log.Fatal(err)
		return
	}

	// // Read list of Downloadpool
	files, err := ioutil.ReadDir(configs.DOWNLOAD_PATH)
	if err != nil {
		fmt.Printf("Read downloadpool error! %v", err)
		log.Fatal(err)
		return
	}

	// filePoolList := ""
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
			// filePoolList = filePoolList + file.Name() + "\n"
		}
	}

	controllers.ListFiles()

}
