package agent

import (
	"api_server/internal/model"
	expression2 "api_server/internal/model/expression"
	"log"
	"math"
	"time"
)

type Calculator struct {
	miniCalc *expression2.MiniCalculator
	taskChan chan *expression2.LeastExpression
	isBusy   bool
}

func NewCalculator(i int) *Calculator {
	c := &Calculator{
		miniCalc: expression2.NewMiniCalculator(i),
		taskChan: make(chan *expression2.LeastExpression),
		isBusy:   false,
	}

	c.Start()

	return c
}

func (c *Calculator) AddTask(task *expression2.LeastExpression) bool {
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
				c.miniCalc.LeastExpression = task
				c.isBusy = true

				c.SolveExpression(task)

				c.miniCalc.LeastExpression = nil
				c.isBusy = false
			}
		}
	}()
}

func (c *Calculator) GetCurrentMiniCalculator() *expression2.MiniCalculator {
	return c.miniCalc
}

func (c *Calculator) SolveExpression(le *expression2.LeastExpression) {
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
