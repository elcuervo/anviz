package anviz

import (
	"log"
	"net"
)

type Door struct {
	Url string
	Id  uint16

	conn net.Conn
}

func NewDoor(url string) *Door {
	door := &Door{Url: url}
	door.Id = 1
	door.Connect()

	return door
}

func (d *Door) Connect() {
	d.conn, _ = net.Dial("tcp", d.Url)
}

func (d *Door) ForceOpen() {
	cmd := newCommand(d.Id)
	cmd.Build(0x5E, []byte{})
	d.conn.Write(cmd.bytes)

	response := make([]byte, 25)
	d.conn.Read(response)

	log.Printf("%#x\n", response)
}

func (d *Door) GetInformation() {
	cmd := newCommand(d.Id)
	cmd.Build(0x30, []byte{})
	d.conn.Write(cmd.bytes)

	response := make([]byte, 29)
	d.conn.Read(response)

	stx := response[0]
	channel := response[1:5]
	ack := response[5]
	status := response[6]
	length := response[7:9]
	data := response[9:28]

	log.Printf("%#x\n", stx)
	log.Printf("%#x\n", channel)
	log.Printf("%#x\n", ack)
	log.Printf("%#x\n", status)
	log.Printf("%#x\n", length)
	log.Printf("%#x\n", data)

	log.Println()
	log.Printf("%i\n", uint8(data[12]))
}
