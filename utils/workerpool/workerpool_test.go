package workerpool

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func TestOrderedWorkerPool(t *testing.T) {
	wp := NewOrderedWorkerPool(10)
	tasks := make(map[string]chan int) // key -> goroutine id
	for i := 0; i < 30; i++ {
		tasks["task"+strconv.Itoa(i)] = make(chan int, 1000)
	}

	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		for key, ch := range tasks {
			k := key
			c := ch
			wg.Add(1)
			wp.Submit(func() {
				c <- goid()
				time.Sleep(5 * time.Millisecond)
				wg.Done()
			}, k)
		}
	}

	wg.Wait()
	for k, ch := range tasks {
		if len(ch) < 30 {
			require.FailNow(t, "no value received", k)
		}
		firstVal := <-ch

	test_equal:
		for {
			select {
			case val := <-ch:
				require.Equal(t, firstVal, val)
			default:
				break test_equal
			}
		}
	}
}
