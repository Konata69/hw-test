package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := getOutChannel(in, done)

	for _, stage := range stages {
		out = stage(getOutChannel(out, done))
	}

	return out
}

func getOutChannel(in In, done In) Out {
	out := make(Bi)

	go func(out Bi) {
		defer close(out)
		for {
			select {
			case _, ok := <-done:
				if !ok {
					return
				}
			case val, ok := <-in:
				if !ok {
					return
				}
				out <- val
			}
		}
	}(out)

	return out
}
