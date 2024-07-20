package singleton

import (
	"fmt"
	"sync"
)

var treadLocker = &sync.Mutex{}

type Singleton struct {
}

var singeltonInstance *Singleton

// GetInstance is a function that returns a Singleton instance
// This function is thread safe
// Singleton is a creational design pattern that lets you ensure that a class has only one instance,
// while providing a global access point to this instance.
func GetInstance() *Singleton {
	if singeltonInstance == nil {
		treadLocker.Lock()
		defer treadLocker.Unlock()

		if singeltonInstance == nil {
			fmt.Println("Creating Singleton Instance")
			singeltonInstance = &Singleton{}
		} else {
			fmt.Println("Singleton Instance already created")
		}
	} else {
		fmt.Println("Singleton Instance already created")
	}

	return singeltonInstance
}
