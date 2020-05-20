package octodex

import "flag"

const (
	flagName        = "octodex"
	defaultSavePath = "data"
)

func NewCmdParser() (string, *CmdParser) {
	cp := &CmdParser{}
	return flagName, cp
}

type CmdParser struct {
	path string
}

func (cp *CmdParser) ParseCmd(params []string) error {
	flagSet := flag.NewFlagSet(flagName, flag.ExitOnError)
	flagSet.StringVar(&cp.path, "path", defaultSavePath, "path specifies the storage path of the picture.")

	return flagSet.Parse(params)
}

func (cp *CmdParser) Run() {
	oc := NewCrawler(cp)
	oc.PushURL()
}
