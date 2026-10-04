package pattern

import (
	"fmt"
	"sync"
)

type singleInstance0 struct {
}

var instance0 *singleInstance0

func init() {
	instance0 = &singleInstance0{}
}

func getSingleton0() *singleInstance0 {
	return instance0
}

type singleInstance struct {
}

var (
	singleton *singleInstance
	once      sync.Once
)

func GetSingleton() *singleInstance {
	once.Do( // функция ниже выполнится только один раз
		func() {
			// инициализируем объект
			fmt.Println("Инициализируем singleton")
			singleton = &singleInstance{}
		})

	return singleton
}
