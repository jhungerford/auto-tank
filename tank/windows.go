//+build windows

package tank

import "log"

func (t Tank) Init() error {
	t.Left.Init()
	t.Right.Init()
	return nil
}

func (t Tread) Init() {
	log.Print("Stub tread init().")
}

func (t Tread) Move(dir TreadDirection) {
	log.Print("Stub tread move().")
}
