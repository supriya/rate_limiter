package main

import (
	"fmt"
	"time"
)

type Request struct {
	requestTime int // probably not needed.
}

type SlidingWindow struct {
	windowDuration int
	windowCapacity int
	queue          []int
}

func (sw *SlidingWindow) allowRequest(request Request) bool {
	pReq := sw.processedRequestCount(request)
	if pReq >= sw.windowCapacity {
		return false
	}
	sw.queue = append(sw.queue, request.requestTime)
	return true
}

func (sw *SlidingWindow) cleanupQueue(newFirstElement int) {
	sw.queue = sw.queue[newFirstElement:]
}

func (sw *SlidingWindow) processedRequestCount(currentRequest Request) int {
	qLen := len(sw.queue)
	if qLen == 0 {
		return 0
	}
	var processedRequestCount int
	currentWindow := currentRequest.requestTime - sw.windowDuration
	ret := sw.findWindowStart(0, qLen-1, currentWindow)
	if ret == -1 {
		sw.queue = []int{}
		return 0
	}

	windowStart := ret + 1
	if len(sw.queue) > sw.windowCapacity {
		sw.cleanupQueue(windowStart)
	}
	processedRequestCount = qLen - int(windowStart)
	return processedRequestCount

}

// return the first element that is less than currentWindow
func (sw *SlidingWindow) findWindowStart(begin, end, currentWindow int) int {
	if begin == end {
		return begin
	}
	if begin > end {
		return -1
	}
	median := (end-begin)/2 + begin
	if sw.queue[median] < currentWindow && median != end && sw.queue[median+1] >= currentWindow {
		return median
	} else if sw.queue[median] < currentWindow {
		begin = median + 1
		return sw.findWindowStart(begin, end, currentWindow)
	} else {
		end = median - 1
		return sw.findWindowStart(begin, end, currentWindow)
	}
}

func main() {
	windowDurationInMinutes, _ := time.ParseDuration("10s")
	slidingWindow := SlidingWindow{windowDuration: int(windowDurationInMinutes.Nanoseconds()),
		windowCapacity: 2, queue: []int{}}

	accepted := 0
	called := 0
	for startTime := time.Now(); time.Now().Sub(startTime) < time.Second*60; called++ {
		request := Request{requestTime: int(time.Now().UnixNano())}
		if slidingWindow.allowRequest(request) {
			accepted++
		}
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("Called", called)
	fmt.Println("Accepted", accepted)
}
