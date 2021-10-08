package main

import (
	"fmt"
	"car-main/headlights"
	"car-main/stereo"
	"car-main/wheels"
)

func main()  {
	OpenDoor()
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()
}
func OpenDoor() {
	fmt.Println("Opening door")
}
