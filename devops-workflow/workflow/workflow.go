package workflow

type endState int8

const (
	// esUnknown indicates we haven't reached and end state.
	esUnknown endState = 0
	// esSuccess means that the workflow has completed successfully. This
	// does not mean there haven't been failurea.
	esSuccess endState = 1
	// esPreconditionFailure means no work was done as we failed on a precondition.
	esPreconditionFailure endState = 2
	// esCanaryFailure indicates one of the canaries failed, stopping the workflow.
	esCanaryFailure endState = 3
	// esMaxFailures indicates that the workflow passed the canary phase, but failed
	// at a later phase.
	esMaxFailures endState = 4
)

type Workflow struct {
	config   *config
	lb       *client.client
	failures int32
	endState endState
	actions  []*actions
}
