package store

import "github.com/pachyderm/pachyderm/src/pps"

type Client interface {
	AddPipelineRun(pipelineRun *pps.PipelineRun) error
	GetPipelineRun(id string) (*pps.PipelineRun, error)
	AddPipelineRunStatus(id string, runStatusType pps.PipelineRunStatusType) error
	GetPipelineRunStatusLatest(id string) (*pps.PipelineRunStatus, error)
	AddPipelineRunContainerIDs(id string, containerIDs ...string) error
	GetPipelineRunContainerIDs(id string) ([]string, error)
}

func NewInMemoryClient() Client {
	return newInMemoryClient()
}

func NewRethinkClient(address string) (Client, error) {
	return newRethinkClient(address)
}
