package sub

import (
	"github.com/fatedier/golib/crypto"
	"github.com/panglove/frp/pkg/util"
	"log"
	"testing"
)

func TestRunClient(t *testing.T) {
	crypto.DefaultSalt = "frp"
	ips := util.ExternalIP()
	ipsStr := ""
	if ips!=nil && len(ips)>0 {
		for _,v :=range  ips {
			ipsStr =ipsStr+ v+"|"
		}
		ipsStr = ipsStr[:len(ipsStr)-1]
	}
	err := RunClient("127.0.0.1", 7000, 22, 6000, "minername",ipsStr)
	log.Println(err)
}
