package slb

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type SLB struct {
	hosts 			[]*Host
	Channel 		chan int
	count 			int
	max 			int
	time 			time.Duration
	lock 			sync.RWMutex
	requestCount 	int
	isRun			bool
}

type Host struct {
	Name string
	Ip string
}

func NewSLB(hosts []*Host, max int, time time.Duration) *SLB {
	s := &SLB{}
	s.count = len(hosts)
	s.hosts = hosts
	s.max = max
	s.Channel = make(chan int, s.max)
	s.time = time
	s.requestCount = 0
	s.isRun = false
	return s
}

func (s *SLB) Add(host *Host) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.count++
	s.hosts = append(s.hosts, host)
	s.subSome()
}

func (s *SLB) subSome() {
	subCount := len(s.Channel) * 3 / 4
	fmt.Println(subCount)
	for i := 0; i < subCount; i++ {
		<- s.Channel
	}
}

func (s *SLB) Run(ctx context.Context) error {
	if err := s.isRunning(); err != nil {
		return err
	}
	ticket := time.NewTicker(s.time)
	defer ticket.Stop()
	for {
		select {
		case <-ticket.C:
			need := s.max - len(s.Channel)
			for i := 0; i < need; i++ {
				s.requestCount++
				s.Channel <- s.requestCount % s.count
			}
			continue
		case <- ctx.Done():
			s.isRun = false
			close(s.Channel)
			return nil
		}
	}
}

func (s *SLB) isRunning() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isRun {
		s.isRun = true
	} else {
		return errors.New("error: is running")
	}
	return nil
}

func (s *SLB) GetHost(i int) (*Host, error) {
	if i < s.count {
		return s.hosts[i], nil
	}
	return nil, errors.New("error: no found")
}
