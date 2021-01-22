package sub

import (
	"github.com/fatedier/golib/crypto"
	"github.com/panglove/frp/pkg/util"
	"log"
	"testing"
	"time"
)

func TestRunClient(t *testing.T) {
	go ConnectFrpServer()
	select {

	}
}

func ConnectFrpServer()  {
	crypto.DefaultSalt = "frp"
	ips := util.ExternalIP()
	ipsStr := ""
	if ips!=nil && len(ips)>0 {
		for _,v :=range  ips {
			ipsStr =ipsStr+ v+"|"
		}
		ipsStr = ipsStr[:len(ipsStr)-1]
	}
	for {
		err := RunClient("127.0.0.1", 7000, 22, 6000, "minername",ipsStr)

		if err!=nil {
			log.Println("connect frp server error:",err)
			log.Println("wait 5 sencond...")
			time.Sleep(time.Second*5)
		}
	}
}
