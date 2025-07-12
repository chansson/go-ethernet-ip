package listInterface

import (
	"github.com/chansson/ethernet-ip/command"
	"github.com/chansson/ethernet-ip/messages/packet"
	"github.com/chansson/ethernet-ip/types"
)

func New(context types.ULInt) (*packet.Packet, error) {
	return &packet.Packet{
		Header: packet.Header{
			Command:       command.ListInterfaces,
			Length:        0,
			SessionHandle: 0,
			Status:        0,
			SenderContext: context,
			Options:       0,
		},
		SpecificData: nil,
	}, nil
}
