package simpleTask

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var ()

type task_t struct {
	on_pause    int32
	pause_cond  *sync.Cond
	pause_mutex *sync.Mutex
}

func new(m *sync.Mutex, c *sync.Cond) *task_t {
	return &task_t{
		pause_cond:  c,
		pause_mutex: m,
	}
}

func (t *task_t) simpe_task() {
	for i := 0; i < 30; i++ {
		fmt.Printf("Simple task 1\n")
		t.pause_mutex.Lock()
		for atomic.LoadInt32((*int32)(&t.on_pause)) == 1 {
			t.pause_cond.Wait()
		}
		t.pause_mutex.Unlock()
		time.Sleep(time.Second * time.Duration(1))
	}
}

func (t *task_t) puase() {
	atomic.AddInt32((*int32)(&t.on_pause), 1)
}

func (t *task_t) resume() {
	atomic.AddInt32(&t.on_pause, 0)
	t.pause_cond.Signal()
}
