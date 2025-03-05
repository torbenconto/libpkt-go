package libpkt_go

import (
	"os"
	"testing"
)

var filePath = "test.pkt"
var testHeader = &Header{
	MagicNumber: MagicNumber,
	Version:     0x01,
	Reserved:    0x00,
	Length:      0x00,
}

func TestWriteHeader(t *testing.T) {
	// if file exists, remove it
	if _, err := os.Stat(filePath); err == nil {
		err = os.Remove(filePath)
		if err != nil {
			t.Error(err)
		}
	}

	// Open a .pkt file
	file, err := os.Create(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	// Write the header
	err = WriteHeader(file, testHeader)
	if err != nil {
		t.Error(err)
	}
}

func TestReadHeader(t *testing.T) {
	// Ensure TestWriteHeader succeeded by checking the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Skip("TestWriteHeader did not create the file, skipping TestReadHeader")
	}
	defer os.Remove(filePath)

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

	// check if header is equal to testHeader
	if header.MagicNumber != testHeader.MagicNumber {
		t.Error("MagicNumber does not match")
	}

	if header.Version != testHeader.Version {
		t.Error("Version does not match")
	}

	if header.Reserved != testHeader.Reserved {
		t.Error("Reserved does not match")
	}

	if header.Length != testHeader.Length {
		t.Error("Length does not match")
	}

	// Remove the test file
	err = os.Remove(filePath)
	if err != nil {
		t.Error(err)
	}
}
