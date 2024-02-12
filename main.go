package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//参数设置
	name := flag.String("name", "尼格买提热合曼", "你的姓名")
	region := flag.Int("region", 2, "南方人=1, 北方人=2, 不确定=3")
	gender := flag.Int("gender", 1, "男生=1, 女生=2")
	flag.Parse()
	magic(*name, *region, *gender)
}

// 从牌面抽一张移到牌底
func moveFirstToLast(cards []int) []int {
	first := cards[0]
	cards = cards[1:]
	cards = append(cards, first)
	return cards
}

// 把前n张牌插进剩下卡片的中间去, 注意如果不是插中间,而是放到最后一张牌后面就会不成功
// 我估计尼格买提就是这里错了,但是电视镜头给了刘谦特写,看不到尼格买提的操作
func moveTopN(cards []int, n int) []int {
	randInt := randIntRange(n+1, 7) //这里7改成8就可能插到最后一张牌后面从而引发错误
	firstThree, beforeInsert, afterInsert := cards[:n], cards[n:randInt], cards[randInt:]
	// fmt.Println("firstThree=", firstThree)
	// fmt.Println("beforeInsert=", beforeInsert)
	// fmt.Println("afterInsert=", afterInsert)

	// 组合这三个部分形成新的切片
	newCards := make([]int, 0)
	newCards = append(newCards, beforeInsert...)
	newCards = append(newCards, firstThree...)
	newCards = append(newCards, afterInsert...)

	return newCards
}

// 丢弃的牌
func discardTopN(cards []int, n int) []int {
	cards = cards[n:]
	return cards
}

// 生成[min, max)范围内的随机整数
func randIntRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	return min + rand.Intn(max-min)
}

func magic(name string, region, gender int) bool {
	// 第一步,4张牌,对半撕开(共8张)
	cards := []int{1, 2, 3, 4, 1, 2, 3, 4}
	fmt.Println("原始牌", cards)

	// 第二步,名字几个字放几张牌下去
	for i := 0; i < len([]rune(name)); i++ {
		cards = moveFirstToLast(cards)
	}
	fmt.Println("名字几个字放几张牌下去:", cards)

	// 第三步,前3张插进剩下的牌的任何一个位置
	cards = moveTopN(cards, 3)
	fmt.Println("前3张插进剩下的牌的任何一个位置:", cards)

	// 第四步,第一张牌藏屁股底下,供后面对比用
	buttCard := cards[0]
	fmt.Println("屁股下的牌:", buttCard)

	cards = discardTopN(cards, 1)
	fmt.Println("剩下的牌:", cards)

	// 第五步,南方人前1张,北方人前2张,不确定前3张,插进剩下的牌的任何一个位置(重复第3步)
	cards = moveTopN(cards, region)
	fmt.Println("南方人1张,北方人2张,不确定3张,插进剩下的牌的任何一个位置:", cards)

	// 第六步,男生丢1张,女生丢2张
	cards = discardTopN(cards, gender)
	fmt.Println("男生丢1张,女生丢2张:", cards)

	// 第七步,见证奇迹的时刻七个字每念一个字放一张下去
	spell := "见证奇迹的时刻"
	for _, c := range spell {
		cards = moveFirstToLast(cards)
		fmt.Println(string(c), cards)
	}

	// 第八步,“好运留下来,烦恼丢出去”
	// 第一张去最下面,第二张丢弃,以此类推,直到最后男生剩2张,女生剩1张(因为女生第六步多丢一张)
	for len(cards) >= 2 {
		cards = moveFirstToLast(cards)
		fmt.Println("好运留下来", cards)

		cards = discardTopN(cards, 1)
		fmt.Println("烦恼丢出去", cards)
	}

	//见证奇迹
	fmt.Println("最后手里的牌:", cards[0])
	return cards[0] == buttCard
}
