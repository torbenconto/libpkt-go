package libpkt_go

import (
	"os"
	"testing"
)

var filePath = "test.pkt"

func TestWriteHeader(t *testing.T) {

	// Open a .pkt file
	file, err := os.Create(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	// Create a header
	header := &Header{
		MagicNumber: MagicNumber,
		Version:     0x01,
		Reserved:    0x00,
		Length:      0x00,
	}

	// Write the header
	err = WriteHeader(file, header)
	if err != nil {
		t.Error(err)
	}
}

func TestReadHeader(t *testing.T) {
	// Ensure TestWriteHeader succeeded by checking the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skip("TestWriteHeader did not create the file, skipping TestReadHeader")
	}

	// Open the .pkt file for reading
	file, err := os.Open(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	// Read the header
	header, err := ReadHeader(file)
	if err != nil {
		t.Error(err)
	}

	// Check if the header is correct
	if header.MagicNumber != MagicNumber {
		t.Error("Invalid magic number")
	}
}
