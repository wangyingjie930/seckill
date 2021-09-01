package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var filePath = "/Users/wangyingjie/Desktop/sys_region_streets.txt"

/**
文件全部读入内存
 */
func readAll(filePath string) {
	start1 := time.Now()
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	//readAll spend :  45.840708ms
	fmt.Println("readAll spend : ", time.Now().Sub(start1))
}

/**
按行读取, 有带缓冲区的
 */
func readLine(filePath string) {
	start1 := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 200000)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
	}
	fmt.Println("readEachLineReader spend : ", time.Now().Sub(start1))
}

/**
大文件切分成多个小文件
 */
func sliceBigFile(filePath, writeDir string, count int)  {
	start1 := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 200000)
	for j := 0; ; j ++{
		outPutFile := fmt.Sprintf("%s/%d.txt", writeDir, j)

		writeFile, err := os.OpenFile(outPutFile, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 777)
		if err != nil{
			fmt.Println("Open file err =", err)
			writeFile.Close()
			return
		}
		writer := bufio.NewWriter(writeFile)

		//buf := make([]string, 0, count)
		for i := 0; i < count; i ++ {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				j = -1
				break
			}
			//buf = append(buf, string(line))
			_, err = writer.WriteString(string(line) + "\r\n")
			if err != nil{
				fmt.Println("Write file err =", err)
				return
			}
		}
		if j == -1 {
			break
		}
		writer.Flush()
		writeFile.Close()
		fmt.Printf("导出文件: %s\n", outPutFile)
	}
	fmt.Println("readEachLineReader spend : ", time.Now().Sub(start1))
}

func main() {
	//readAll(filePath)
	//readLine(filePath)
	sliceBigFile(filePath, "/Users/wangyingjie/Documents/code/go/src/imooc-product/storages", 1000000)
}
