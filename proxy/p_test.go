package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxy := CreateProxy("tom")
	proxy.show()
}

func TestClient(t *testing.T) {
	client := Client{}
	client.newShow()
}


