package silk2H5

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestTransToWavByte(t *testing.T) {
	//fr, err := os.Open("./testdata/test.silk")
	fr, err := os.Open("./testdata/test.go")
	if err != nil {
		fmt.Printf("Open test.silk err: %v", err)
		panic(err)
	}
	defer fr.Close()
	body, err := ioutil.ReadAll(fr)
	if err != nil {
		fmt.Printf("Read test.silk err: %v", err)
		panic(err)
	}
	transBody, err := TransToWavByte(body)
	if err != nil {
		fmt.Printf("TransToWavByte err: %v", err)
		panic(err)
	}

	fw, err := os.Create("./testdata/test_wav.wav")
	if err != nil {
		fmt.Printf("Create test_wav.wav err: %v", err)
		panic(err)
	}
	_, err = fw.Write(transBody)
	if err != nil {
		fmt.Printf("Write test_wav.wav err: %v", err)
		panic(err)
	}
	fw.Close()
}

func TestTransToMp3Byte(t *testing.T) {
	fr, err := os.Open("./testdata/test.silk")
	if err != nil {
		fmt.Printf("Open test.silk err: %v", err)
		panic(err)
	}
	defer fr.Close()
	body, err := ioutil.ReadAll(fr)
	if err != nil {
		fmt.Printf("Read test.silk err: %v", err)
		panic(err)
	}
	transBody, err := TransToMp3Byte(body)
	if err != nil {
		fmt.Printf("TransToWavByte err: %v", err)
		panic(err)
	}

	fw, err := os.Create("./testdata/test_mp3.mp3")
	if err != nil {
		fmt.Printf("Create test_mp3.mp3 err: %v", err)
		panic(err)
	}
	_, err = fw.Write(transBody)
	if err != nil {
		fmt.Printf("Write test_mp3.mp3 err: %v", err)
		panic(err)
	}
	fw.Close()
}
