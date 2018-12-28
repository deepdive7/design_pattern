package structural

import "testing"

func TestProxy(t *testing.T) {
	proxyImage := ProxyImage{}
	proxyImage.SetProxy(RealImage{"imada.jpg"})
	proxyImage.Display()
	proxyImage.Display()
	proxyImage.Display()
}
