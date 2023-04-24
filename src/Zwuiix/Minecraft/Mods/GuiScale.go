package Mods

import (
	"erodialuncher/src/Zwuiix/Game"
	"github.com/kbinani/win"
	"unsafe"
)

type GuiScale struct{}

func (r GuiScale) SetGuiScale(h *Game.Handler, number float64) {
	value := float32(number)
	var num win.DWORD
	var bytesWritten win.SIZE_T

	addresses := []uintptr{0x3F74100}

	for i := 0; i < len(addresses); i++ {
		address := win.LPVOID(h.GameID() + addresses[i])
		win.VirtualProtectEx(h.Handle(), address, 4, 0x40, &num)
		win.WriteProcessMemory(h.Handle(), address, uintptr(unsafe.Pointer(&value)), 4, &bytesWritten)
	}

}
