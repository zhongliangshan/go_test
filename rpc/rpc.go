package rpcdemo

import "errors"

type DemoRpc struct {
}

type Args struct {
	A, B int
}

func (DemoRpc) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("devisio by zero")
	}

	*result = float64(args.A) / float64(args.B)

	return nil
}
