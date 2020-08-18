package tuchong

import (
	"fmt"
	"net"
	"testing"

	"github.com/jdxj/wallpaper/client"
)

func TestLookupHost(t *testing.T) {
	ips, err := net.LookupHost(Host)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for _, ip := range ips {
		fmt.Printf("%v\n", ip)
	}
}

func TestFollow(t *testing.T) {
	flags := &Flags{
		SiteID: "1064195",
	}
	tcd := NewTuChongDLI(flags)
	cli := client.New(30)
	tcd.SetClient(cli)

	if err := tcd.follow(); err != nil {
		t.Fatalf("%s\n", err)
	}
}
