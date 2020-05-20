package octodex

import "flag"

const (
	flagName        = "octodex"
	defaultSavePaht = "data"
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

	flagSet.StringVar(&cp.path, "path", defaultSavePaht, "path specifies the storage path of the picture.")

	return flagSet.Parse(params)
}

func (cp *CmdParser) Run() {
	oc := NewCrawler(cp)
	oc.PushURL()
}
