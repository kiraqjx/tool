package slb

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSLB(t *testing.T) {
	host1 := &Host{Name: "host1.com", Ip: "192.168.1.1"}
	host2 := &Host{Name: "host2.com", Ip: "192.168.1.2"}
	hosts := []*Host{host1, host2}
	slb := NewSLB(hosts, 10, 2 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	go slb.Run(ctx)
	go func(slb *SLB, host *Host) {
		time.Sleep(3 * time.Second)
		host3 := &Host{Name: "host3.com", Ip: "192.168.1.3"}
		slb.AddHost(host3)
		slb.RemoveHost(host1)
	}(slb, host1)
	for i := 0; i < 5; i++ {
		index := <- slb.IndexChan
		fmt.Println(slb.GetHost(index))
	}
	time.Sleep(1 * time.Second)
	fmt.Println("==========================")
	for i := 0; i < 15; i++ {
		index := <-slb.IndexChan
		fmt.Println(slb.GetHost(index))
	}
	cancel()
}
