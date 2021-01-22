package sub

import (
	"github.com/fatedier/golib/crypto"
	"log"
	"testing"
)

func TestRunClient(t *testing.T) {
	crypto.DefaultSalt = "frp"
	err := RunClient("127.0.0.1", 7000, 22, 6000, "minername")
	log.Println(err)
}
