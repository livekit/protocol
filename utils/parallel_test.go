package utils

import (
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParallel(t *testing.T) {
	for _, cs := range []int{1, 10, 100, 10000} {
		t.Run(strconv.Itoa(cs), func(t *testing.T) {
			sendSlice := make([]int, 0, cs)
			recvSlice := make([]int, 0, cs)
			recvCh := make(chan int, cs)
			for i := 0; i < cs; i++ {
				sendSlice = append(sendSlice, i)
			}

			ParallelExec(sendSlice, 10, 2, func(i int) {
				recvCh <- i
			})
			close(recvCh)

			for i := range recvCh {
				recvSlice = append(recvSlice, i)
			}

			sort.Ints(recvSlice)

			require.Equal(t, sendSlice, recvSlice)
		})
	}
}
