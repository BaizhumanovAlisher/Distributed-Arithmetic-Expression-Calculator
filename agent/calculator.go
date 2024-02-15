package agent

import (
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
	"math"
	"time"
)

type Calculator struct {
	miniCalc *expression.MiniCalculator
	taskChan chan *expression.LeastExpression
	closed   chan bool
}

func NewCalculator(id int, queue chan *expression.LeastExpression) *Calculator {
	return &Calculator{
		miniCalc: expression.NewMiniCalculator(id),
		taskChan: queue,
		closed:   make(chan bool),
	}
}

func (c *Calculator) Start() {
	defer func() {
		if r := recover(); r != nil {
			c.Start()
		}
	}()

	for task := range c.taskChan {
		if c.miniCalc.LeastExpression != nil {
			c.miniCalc.LeastExpression = task // Store the current task
		}

		c.SolveExpression(task)

		c.miniCalc.LeastExpression = nil // Reset the current task
	}
}

func (c *Calculator) GetCurrentMiniCalculator() *expression.MiniCalculator {
	return c.miniCalc
}

func (c *Calculator) Close() {
	close(c.taskChan)
	<-c.closed
}

func (c *Calculator) SolveExpression(le *expression.LeastExpression) {
	time.Sleep(time.Duration(int64(le.DurationInSecond) * int64(time.Second)))

	le.ResultIsCorrect = make(chan bool, 1)
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
