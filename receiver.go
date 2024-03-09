package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/riccardotornesello/f1-23-parser-go/packets"
)

type Receiver struct {
	server *net.UDPConn

	eventHandlers map[packets.PacketId]interface{}
}

func NewReceiver(host string, port int) *Receiver {
	// Start UDP server
	server, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(host),
	})
	if err != nil {
		log.Fatal("Error starting UDP server: ", err)
	}

	return &Receiver{
		server:        server,
		eventHandlers: make(map[packets.PacketId]interface{}),
	}
}

func (r *Receiver) Start() {
	// Read UDP packets
	for {
		buffer := make([]byte, 4096)
		_, _, err := r.server.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}

		header := new(packets.PacketHeader)
		reader := bytes.NewReader(buffer)
		if err = binary.Read(reader, binary.LittleEndian, header); err != nil {
			panic(err)
		}

		if header.PacketFormat != 2023 {
			panic(fmt.Sprintf("Unknown packet format: %d", header.PacketFormat))
		}

		pack := NewPacketById(header.PacketId)
		if pack == nil {
			panic(fmt.Sprintf("Unknown packet id: %d", header.PacketId))
		}

		reader = bytes.NewReader(buffer)
		if err = binary.Read(reader, binary.LittleEndian, pack); err != nil {
			panic(err)
		}

		if handler, ok := r.eventHandlers[header.PacketId]; ok {
			handler.(func(interface{}))(pack)
		}
	}
}

func (r *Receiver) On(eventId packets.PacketId, fn interface{}) {
	r.eventHandlers[eventId] = fn
}

func (r *Receiver) Close() {
	r.server.Close()
}
