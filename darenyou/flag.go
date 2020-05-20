package darenyou

import (
	"flag"
)

const (
	flagName = "darenyou"

	defaultSavePath = "data"

	// project
	Chaos        = "chaos"
	Commissioned = "commissioned"

	// size
	Src       = "src"
	SrcO      = "src_o"
	DataHiRes = "data-hi-res" // goquery 无法解析
)

func NewCmdParser() (string, *CmdParser) {
	cp := &CmdParser{}
	return flagName, cp
}

type CmdParser struct {
	project string
	size    string
	path    string
}

func (cp *CmdParser) ParseCmd(param []string) error {
	flagSet := flag.NewFlagSet(flagName, flag.ExitOnError)

	flagSet.StringVar(&cp.project, "project", Chaos, "project specifies a different album. you can choose are [chaos | commissioned]")
	flagSet.StringVar(&cp.size, "size", Src, "size specifies the resolution of the image to be downloaded. you can choose [src | src_o | data-hi-res]")
	flagSet.StringVar(&cp.path, "path", defaultSavePath, "path specifies the storage path of the picture.")

	return flagSet.Parse(param)
}

func (cp *CmdParser) Run() {
	c := NewCrawler(cp)
	c.PushURL()
}
