// Package sysinfo provides function for gathering system info
package sysinfo

import (
	"context"
	"sync"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Fetch(argNames []string) ([]entities.Arg, error) {
	if len(argNames) == 0 {
		argNames = entities.BasicArgs
	}

	args := make([]entities.Arg, len(argNames))

	wg := sync.WaitGroup{}
	wg.Add(len(args))

	ctx, cancel := context.WithCancel(context.Background())

	errCh := make(chan error, 1)
	defer close(errCh)

	for i := range argNames {
		go func() {
			defer wg.Done()

			select {
			case <-ctx.Done():

			default:
				args[i] = entities.ArgsMap[argNames[i]]

				var err error
				args[i].Contents, err = args[i].Command()
				if err != nil {
					select {
					case errCh <- err:
						cancel()

					default:
					}
				}
			}
		}()
	}

	wg.Wait()
	select {
	case err := <-errCh:
		return nil, err
	default:
	}

	return args, nil
}

func FetchLogo(name string) (entities.LogoInfo, error) {
	if name == "" {
		var err error
		name, err = entities.GetDistro()
		if err != nil {
			return entities.LogoInfo{}, nil
		}
	}

	if logo, ok := entities.LogosMap[name]; ok {
		return logo, nil
	}

	return entities.BasicLogo, nil
}
