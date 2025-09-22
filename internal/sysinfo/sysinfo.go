// Package sysinfo provides function for gathering system info
package sysinfo

import (
	"context"
	"sync"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func Fetch(args []entities.Arg) ([]entities.Arg, error) {
	wg := sync.WaitGroup{}
	wg.Add(len(args))

	ctx, cancel := context.WithCancel(context.Background())

	errCh := make(chan error, 1)
	defer close(errCh)

	for i := range args {
		go func() {
			defer wg.Done()

			select {
			case <-ctx.Done():

			default:
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

func ResolveIcon(logoName string) (entities.LogoInfo, error) {
	logo, err := entities.GetLogo(logoName)
	if err != nil {
		return entities.LogoInfo{}, err
	}

	return logo, err
}
