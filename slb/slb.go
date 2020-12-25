package slb

import (
	"context"
	"errors"
	"sync"
	"time"
)

type SLB struct {
	hosts         []*Host
	IndexChan     chan int
	hostCount     int
	indexChanMax  int
	addIndexTime  time.Duration
	lock          sync.RWMutex
	requestCount  int
	isRun         bool
	addIndexLock  chan int
	addIndexStart chan int
}

type Host struct {
	Name string
	Ip   string
}

func NewSLB(hosts []*Host, indexChanMax int, addIndexTime time.Duration) *SLB {
	s := &SLB{}
	s.hostCount = len(hosts)
	s.hosts = hosts
	s.indexChanMax = indexChanMax
	s.IndexChan = make(chan int, s.indexChanMax)
	s.addIndexTime = addIndexTime
	s.requestCount = 0
	s.isRun = false
	s.addIndexLock = make(chan int)
	s.addIndexStart = make(chan int, 1)
	return s
}

func (s *SLB) AddHost(host *Host) {
	s.lock.Lock()
	s.addIndexLock <- 1
	defer func() {
		s.addIndexStart <- 1
		s.lock.Unlock()
	}()
	s.hostCount++
	s.hosts = append(s.hosts, host)
	s.subSomeIndex(len(s.IndexChan))
}

func (s *SLB) RemoveHost(removeHost *Host) {
	s.lock.Lock()
	s.addIndexLock <- 1
	defer func() {
		s.addIndexStart <- 1
		s.lock.Unlock()
	}()
	s.hostCount--
	for index, host := range s.hosts {
		if host == removeHost {
			s.hosts = append(s.hosts[:index], s.hosts[index+1:]...)
		}
	}
	s.subSomeIndex(len(s.IndexChan))
}

func (s *SLB) subSomeIndex(count int) {
	for i := 0; i < count; i++ {
		if len(s.IndexChan) == 0 {
			break
		}
		<-s.IndexChan
	}
}

func (s *SLB) Run(ctx context.Context) error {
	if err := s.isRunning(); err != nil {
		return err
	}
	ticket := time.NewTicker(s.addIndexTime)
	defer func() {
		ticket.Stop()
		s.isRun = false
		close(s.IndexChan)
		close(s.addIndexLock)
		close(s.addIndexStart)
	}()
	for {
		select {
		case <-ticket.C:
			s.addIndex()
		case <-s.addIndexLock:
			select {
			case <-s.addIndexStart:
				s.addIndex()
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (s *SLB) addIndex() {
	need := s.indexChanMax - len(s.IndexChan)
	for i := 0; i < need; i++ {
		s.requestCount++
		s.IndexChan <- s.requestCount % s.hostCount
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
	if i < s.hostCount {
		return s.hosts[i], nil
	}
	return nil, errors.New("error: no found")
}
