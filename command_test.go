package anviz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	cmd := newCommand(1)
	cmd.Build(0x3C, []byte{})

	samplePacket := []byte{0xA5, 0x0, 0x0, 0x0, 0x1, 0x3C, 0x0, 0x0, 0x49, 0xa9}

	assert.Equal(t, cmd.bytes, samplePacket)
}
