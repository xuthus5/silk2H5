package silk2H5

var (
	std = New()
)

func TransToWavByte(src []byte) ([]byte, error) {
	return std.TransSilkByteToWavByte(src)
}

func TransToMp3Byte(src []byte) ([]byte, error) {
	return std.TransSilkByteToMp3Byte(src)
}
