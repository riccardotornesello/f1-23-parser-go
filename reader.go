package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/riccardotornesello/f1-23-parser-go/packets"
)

type FileReader struct {
	fd *os.File

	eventHandlers map[packets.PacketId]interface{}
}

func NewFileReader(filename string) *FileReader {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return &FileReader{
		fd:            fd,
		eventHandlers: make(map[packets.PacketId]interface{}),
	}
}

func (r *FileReader) Read() {
	for {
		currentFilePosition, err := r.fd.Seek(0, io.SeekCurrent)
		if err != nil {
			panic(err)
		}

		header := new(packets.PacketHeader)
		if err := binary.Read(r.fd, binary.LittleEndian, header); err != nil {
			if err.Error() == "EOF" {
				return
			}
			panic(err)
		}

		if header.PacketFormat != 2023 {
			panic(fmt.Sprintf("Unknown packet format: %d", header.PacketFormat))
		}

		pack := NewPacketById(header.PacketId)
		if pack == nil {
			panic(fmt.Sprintf("Unknown packet id: %d", header.PacketId))
		}

		// Seek back to the start of the packet
		if _, err := r.fd.Seek(currentFilePosition, io.SeekStart); err != nil {
			panic(err)
		}
		if err := binary.Read(r.fd, binary.LittleEndian, pack); err != nil {
			panic(err)
		}

		if handler, ok := r.eventHandlers[header.PacketId]; ok {
			reflect.ValueOf(handler).Call([]reflect.Value{
				reflect.ValueOf(pack),
			})
		}
	}
}

func (r *FileReader) On(event packets.PacketId, handler interface{}) {
	r.eventHandlers[event] = handler
}

func (r *FileReader) Close() {
	r.fd.Close()
}
