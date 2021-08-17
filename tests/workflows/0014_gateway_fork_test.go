package workflows

import (
	"context"
	"testing"

	autTypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/stretchr/testify/require"
)

func Test0014_gateway_fork(t *testing.T) {
	var (
		ctx = superUser(context.Background())
		req = require.New(t)
	)

	loadScenario(ctx, t)

	for i := 0; i < 101; i++ {
		_, trace := mustExecWorkflow(ctx, t, "case1", autTypes.WorkflowExecParams{})
		req.Len(trace, 9)

		// These steps are fixed; the fork, join, and end
		req.Equal(uint64(10), trace[0].StepID)
		req.Equal(uint64(14), trace[6].StepID)
		req.Equal(uint64(15), trace[7].StepID)
		req.Equal(uint64(0), trace[8].StepID)

		logCount := 0
		for _, f := range trace {
			if f.StepID == 11 || f.StepID == 12 || f.StepID == 13 {
				logCount++
			}
		}
		req.Equal(logCount, 3)

		joinCount := 0
		for _, f := range trace {
			if f.StepID == 14 {
				joinCount++
			}
		}
		req.Equal(joinCount, 3)
	}

}
