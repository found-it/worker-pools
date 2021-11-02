package pool

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type WorkerPool struct {
	num     int
	workers []Worker
}

type Worker struct {
	stopper chan struct{}
	id      int
	logger  *logrus.Entry
}

func NewWorkerPool(n int) *WorkerPool {

	w := WorkerPool{
		num:     n,
		workers: make([]Worker, n),
	}

	for i := 0; i < w.num; i++ {
		w.workers[i] = Worker{
			id:     i,
			logger: logrus.WithField("worker", i),
		}
	}
	return &w
}

func (wp *WorkerPool) WorkOn(queue *[]string) error {

	work := make(chan string, 10)
	for _, f := range *queue {
		work <- f
	}

	close(work)

	var wg sync.WaitGroup
	wg.Add(wp.num)

	for i := range wp.workers {
		go wp.workers[i].run(work, &wg)
	}

	wg.Wait()

	return nil
}

func (w *Worker) run(queue chan string, wg *sync.WaitGroup) {

	for item := range queue {
		select {
		case <-w.stopper:
			w.logger.Info("Stopping")
			wg.Done()
			return
		default:
			w.logger.Info(item)
		}
	}
	wg.Done()
}
