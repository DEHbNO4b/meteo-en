package semaphore

type Semaphore struct {
	semCh chan struct{}
}

func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		semCh: make(chan struct{}, maxReq),
	}
}

func (s *Semaphore) Acquire() {
	s.semCh <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.semCh
}
