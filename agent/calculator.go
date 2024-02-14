package agent

import (
	"distributed_calculator/model"
	"time"
)

type Calculator struct {
	miniCalc *model.MiniCalculator
	taskChan chan *model.LeastExpression
	closed   chan bool
}

func NewCalculator(id int, queue chan *model.LeastExpression) *Calculator {
	return &Calculator{
		miniCalc: model.NewMiniCalculator(id),
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

func (c *Calculator) GetCurrentMiniCalculator() *model.MiniCalculator {
	return c.miniCalc
}

func (c *Calculator) Close() {
	close(c.taskChan)
	<-c.closed
}

func (c *Calculator) SolveExpression(le *model.LeastExpression) {
	time.Sleep(time.Duration(int64(le.DurationInSecond) * int64(time.Second)))

	switch le.Operation {
	case model.Addition:
		le.Result = le.Number1 + le.Number2
	case model.Subtraction:
		le.Result = le.Number1 - le.Number2
	case model.Multiplication:
		le.Result = le.Number1 * le.Number2
	case model.Division:
		le.Result = le.Number1 / le.Number2
	}

	le.ResultIsReady = make(chan bool, 1)
	le.ResultIsReady <- true
}
