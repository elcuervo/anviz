package anviz

import (
	"os"
	"testing"
)

func TestDoorInformation(t *testing.T) {
	door := NewDoor(os.Getenv("DOOR_URL"))
	door.GetInformation()
}

func TestForceOpen(t *testing.T) {
	door := NewDoor(os.Getenv("DOOR_URL"))
	door.ForceOpen()
}
