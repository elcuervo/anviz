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
