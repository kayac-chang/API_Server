package utils

import (
	"context"
	"sync"
)

type Task func() (interface{}, error)
type Promise func() Resolver
type Resolver <-chan Task

func Wait(task Promise) (interface{}, error) {

	return WaitCtx(context.Background(), task)
}

func WaitCtx(ctx context.Context, task Promise) (interface{}, error) {

	select {
	case <-ctx.Done():
		return nil, nil

	case fn := <-task():
		return fn()
	}
}

func WaitAll(tasks ...Promise) ([]interface{}, error) {

	return WaitAllCtx(context.Background(), tasks...)
}

func WaitAllCtx(ctx context.Context, tasks ...Promise) ([]interface{}, error) {
	counts := len(tasks)

	results := make([]interface{}, counts)
	var resErr error

	wg := sync.WaitGroup{}
	wg.Add(counts)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for idx, task := range tasks {

		go func(idx int, task Promise) {

			defer wg.Done()

			res, err := WaitCtx(ctx, task)

			if err != nil {
				resErr = err

				cancel()

				return
			}

			results[idx] = res
		}(idx, task)
	}

	wg.Wait()

	return results, resErr
}

func Promisefy(task Task) Promise {

	return func() Resolver {

		resolver := make(chan Task)

		go func() {
			defer close(resolver)

			res, err := task()

			resolver <- func() (interface{}, error) { return res, err }
		}()

		return resolver
	}
}
