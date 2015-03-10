package anviz

import "testing"

func TestDoorInformation(t *testing.T) {
	door := NewDoor("192.168.2.9:5010")
	door.GetInformation()
}

func TestForceOpen(t *testing.T) {
	door := NewDoor("192.168.2.9:5010")
	door.ForceOpen()
}
