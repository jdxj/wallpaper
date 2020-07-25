package darenyou

const (
	flagName = "darenyou"

	defaultSavePath = "data"

	// project
	Chaos        = "chaos"
	Hysteresis   = "Hysteresis"
	Commissioned = "commissioned"

	// size
	// 注意: Hysteresis 系列的图片中,
	// src_o 与 data-hi-res 大小是相反的.
	Src       = "src"
	SrcO      = "src_o"
	DataHiRes = "data-hi-res" // goquery 无法解析
)

type Flags struct {
	Project string
	Size    string
}
