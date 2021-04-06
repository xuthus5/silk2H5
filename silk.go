package silk2H5

import (
	"errors"

	"github.com/xuthus5/silk2H5/silk"
)

type Silk2H5 struct {
}

func New() *Silk2H5 {
	return &Silk2H5{}
}

func (s *Silk2H5) TransSilkByteToWavByte(src []byte) ([]byte, error) {
	if src == nil {
		return nil, errors.New("src empty")
	}

	body, err := silk.TransByte(src, silk.TransTypeWav)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *Silk2H5) TransSilkByteToMp3Byte(src []byte) ([]byte, error) {
	if src == nil {
		return nil, errors.New("src empty")
	}

	body, err := silk.TransByte(src, silk.TransTypeMp3)
	if err != nil {
		return nil, err
	}
	return body, nil
}
