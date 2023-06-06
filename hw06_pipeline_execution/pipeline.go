package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := GetReadOnlyChannel(in)

	for _, stage := range stages {
		out = stage(GetReadOnlyChannel(out))
	}

	return out
}

func GetReadOnlyChannel(in In) Out {
	tmpChannel := make(Bi)

	go func(tmpChannel Bi) {
		defer close(tmpChannel)
		for {
			select {
			case val, ok := <-in:
				if !ok {
					return
				}
				tmpChannel <- val
			}
		}
	}(tmpChannel)

	return tmpChannel
}
