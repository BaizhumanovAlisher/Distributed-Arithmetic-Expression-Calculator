package agent

import (
	"distributed_calculator/model"
	"time"
)

type Calculator struct {
	miniCalc *model.MiniCalculator
	taskChan chan *model.LeastExpression
	done     chan bool
}

func NewCalculator(id int) *Calculator {
	return &Calculator{
		miniCalc: model.NewMiniCalculator(id),
		taskChan: make(chan *model.LeastExpression),
		done:     make(chan bool),
	}
}

func (c *Calculator) Start() {
	defer func() {
		if r := recover(); r != nil {
			// Restart the worker if it panics
			c.Start()
		}
	}()

	for task := range c.taskChan {
		c.miniCalc.LeastExpression = task // Store the current task

		c.SolveExpression(task)
	}
}

func (c *Calculator) GetCurrentMiniCalculator() *model.MiniCalculator {
	return c.miniCalc
}

// Close closes the worker's task channel and waits for all tasks to complete.
func (c *Calculator) Close() {
	close(c.taskChan)
	<-c.done
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
