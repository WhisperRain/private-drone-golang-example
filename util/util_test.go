package util

import (
	"fmt"
	"math"
	"testing"
	"time"
)



func TestMaxTimeOfCI(t *testing.T) {
	t.Log("begin")

	for  i:=0;i<math.MaxInt32;i++ {
		time.Sleep(time.Second*1)
		t.Log("the ",i," second")
	}

	fmt.Println("hello world 2")
	if HelloWorld() != "hello world" {
		t.Errorf("got %s expected %s", HelloWorld(), "hello world")
	}
}
