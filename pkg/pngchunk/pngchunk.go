package pngchunk

import (
	"encoding/binary"
	"io"
)

type (
	Header struct {
		data [8]byte
	}

	Chunk struct {
		length int32
		ctype  [4]byte
		data   []byte
		crc    [4]byte
	}
)

// header functions
func (h *Header) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, &h.data)
}

func (h *Header) Write(w io.Writer) error {
	return binary.Write(w, binary.BigEndian, &h.data)
}

// chunk functions
func (c *Chunk) Read(r io.Reader) error {
	// chunk length
	err := binary.Read(r, binary.BigEndian, &c.length)
	if err != nil {
		return err
	}

	// chunk type
	err = binary.Read(r, binary.BigEndian, &c.ctype)
	if err != nil {
		return err
	}

	// chunk data if length > 0
	if c.length > 0 {
		c.data = make([]byte, c.length)
		err = binary.Read(r, binary.BigEndian, &c.data)
		if err != nil {
			return err
		}
	}

	// chunk CRC
	err = binary.Read(r, binary.BigEndian, &c.crc)
	if err != nil {
		return err
	}

	return nil
}

func (c *Chunk) Write(w io.Writer) error {
	// chunk length
	err := binary.Write(w, binary.BigEndian, c.length)
	if err != nil {
		return err
	}

	// chunk type
	err = binary.Write(w, binary.BigEndian, c.ctype)
	if err != nil {
		return err
	}

	// chunk data
	if c.length > 0 {
		err = binary.Write(w, binary.BigEndian, c.data)
		if err != nil {
			return err
		}
	}

	// chunk crc
	err = binary.Write(w, binary.BigEndian, c.crc)
	if err != nil {
		return err
	}

	return nil
}
