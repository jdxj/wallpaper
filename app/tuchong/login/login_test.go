package login

import (
	"fmt"
	"net/url"
	"testing"
)

func TestAccount_Login(t *testing.T) {
	flags := &Flags{
		Phone: "",
		Pass:  "",
	}
	acc := NewLogin(flags)
	resp, err := acc.login()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("id: %s, token: %s", resp.Identity, resp.Token)
}

func TestBody(t *testing.T) {
	vs := url.Values{
		"account":  []string{""},
		"password": []string{""},
	}
	fmt.Printf("%s\n", vs.Encode())
}

func TestAccount_RecordNewToken(t *testing.T) {
	flags := &Flags{
		Phone: "",
		Pass:  "",
	}
	acc := NewLogin(flags)
	acc.RecordNewToken()
}
