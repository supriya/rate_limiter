package main

import (
	"math"
	"time"
)

type TokenBucket struct {
	currentBucketSize float64
	lastRefillTime    int
	refillRate        int
	bucketCapacity    float64
}

func (tb *TokenBucket) allowRequest(request float64) bool {
	tb.refillBucket()
	if tb.currentBucketSize > request {
		tb.currentBucketSize -= request
		return true
	}
	return false
}

func (tb *TokenBucket) refillBucket() {
	currentTime := time.Now().Nanosecond()
	tokensToAdd := (currentTime - tb.lastRefillTime) * tb.refillRate / 1e9
	tb.currentBucketSize = math.Min(float64(tokensToAdd)+tb.currentBucketSize, tb.bucketCapacity)
	tb.lastRefillTime = currentTime
}

func main() {
	tb := TokenBucket{currentBucketSize: 3, lastRefillTime: 0, refillRate: 1, bucketCapacity: 3}
	tb.allowRequest(1)
}
