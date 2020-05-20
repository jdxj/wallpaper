package polayoutu

import "flag"

const (
	flagName = "polayoutu"

	defaultSavePath = "data"

	FullRes = "full"
	Thumb   = "thumb"
)

func NewCmdParser() (string, *CmdParser) {
	cp := &CmdParser{}
	return flagName, cp
}

type CmdParser struct {
	size    string
	edition int
	path    string
}

func (cp *CmdParser) ParseCmd(params []string) error {
	flagSet := flag.NewFlagSet(flagName, flag.ExitOnError)

	flagSet.StringVar(&cp.size, "size", Thumb, "size specifies the resolution of the image to be downloaded. you can choose [full | thumb].")
	flagSet.IntVar(&cp.edition, "edition", 1, "edition specifies which period to download.")
	flagSet.StringVar(&cp.path, "path", defaultSavePath, "path specifies the storage path of the picture.")

	return flagSet.Parse(params)
}

func (cp *CmdParser) Run() {
	pc := NewCrawler(cp)
	pc.PushURL()
}
