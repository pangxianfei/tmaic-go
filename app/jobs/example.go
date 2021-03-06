package jobs

import (
	"github.com/golang/protobuf/proto"

	"github.com/pangxianfei/framework/helpers/debug"
	"github.com/pangxianfei/framework/job"
	pbs "tmaic/app/jobs/protocol_buffers"
)

func init() {
	job.Add(&ExampleJob{})
}

type ExampleJob struct {
	job.Job
}

func (e *ExampleJob) Retries() uint32 {
	return 3
}

func (e *ExampleJob) Name() string {
	return "example-job"
}

func (e *ExampleJob) ParamProto() proto.Message {
	return &pbs.ExampleJob{}
}

func (e *ExampleJob) Handle(paramPtr proto.Message) error {
	obj := paramPtr.(*pbs.ExampleJob)
	debug.Dump(obj)
	return nil
}
