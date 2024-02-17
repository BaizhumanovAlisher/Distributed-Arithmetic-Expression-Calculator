package agent

import (
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
	"log"
	"math"
	"sync"
	"time"
)

type Calculator struct {
	mtx      sync.Mutex
	miniCalc *expression.MiniCalculator
	taskChan chan *expression.LeastExpression
	isBusy   bool
}

func NewCalculator(i int) *Calculator {
	c := &Calculator{
		miniCalc: expression.NewMiniCalculator(i),
		taskChan: make(chan *expression.LeastExpression),
		isBusy:   false,
	}

	c.Start()

	return c
}

func (c *Calculator) AddTask(task *expression.LeastExpression) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.isBusy {
		return false
	}

	c.taskChan <- task
	return true
}

func (c *Calculator) Start() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				c.Start()
			}
		}()

		for {
			select {
			case task := <-c.taskChan:
				c.mtx.Lock()
				c.miniCalc.LeastExpression = task
				c.isBusy = true
				c.mtx.Unlock()

				c.SolveExpression(task)

				c.mtx.Lock()
				c.miniCalc.LeastExpression = nil
				c.isBusy = false
				c.mtx.Unlock()
			}
		}
	}()
}

func (c *Calculator) GetCurrentMiniCalculator() *expression.MiniCalculator {
	c.mtx.Lock()
	copyMiniCalc := *c.miniCalc
	c.mtx.Unlock()
	return &copyMiniCalc
}

func (c *Calculator) SolveExpression(le *expression.LeastExpression) {
	time.Sleep(time.Duration(le.DurationInSecond) * time.Second)

	switch le.Operation {
	case model.Addition:
		le.Result = le.Number1 + le.Number2
	case model.Subtraction:
		le.Result = le.Number1 - le.Number2
	case model.Multiplication:
		le.Result = le.Number1 * le.Number2
	case model.Division:
		if almostEqual(le.Number2, 0.0) {
			le.ResultIsCorrect <- false
			return
		}

		le.Result = le.Number1 / le.Number2
	}

	le.ResultIsCorrect <- true
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
