package agent

import (
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
	"errors"
	"sync"
)

type Agent struct {
	mtx         sync.Mutex
	Calculators []*Calculator
	Queue       chan *expression.LeastExpression
}

func NewAgent(countCalculators int) *Agent {
	queue := make(chan *expression.LeastExpression)
	miniCalcs := make([]*Calculator, countCalculators)

	for i := 0; i < countCalculators; i++ {
		miniCalcs[i] = NewCalculator(i, queue)
	}

	return &Agent{
		Calculators: miniCalcs,
		Queue:       queue,
	}
}

func (a *Agent) AddTask(exp *expression.LeastExpression) {
	a.Queue <- exp
}

func (a *Agent) GetStatusAllCalculators() []*model.MiniCalculator {
	miniCalculators := make([]*model.MiniCalculator, len(a.Calculators))

	for i, calc := range a.Calculators {
		miniCalculators[i] = calc.GetCurrentMiniCalculator()
	}

	return miniCalculators
}

// SetNewCountCalculator can add and remove calculators
func (a *Agent) SetNewCountCalculator(count int) error {
	a.mtx.Lock()

	if count < 1 {
		return errors.New("count of calculator must be greater than 0")
	}

	previousCount := len(a.Calculators)

	if count == previousCount {
		return nil
	}

	if count-previousCount > 0 {
		for i := previousCount; i < count; i++ {
			a.Calculators = append(a.Calculators, NewCalculator(len(a.Calculators), a.Queue))
		}
	} else {
		for i := len(a.Calculators) - 1; i >= count; i-- {
			a.Calculators[i].Close()
			<-a.Calculators[i].closed
		}

		a.Calculators = a.Calculators[:count]
	}

	a.mtx.Unlock()
	return nil
}