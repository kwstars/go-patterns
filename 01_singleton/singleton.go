package singleton

import (
	"fmt"
	"sync"
)

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			fmt.Println("Create Single Instance Now")
			instance = &singleton{}
		})
	} else {
		fmt.Println("Single Instance already created")
	}

	return instance
}
