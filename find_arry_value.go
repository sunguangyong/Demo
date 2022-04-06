package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

)
// 开启多个 goroutine 从一个超大的[]int 里寻找特定值, 找到后或者5秒后未找到推出 所有 goroutine

var  signal int64 =  0

func SearchValues(cancel context.CancelFunc, int_arry []int, start, end, value int){
	for i:= start; i <= end; i ++ {
		if int_arry[i] == value  {
			atomic.AddInt64(&signal,1)
			fmt.Println("find value")
			cancel()
			return
		}

		if  signal > 0 {
			fmt.Println("Has been found value")
			return
		}
	}
}

func GetIntArrys()(intArr [] int){
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <1000000; i++ {
		intArr = append(intArr, i)
	}
	intArr = append(intArr,8)
	return intArr
}

func Paging(total_len, share, num int)  (start, end int) {
	average := total_len / share
	if share != num {
		return (num-1) * average , num * average
	} else {
		return (num-1) * average, total_len - 1
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	var int_arry = GetIntArrys()
	var value int = 2
	var go_num int = 10  // 协程数
	total_len := len(int_arry)
	for i := 1; i <= go_num ; i ++ {
		start,end := Paging(total_len, go_num, i)
		go SearchValues(cancel, int_arry, start,end, value)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
		default:
		}
	}
	time.Sleep(5 * time.Second)
}

