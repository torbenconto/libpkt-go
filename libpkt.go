package libpkt_go

import (
	"encoding/binary"
	"errors"
	"os"
)

const MagicNumber = 0x504B5400

// Header struct represents the header of a .pkt file, magic number is defined as constant libpkt_go.MagicNumber
// Header length is fixed at 11 bytes. (32-bit + 8-bit + 16-bit + 32-bit) = 11 bytes
type Header struct {
	MagicNumber uint32

	// Version is the version of the .pkt file format
	// 0x01 - Original version
	Version  uint8
	Reserved uint16
	// Length is NOT the length of any single packet or the total length of the file in bytes, but the count of packets in the file.
	Length uint32
}

// Packet struct represents a single packet in a .pkt file and is the main data structure of the library.
type Packet struct {
	// Type is the type of packet (eg: ARP, IP, TCP, UDP, etc.)
	// The type is of a custom go type PacketType, which is an unsigned 16-bit integer.
	// The type is defined in the types.go file.
	// This is a custom type for the purpose of ensuring that only valid/known types are used.
	Type uint16
	// Unix timestamp of the packet
	Timestamp uint64
	// Length of the packet in bytes
	Length uint32
	// Data is the actual packet data
	// Data is a byte slice equivalent to a dynamic []uint8_t in C.
	Data []byte
}

// ReadHeader reads the header from the file
func ReadHeader(file *os.File) (*Header, error) {
	if file == nil {
		return nil, errors.New("invalid file")
	}

	header := &Header{}
	_, err := file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	err = binary.Read(file, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	if header.MagicNumber != MagicNumber {
		return nil, errors.New("invalid .pkt file")
	}

	return header, nil
}

func WriteHeader(file *os.File, header *Header) error {
	if file == nil {
		return errors.New("invalid file")
	}

	if header == nil {
		return errors.New("invalid header")
	}

	_, err := file.Seek(0, 0)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	return nil
}
