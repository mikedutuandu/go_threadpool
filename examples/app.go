package main

import (
	"fmt"
	"threadpool"
	"time"
)

func main() {
	pool := threadpool.NewThreadPool(1, 1)
	task := &myTask{ID: 1}
	task2 := &myTask{ID: 2}
	task3 := &myTask{ID: 3}
	task4 := &myTask{ID: 4}
	task5 := &myTask{ID: 5}
	pool.Execute(task)
	pool.Execute(task2)
	pool.Execute(task3)
	pool.Execute(task4)
	pool.Execute(task5)
	println(pool)
	time.Sleep(5 * time.Hour)
}

type myTask struct {
	ID int64
}

func (m *myTask) Run() {
	fmt.Println("----Running my task ", m.ID)
}
