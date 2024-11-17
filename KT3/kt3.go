package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
)

type Zebra struct {
	HighSpeed int
	Size      int
	Stripes   int
}

type Tiger struct {
	HighSpeed     int
	ClimbTree     bool
	StealthRating int
}

type Panda struct {
	Size          int
	BambooEaten   int
	RecognizeDiseases bool
}

func processZebra(wg *sync.WaitGroup) {
	defer wg.Done()
	zebra := Zebra{HighSpeed: 60, Size: 250, Stripes: 80}
	fmt.Printf("Зебра - Скорость: %d км/ч, Размер: %d кг, Полоски: %d\n", zebra.HighSpeed, zebra.Size, zebra.Stripes)

	go func() {
		beeep.Notify("Зебра", "Обработана информация о зебре", "")
	}()
	time.Sleep(1 * time.Second) 
}

func processTiger(wg *sync.WaitGroup) {
	defer wg.Done()
	tiger := Tiger{HighSpeed: 80, ClimbTree: true, StealthRating: 9}
	fmt.Printf("Тигр - Скорость: %d км/ч, Лазает по деревьям: %v, Рейтинг скрытности: %d/10\n", tiger.HighSpeed, tiger.ClimbTree, tiger.StealthRating)

	go func() {
		beeep.Notify("Тигр", "Обработана информация о тигре", "")
	}()
	time.Sleep(1 * time.Second) 
}

func processPanda(wg *sync.WaitGroup) {
	defer wg.Done()
	panda := Panda{Size: 150, BambooEaten: 20, RecognizeDiseases: true}
	fmt.Printf("Панда - Размер: %d кг, Съедено бамбука: %d кг, Распознаёт болезни: %v\n", panda.Size, panda.BambooEaten, panda.RecognizeDiseases)

	go func() {
		beeep.Notify("Панда", "Обработана информация о панде", "")
	}()
	time.Sleep(1 * time.Second) 
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go processZebra(&wg)
	go processTiger(&wg)
	go processPanda(&wg)

	wg.Wait()
	fmt.Println("Все животные обработаны.")
}
