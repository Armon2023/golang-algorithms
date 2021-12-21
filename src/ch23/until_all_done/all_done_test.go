package untilalldone_test

/* 不使用wg的方式 */
import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	// 这里使用buffered channel
	//ch := make(chan string) 当无消费者接受数据时, 会阻塞剩余未返回的协程, 导致协程泄露
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for i := 0; i < numOfRunner; i++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(1 * time.Second)
	t.Log("After:", runtime.NumGoroutine())
}
