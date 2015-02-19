package anviz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCRC(t *testing.T) {
	get_record_information := []byte{0xA5, 0x00, 0x00, 0x00, 0x01, 0x3C, 0x00, 0x00}
	h, l := checksum(get_record_information)

	assert.Equal(t, h, uint16(0x49))
	assert.Equal(t, l, uint16(0xA9))
}

func TestCommand(t *testing.T) {
	cmd := newCommand(1)
	cmd.Build(0x3C, []byte{})

	samplePacket := []byte{0xA5, 0x0, 0x0, 0x0, 0x1, 0x3C, 0x0, 0x0, 0x49, 0xa9}

	assert.Equal(t, cmd.bytes, samplePacket)
}

func TxestDowloadStaffInformation(t *testing.T) {
	door := NewDoor("192.168.2.9:5010")
	door.DownloadStaffInformation()
}

func TxestDoorInformation(t *testing.T) {
	door := NewDoor("192.168.2.9:5010")
	door.GetInformation()
}

func TxestRecordInformation(t *testing.T) {
	door := NewDoor("192.168.2.9:5010")
	door.GetRecordInformation()
}
