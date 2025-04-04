package corredor

import (
	"context"
	"fmt"

	"google.golang.org/grpc/connectivity"
)

// Healtcheck for global
func Healthcheck(_ context.Context) error {
	svc := Service()

	if svc == nil {
		return fmt.Errorf("uninitialized")
	}

	if !svc.opt.Enabled {
		return nil
	}

	if state := svc.conn.GetState(); state != connectivity.Ready {
		return fmt.Errorf("connection is %s", state)
	}

	return nil
}
