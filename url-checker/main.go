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

	// go sexyCount("ella")
	// go sexyCount("moon")
	// time.Sleep(time.Second * 5) // 5초간 main함수가 살아있으므로, 5초간 goroutine이 실행된다. 단, 이후 숫자는 프린트 되지 않는다.

	// main함수는 결과를 저장하는곳!!

	// c := make(chan bool) // 채널 생성
	c := make(chan string) // 채널2 생성

	people := [2]string{"ella", "moon"}
	for _, person := range people {
		go isSexy(person, c)
	}

	fmt.Println("Waiting..")
	// 채널에서 값을 받을 때까지 대기한다.
	// resultOne := <-c   // 채널에서 받은 첫번째 메시지
	// resultTwo := <-c   // 채널에서 받은 두번째 메시지
	// resultThree := <-c // 채널에서 받은 세번째 메시지

	// fmt.Println("Received", resultOne)
	// fmt.Println("Received", resultTwo)
	// fmt.Println("Received", resultThree) // 3개를 받으려고 하면 메시지를 대기하고 있으나 goroutines은 이미 끝났으므로(2개출력) deadlock이 발생한다. 채널은 blocking operation이다.

	// 몇개의 메시지가 올 지 모르기떄문에 Loof를 사용한다.
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy" // 채널에 값을 보낸다.
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
