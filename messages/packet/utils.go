package packet

import (
	"github.com/chansson/ethernet-ip/bufferx"
)

func Paths(arg ...[]byte) []byte {
	io := bufferx.New(nil)
	for i := 0; i < len(arg); i++ {
		io.WL(arg[i])
	}
	return io.Bytes()
}
