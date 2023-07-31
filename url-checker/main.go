package main

import (
	"fmt"
	"time"
)

func main() {
	// 이 경우, ella가 먼저 10번 실행되고, moon이 실행된다. - Top-down
	// sexyCount("ella")
	// sexyCount("moon")

	// go sexyCount("ella")
	// sexyCount("moon") // 여기에도 go를 붙이면 main함수가 끝나버리므로 종료된다.

	go sexyCount("ella")
	go sexyCount("moon")
	time.Sleep(time.Second * 5) // 5초간 main함수가 살아있으므로, 5초간 goroutine이 실행된다. 단, 이후 숫자는 프린트 되지 않는다.

	// main함수는 결과를 저장하는곳!!
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
