package silk2H5

var (
	std = New()
)

// ToWav 包函数 直接转换src为wav
func ToWav(src []byte) ([]byte, error) {
	return std.ToWavByte(src)
}

// ToMp3 包函数 直接转换src为mp3
func ToMp3(src []byte) ([]byte, error) {
	return std.ToMp3Byte(src)
}
