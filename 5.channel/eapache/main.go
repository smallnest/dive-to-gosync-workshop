package main

import (
	"fmt"

	"github.com/eapache/channels"
)

func testPipe() {
	fmt.Println("pipe:")
	a := channels.NewNativeChannel(channels.None)
	b := channels.NewNativeChannel(channels.None)

	channels.Pipe(a, b)
	// channels.WeakPipe(a, b)

	go func() {
		for i := 0; i < 5; i++ {
			a.In() <- i
		}
		a.Close()
	}()

	for v := range b.Out() {
		fmt.Printf("%d ", v)
	}
}

func testDist() {
	fmt.Println("dist:")
	a := channels.NewNativeChannel(channels.None)
	outputs := []channels.Channel{
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
	}

	channels.Distribute(a, outputs[0], outputs[1], outputs[2], outputs[3])
	//channels.WeakDistribute(a, outputs[0], outputs[1], outputs[2], outputs[3])

	go func() {
		for i := 0; i < 5; i++ {
			a.In() <- i
		}
		a.Close()
	}()

	for i := 0; i < 6; i++ {
		var v interface{}
		var j int
		select {
		case v = <-outputs[0].Out():
			j = 0
		case v = <-outputs[1].Out():
			j = 1
		case v = <-outputs[2].Out():
			j = 2
		case v = <-outputs[3].Out():
			j = 3
		}
		fmt.Printf("channel#%d: %d\n", j, v)
	}

}

func testTee() {
	fmt.Println("tee:")
	a := channels.NewNativeChannel(channels.None)
	outputs := []channels.Channel{
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
	}

	channels.Tee(a, outputs[0], outputs[1], outputs[2], outputs[3])
	//channels.WeakTee(a, outputs[0], outputs[1], outputs[2], outputs[3])

	go func() {
		for i := 0; i < 5; i++ {
			a.In() <- i
		}
		a.Close()
	}()

	for i := 0; i < 20; i++ {
		var v interface{}
		var j int
		select {
		case v = <-outputs[0].Out():
			j = 0
		case v = <-outputs[1].Out():
			j = 1
		case v = <-outputs[2].Out():
			j = 2
		case v = <-outputs[3].Out():
			j = 3
		}
		fmt.Printf("channel#%d: %d\n", j, v)
	}
}

func testMulti() {
	fmt.Println("multi:")
	a := channels.NewNativeChannel(channels.None)
	inputs := []channels.Channel{
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
		channels.NewNativeChannel(channels.None),
	}

	channels.Multiplex(a, inputs[0], inputs[1], inputs[2], inputs[3])
	//channels.WeakMultiplex(a, inputs[0], inputs[1], inputs[2], inputs[3])

	go func() {
		for i := 0; i < 5; i++ {
			for j := range inputs {
				inputs[j].In() <- i
			}
		}
		for i := range inputs {
			inputs[i].Close()
		}
	}()

	for v := range a.Out() {
		fmt.Printf("%d ", v)
	}
}

func main() {
	testPipe()
	testDist()
	testTee()
	testMulti()
}
