/*
Необходимо написать функцию для параллельного выполнения заданий в 'n' параллельных горутинах:
	- кол-во создаваемых горутин не должно зависеть от числа заданий, т.е функция должна запускать 'n' горутин для конкурентной обработки заданий и
	возможно, еще несколько вспомогательных горутин;
	- функция должна остановить свою работу, если произошло 'm' ошибок;
	- после завершения работы функции (успешного или из-за превышения 'm') не должно оставаться работающих горутин;
	- нужно учесть, что задания могут выполняться разное время, а длина списка задач может быть больше или меньше 'n';
	- значение 'm' всегда больше 0;
Граничные случаи:
	- если задачи работают без ошибок, то выпонятся все задачи;
	- если в первых выполненных 'm' задачах (или вообще всех) происходят ошибки, то всего выполнится не более 'm + n' задач;
*/

package tasks

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Task1() {
	tasks := make([]Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = func() error {
			time.Sleep(time.Millisecond * 50)
			return nil
		}
	}

	_ = run(tasks, 5, 10)
	fmt.Println("done")
}

func run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var errCounter int32
	taskChan := make(chan Task)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for task := range taskChan {
				if err := task(); err != nil {
					atomic.AddInt32(&errCounter, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errCounter) >= int32(m) {
			break
		}

		taskChan <- task
	}

	close(taskChan)
	wg.Wait()

	if errCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
