package anviz

import "testing"

func TxestDoorInformation(t *testing.T) {
	door := NewDoor("192.168.2.9:5010")
	door.GetInformation()
}
