package fun

import (
	"errors"
	"bytes"
	"encoding/binary"
	"net"
	"io"
)

var (
	ErrMessageInvalid = errors.New("message invalid")
)


type Message interface {
	Serialize() ([]byte, error)
	MessageType() int32
}

//TestMessage implement Message interface
type TestMessage string

func (t TestMessage) Serialize() ([]byte, error) {
	if len(string(t)) == 0 {
		return nil, ErrMessageInvalid
	}
	return []byte(string(t)), nil
}

func (t TestMessage) MessageType() int32 {
	return 0
}

type Codecs interface {
	Encode(message Message) ([]byte, error)
	Decode(conn net.Conn) (Message, error)
}

//DefaultCodecs implement Codecs interface
type DefaultCodecs struct {}

func (df *DefaultCodecs) Encode(message Message) ([]byte, error) {
	data, err := message.Serialize()
	if err != nil {
		return nil,err
	}
	msgType := message.MessageType()
	msgLen := len(data)
	w := new(bytes.Buffer)
	err = binary.Write(w,binary.LittleEndian,&msgType)
	if err != nil {
		return nil, err
	}
	err = binary.Write(w,binary.LittleEndian,&msgLen)
	if err != nil {
		return nil, err
	}
	w.Write(data)
	return w.Bytes(),nil
}

func (df *DefaultCodecs) Decode(conn net.Conn) (Message, error) {
	msgType := make([]byte, 4)
	msgLen := make([]byte, 4)
	_, err := io.ReadFull(conn, msgType)
	if err != nil {
		return nil, err
	}
	_, err = io.ReadFull(conn, msgLen)
	if err != nil {
		return nil, err
	}
	var length int32
	r := bytes.NewReader(msgLen)
	err = binary.Read(r, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	data := make([]byte, length)
	_, err = io.ReadFull(conn,data)
	if err != nil {
		return nil, err
	}

}