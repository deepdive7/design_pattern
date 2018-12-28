package structural

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
	fileName string
}

func (ri *RealImage) LoadFromDisk() {
	fmt.Println("Loading ", ri.fileName)
}

func (ri *RealImage) Display() {
	fmt.Println("Displaying ", ri.fileName)
}

type ProxyImage struct {
	realImage RealImage
	fileName  string
	loaded    bool
}

func (pI *ProxyImage) SetProxy(ri RealImage) {
	pI.realImage = ri
	pI.fileName = ri.fileName
	pI.loaded = false
}
func (pI *ProxyImage) Display() {
	if !pI.loaded {
		pI.realImage.LoadFromDisk()
		pI.loaded = true
	}
	pI.realImage.Display()
}
