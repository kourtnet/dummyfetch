// Package sysinfo provides functions for gathering system and logo info
package sysinfo

import (
	"context"
	"sync"

	"github.com/kourtnet/dummyfetch/internal/entities"
	"golang.org/x/sync/errgroup"
)

func fetchArg(argName string, mu *sync.RWMutex) error {
	var err error

	mu.RLock()
	arg := entities.ArgsMap[argName]
	mu.RUnlock()

	if arg.Contents != "" {
		return nil
	}

	arg.Contents, err = arg.Command()
	if err != nil {
		return err
	}

	mu.Lock()
	entities.ArgsMap[argName] = arg
	mu.Unlock()

	return nil
}

func Fetch(argNames []string) error {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	mu := &sync.RWMutex{}

	for _, v := range argNames {
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			return fetchArg(v, mu)
		})
	}

	return g.Wait()
}

func FetchLogo(name string) (string, error) {
	if name == "" {
		var err error
		name, err = entities.GetDistro()
		if err != nil {
			return "", nil
		}
	}

	if _, ok := entities.LogosMap[name]; ok {
		return name, nil
	}

	return entities.BasicLogoName, nil
}
