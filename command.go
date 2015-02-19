package anviz

import (
	"encoding/binary"
	"sort"
)

const ANVIZ = 0xA5

type command struct {
	id    uint16
	bytes []byte
}

func newCommand(id uint16) *command {
	cmd := &command{}
	cmd.id = id
	cmd.bytes = []byte{ANVIZ}

	return cmd
}

func (c *command) add(b []byte) {
	c.bytes = append(c.bytes, b...)
}

type ReverseByteSorter []byte

func (p ReverseByteSorter) Len() int           { return len(p) }
func (p ReverseByteSorter) Less(i, j int) bool { return p[i] > p[j] }
func (p ReverseByteSorter) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p ReverseByteSorter) Sort() { sort.Sort(p) }

func (c *command) Build(cmd uint16, data []byte) {
	var dst = make([]byte, 4)
	var dataLen = make([]byte, 2)
	var cmdBytes = make([]byte, 2)

	binary.LittleEndian.PutUint16(dst, c.id)
	binary.LittleEndian.PutUint16(cmdBytes, cmd)
	binary.LittleEndian.PutUint16(dataLen, uint16(len(data)))

	sort.Sort(sort.Reverse(ReverseByteSorter(dst)))

	c.add(dst)
	c.add([]byte{cmdBytes[0]})
	c.add(dataLen)

	crch, crcl := checksum(c.bytes)

	c.add([]byte{crch})
	c.add([]byte{crcl})
}
