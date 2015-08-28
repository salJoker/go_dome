package main

import(
	"os"
	"fmt"
	"path/filepath"
	"io/ioutil"
)

func main(){
//	for i,val := range os.Args {
//		fmt.Printf("%s : %d",val,i)
//	}
	driver := os.Args[1]
	drv,err := filepath.Abs(driver)
	if err != nil {
		return
	}
	fileindos,err := ioutil.ReadDir(drv)
	if err != nil {
		return
	}
	fmt.Printf("fileinfos length : %d\n",len(fileindos))
	var fileCount int = 0
	for _,fileinfo := range fileindos {
		if !fileinfo.IsDir(){
			fmt.Printf("filename:%s\n",fileinfo.Name())
			fileCount++
		}
	}
	fmt.Println("file count : ",fileCount)
}
