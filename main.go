package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	AMOUNT_OF_LINES = 50000
	AMOUNT_OF_FILES = 100
	BY_CHANNEL      = iota
	WITHOUT_CHANNEL
)

var wg = sync.WaitGroup{}

func main() {

	//uncomment this line if you wanna test without channel
	// writeByChannel()

	//uncomment this line if you wanna test by channel
	// writeWithoutChannel()

}

//write to the files
func writeToFile(file *os.File, i int, opType int) {
	start := time.Now()

	for i := 0; i < AMOUNT_OF_LINES; i++ {
		file.WriteString("dummy line " + strconv.Itoa(i) + " \n")
	}

	duration := time.Since(start)

	if opType == 0 {
		wg.Done()
	}

	fmt.Println("job number "+strconv.Itoa(i)+" end ", duration)
	time.Sleep(time.Millisecond * 1000)
	file.Close()
}

//write file without using the channels
func writeWithoutChannel() {
	start := time.Now()

	os.Mkdir("files", 0755)

	for i := 0; i < AMOUNT_OF_FILES; i++ {
		file, err := os.Create("files/test" + strconv.Itoa(i) + ".txt")
		checkErr(err)

		writeToFile(file, i, 1)

	}

	duration := time.Since(start)

	fmt.Println("All tasks done without using channel in", duration)
}

//write file by using the channels
func writeByChannel() {
	start := time.Now()

	os.Mkdir("files", 0755)

	for i := 0; i < AMOUNT_OF_FILES; i++ {
		file, err := os.Create("files/test" + strconv.Itoa(i) + ".txt")
		checkErr(err)

		go writeToFile(file, i, 0)

		wg.Add(1)

	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Println("All tasks done by using channel in", duration)
}

//error chacker
func checkErr(err error) {
	if err != nil {
		errors.New("we have some problems")
	}
}
