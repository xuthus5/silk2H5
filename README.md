# silk2H5

> 将silk格式的音频格式(silk, amr)数据转为H5能够播放的格式(wav, mp3)

> 需要 header 为 !SILK_V3

- [x] 编译

需要 `CGO` 支持. 下面是基于 `musl-gcc` 进行的编译示例:

```shell
CGO_ENABLED=1 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" GOOS=linux GOARCH=amd64 go build .
```
