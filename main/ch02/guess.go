package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	randomValue := rand.Intn(100) + 1

	// 让用户猜
	count := 10
	log.Println("请猜一个数字，范围是1-100")
	for {
		log.Println("剩余次数", count)
		count--
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		parseInt, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}
		if parseInt < randomValue {
			log.Println("哎呀，你猜低了")
		} else if parseInt > randomValue {
			log.Println("哎呀，你猜高了")
		} else {
			log.Println("恭喜你猜对了")
			break
		}

		if count == 0 {
			log.Println("你已经猜了", count, "次，游戏结束，实际值是", randomValue)
			break
		}
	}

}
