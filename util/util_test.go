package util


import (
	"testing"
)

func TestMaxTimeOfCI(t *testing.T) {
	t.Log("begin")

	for  i:=0;i<math.MaxInt32;i++ {
		time.Sleep(time.Second*1)
		t.Log("the ",i," second")
	}
 
}
