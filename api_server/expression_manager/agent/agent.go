package agent

import (
	expression2 "distributed_calculator/internal/model/expression"
	"sync"
)

type Agent struct {
	mtxToAddTask sync.Mutex
	Calculators  []*Calculator
	Queue        chan *expression2.LeastExpression
}

func NewAgent(countCalculators int) *Agent {
	queue := make(chan *expression2.LeastExpression)
	miniCalculators := make([]*Calculator, countCalculators)

	for i := 0; i < countCalculators; i++ {
		miniCalculators[i] = NewCalculator(i)
	}

	a := &Agent{
		Calculators: miniCalculators,
		Queue:       queue,
	}

	go a.distributeTasks()

	return a
}

func (a *Agent) AddTask(exp *expression2.LeastExpression) {
	a.mtxToAddTask.Lock()
	a.Queue <- exp
	a.mtxToAddTask.Unlock()
}

func (a *Agent) GetAllMiniCalculators() []*expression2.MiniCalculator {
	miniCalculators := make([]*expression2.MiniCalculator, len(a.Calculators))

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
