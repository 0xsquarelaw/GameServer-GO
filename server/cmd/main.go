package main

import (
	"fmt"
	"server/pkg/packets"
)

func main() {
	packet := &packets.Packet{
		SenderId: 79,
		Msg: &packets.Packet_Chat{
			Chat: &packets.ChatMessage{
				Msg: "Hello, world!",
			},
		},
	}

	fmt.Println(packet)

}
