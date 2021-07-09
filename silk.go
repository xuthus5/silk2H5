package silk2H5

import (
	"errors"

	"github.com/xuthus5/silk2H5/silk"
)

type Silk2H5 struct {
	Rate int64 // 声音频率 默认 24000
}

// New 实例化一个转换对象
func New() *Silk2H5 {
	return &Silk2H5{}
}

// ToWavByte 转换silk为wav
func (s *Silk2H5) ToWavByte(src []byte) ([]byte, error) {
	if src == nil {
		return nil, errors.New("src empty")
	}

	body, err := silk.TransByte(src, TransType_TransTypeWAV)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// ToMp3Byte 转换silk为mp3
func (s *Silk2H5) ToMp3Byte(src []byte) ([]byte, error) {
	if src == nil {
		return nil, errors.New("src empty")
	}

	body, err := silk.TransByte(src, TransType_TransTypeMP3)
	if err != nil {
		return nil, err
	}
	return body, nil
}
