package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.

	// стейджы для одного значения должны выполняться последовательно
	// но обработка всех стейджей всех значений должна занимать меньше времени, чем последовательная обработка

	out := getOut(in)

	for _, stage := range stages {
		select {
		default:
			out = stage(getOut(out))
		}
	}

	return out
}

func getOut(in In) Out {
	bi := make(Bi)
	go func() {
		for {
			select {
			case v, _ := <-in:
				bi <- v
			}
		}
	}()

	return bi
}
