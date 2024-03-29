package silk

/*
#cgo LDFLAGS: -L . -lSKP_SILK_SDK
#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include "SKP_Silk_SDK_API.h"
#include "Decoder.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/xuthus5/silk2H5/core"
	"github.com/xuthus5/silk2H5/transcoder/ffmpeg"
	"io/ioutil"
	"os"
	"strings"
	"unsafe"
)

func genPathID() string {
	var id = uuid.New().String()
	return fmt.Sprintf("/tmp/%s", strings.Replace(id, "-", "", -1))
}

func TransByte(buf []byte, transType core.TransType) ([]byte, error) {
	if buf == nil {
		return nil, errors.New("buf empty")
	}

	// 写入本地 经由 ffmpeg 转换本地文件
	namePrefix := genPathID()

	var inputPath = namePrefix + ".silk"

	err := ioutil.WriteFile(inputPath, buf, 0666)
	if err != nil {
		return nil, err
	}

	var suffix string
	switch transType {
	case core.TransType_TransTypeWAV:
		suffix = ".wav"
	case core.TransType_TransTypeMP3:
		suffix = ".mp3"
	default:
		return nil, errors.New("unsupported type")
	}

	var outputPath = namePrefix + ".pcm"
	var wavPath = namePrefix + suffix
	inputPathC := C.CString(inputPath)
	outPathC := C.CString(outputPath)
	// 将 silk 格式转码为 pcm 中间文件
	var retCode = C.Decoder(inputPathC, outPathC)
	rc := int(retCode)
	C.free(unsafe.Pointer(inputPathC))
	C.free(unsafe.Pointer(outPathC))
	if rc != 0 {
		return nil, errors.New("decode amr error")
	}

	err = transPcmToAudio(outputPath, wavPath)
	if err != nil {
		return nil, err
	}

	if err = fileRemove(outputPath); err != nil {
		return nil, err
	}

	// 声音提高
	if err := increaseAudioVolume(wavPath); err != nil {
		return nil, err
	}

	// 读取
	f, err := ioutil.ReadFile(wavPath)
	if err != nil {
		return nil, err
	}

	if err = fileRemove(wavPath); err != nil {
		return nil, err
	}
	if err = fileRemove(inputPath); err != nil {
		return nil, err
	}

	return f, nil
}

func increaseAudioVolume(srcPath string) error {
	audioFilter := "'volume=1.5'"
	opts := ffmpeg.Options{
		AudioFilter: &audioFilter,
	}
	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath: "ffmpeg",
	}
	_, err := ffmpeg.
		New(ffmpegConf).
		Input(srcPath).
		Output(srcPath).
		WithOptions(opts).
		Start(opts)
	if err != nil {
		return err
	}
	return nil
}

func transPcmToAudio(inputPath, OutputPath string) error {
	format := "s16le"
	overwrite := true
	audioCodec := "pcm_s16le"
	audioChannels := 2
	audioRate := 12000
	opts := ffmpeg.Options{
		Overwrite:     &overwrite,
		OutputFormat:  &format,
		AudioChannels: &audioChannels,
		AudioRate:     &audioRate,
		AudioCodec:    &audioCodec,
	}
	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath: "ffmpeg",
	}
	_, err := ffmpeg.
		New(ffmpegConf).
		Input(inputPath).
		Output(OutputPath).
		WithOptions(opts).
		Start(opts)
	if err != nil {
		return err
	}
	return nil
}

func fileRemove(logFile string) error {
	_, err := os.Stat(logFile)
	if err == nil {
		return os.Remove(logFile)
	}
	return err
}
