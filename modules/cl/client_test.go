package cl

import (
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	ClientInit()
	go Broadcaster()

	c1 := make(chan string, 10000)
	client1 := Client{
		Ch:   c1,
		Type: "test",
	}
	client1.Connect()
	go func() {
		for {
			m := <-client1.Ch
			t.Logf("chan1 receive:%v\n", m)
		}
	}()

	c2 := make(chan string, 100000)
	client2 := Client{
		Ch:   c2,
		Type: "test",
	}
	client2.Connect()
	go func() {
		for {
			m := <-client2.Ch
			t.Logf("chan2 receive:%v\n", m)
		}
	}()

	c3 := make(chan string, 10000)
	client3 := Client{
		Ch:   c3,
		Type: "test1",
	}
	client3.Connect()
	go func() {
		for {
			m := <-client3.Ch
			t.Errorf("chan3 receive:%v\n", m)
		}
	}()

	m := Message{
		M:    "Hello",
		Type: "test",
	}

	for i := 0; i < 10; i++ {
		Send(m)
	}

	time.Sleep(5*time.Second)
}
