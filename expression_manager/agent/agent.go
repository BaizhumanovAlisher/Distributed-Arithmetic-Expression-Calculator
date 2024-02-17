package agent

import (
	"distributed_calculator/model/expression"
	"sync"
)

type Agent struct {
	mtxToAddTask sync.Mutex
	Calculators  []*Calculator
	Queue        chan *expression.LeastExpression
}

func NewAgent(countCalculators int) *Agent {
	queue := make(chan *expression.LeastExpression)
	miniCalcs := make([]*Calculator, countCalculators)

	for i := 0; i < countCalculators; i++ {
		miniCalcs[i] = NewCalculator(i)
	}

	a := &Agent{
		Calculators: miniCalcs,
		Queue:       queue,
	}

	go a.distributeTasks()

	return a
}

func (a *Agent) AddTask(exp *expression.LeastExpression) {
	a.mtxToAddTask.Lock()
	a.Queue <- exp
	a.mtxToAddTask.Unlock()
}

func (a *Agent) GetAllMiniCalculators() []*expression.MiniCalculator {
	miniCalculators := make([]*expression.MiniCalculator, len(a.Calculators))

	for i, calc := range a.Calculators {
		miniCalculators[i] = calc.GetCurrentMiniCalculator()
	}

	return miniCalculators
}

func (a *Agent) distributeTasks() {
	i := 0
	countOfCalculators := len(a.Calculators)
	for task := range a.Queue {
		// Try to send the task to each worker in turn
		for {
			// Use a select statement with a default case to avoid blocking
			if a.Calculators[i].AddTask(task) {
				break
			}

			i++

			if i < 0 || i >= countOfCalculators {
				i = 0
			}
		}
	}
}
