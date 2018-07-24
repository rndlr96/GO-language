package main

import (
    "fmt"
)

func main() {
    c := make(chan int)

    go func() {
        for i := 0 ; i < 5 ; i++ {
            c <- i
        }
        close(c)
    }()

    for i := range c {
        fmt.Println("Channel value : ", i)
    }
}

// 이미 닫힌 채널에 값을 보내면 panic 발생
// 채널을 닫으면 range 루프 종료
// 채널이 얼려 있고, 값이 들어오지 않는다면 range는 계속 대기
// 채널을 가져온 뒤 두 번재 리턴값으로 채널이 닫혔는지 확인
