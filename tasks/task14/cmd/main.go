package main

import (
	"fmt"
	"time"
)

// Реализовать функцию, которая будет объединять один или более каналов done (каналов сигнала завершения) в один.
// Возвращаемый канал должен закрываться, как только закроется любой из исходных каналов.
// В этом примере канал, возвращённый or(...), закроется через ~1 секунду,
// потому что самый короткий канал sig(1*time.Second) закроется первым.
//  Ваша реализация or должна уметь принимать на вход произвольное число каналов и завершаться при сигнале
//  на любом из них.

// Подсказка: используйте select в бесконечном цикле для чтения из всех каналов одновременно, либо рекурсивно объединяйте каналы попарно.

func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		c := make(chan interface{})
		close(c)
		return c
	}
	resultFetcher := make(chan interface{})

	go func() {
		defer close(resultFetcher)
		done := make(chan struct{}, 1)

		for _, ch := range channels {
			go func(c <-chan interface{}) {
				<-c
				select {
				case done <- struct{}{}:
				default:
				}
			}(ch)
		}
		<-done
	}()
	return resultFetcher
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	result := or(
		sig(1*time.Second),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	<-result

	fmt.Printf("done after %v\n", time.Since(start))
}
