package main

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

const (
	Rank     = iota
	RF       //皇家同花顺bai(royal flush)
	SF       //同花顺(straight flush)
	FOAK     //4条(four of a kind)
	FH       //满堂红(葫芦)(full house)
	FLUSH    //同花(flush)
	STRAIGHT //顺子
	TOAK     //3条(three of a kind)
	TP       //两对(two pairs)
	OP       //一对(one pair)
	HC       //高牌(high card)

)

func main() {
	fmt.Println()
	fmt.Println("——————德州扑克比大小—————")
	fmt.Println("请按照命名规则给alice发牌（5张）")
	var cardOfAlice string //Alice的手牌，以字符串存储
	var cardOfBob string   //Bob的手牌，以字符串存储
	fmt.Scanln(&cardOfAlice)
	fmt.Println("alice的手牌为:", cardOfAlice)
	var rankOfAlice, acardSet, cardNumAlice = GetRank(cardOfAlice)
	fmt.Println("Alice的手牌等级排名", rankOfAlice)
	fmt.Println("Alice的", acardSet)
	fmt.Println("请按照命名规则给bob发牌（5张）")
	fmt.Scanln(&cardOfBob)
	fmt.Println("bob的手牌为", cardOfBob)
	var rankOfBob, bcardSet, cardNumBob = GetRank(cardOfBob)
	fmt.Println("Bob的手牌等级排名", rankOfBob)
	fmt.Println("Bob的", bcardSet)
	//卡牌比较
	winner := CompareCard(rankOfAlice, rankOfBob, cardNumAlice, cardNumBob)
	fmt.Println("胜者为：", winner)
}

func GetRank(str string) (int, string, []int) {
	rank := 11                  //默认最低牌级为11
	var numList list.List       //存放手牌数字
	var colorList list.List     //存放手牌花色
	var numMap map[string]int   //存放数字种类数量，以只用key,将map作为set使用
	var colorMap map[string]int //存放花色种类数量，以只用key,将map作为set使用
	numMap = make(map[string]int)
	colorMap = make(map[string]int)

	for i := 0; i <= len(str)-2; i += 2 {
		num := str[i : i+1]
		color := str[i+1 : i+2]
		if num == "T" {
			numList.PushBack(10)
			numMap["10"] = 0
		} else if num == "J" {
			numList.PushBack(11)
			numMap["J"] = 0
		} else if num == "Q" {
			numList.PushBack(12)
			numMap["Q"] = 0
		} else if num == "K" {
			numList.PushBack(13)
			numMap["K"] = 0
		} else if num == "A" {
			numList.PushBack(14)
			numMap["A"] = 0
		} else {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			numList.PushBack(i)
			numMap[num] = 0
		}
		colorList.PushBack(color)
		colorMap[color] = 0
	}
	size := numList.Len()
	numArray := GetSorttedArray(numList, size)
	//判断1.皇家同花顺-或者-同花顺-或者-同花不带顺
	if len(colorMap) == 1 {
		fmt.Println("是同花的")
		//len:=numList.Len()
		//numArray:=GetSorttedArray(numList,len)
		fmt.Println("after string :numArray", numArray)
		//皇家同花顺判断
		if numArray[0] == 10 && numArray[1] == 11 && numArray[2] == 12 && numArray[3] == 13 && numArray[4] == 14 {
			rank = RF //rank=RF=1 皇家同花顺
		}

		for i, flag := 0, 0; i <= size-2; i++ {
			if numArray[i+1]-numArray[i] == 1 {
				flag++
				if flag == 4 && rank > 2 {
					rank = SF //rank=SF=2 同花顺
				}
				continue
			} else if rank > 5 {
				rank = FLUSH //rank=FLUSH=5 同花不带顺
			}
		}
	}

	//判断2.四条-或者-葫芦
	if len(numMap) == 2 {
		fmt.Println("只有两种点数")
		//len:=numList.Len()
		//numArray:=GetSorttedArray(numList,len)
		fmt.Println("after string :numArray", numArray)
		if !(numArray[0] == numArray[1]) || !(numArray[size-1] == numArray[size-2]) {
			if rank > 3 {
				rank = FOAK //ran=FOAK=3 四条
			}
		} else if rank > 4 {
			rank = FH //ran=FH=4 葫芦
		}
		switch rank {
		case 3:
			fmt.Println("牌等级是四条")
		case 4:
			fmt.Println("牌等级是葫芦")
		}
	}

	//判断3.顺子
	if len(colorMap) > 1 {
		fmt.Println("不是同花的")
		//len:=numList.Len()
		//numArray:=GetSorttedArray(numList,len)
		fmt.Println("after string :numArray", numArray)
		for i, flag := 0, 0; i <= size-2; i++ {
			if numArray[i+1]-numArray[i] == 1 {
				flag++
				if flag == 4 && rank > 6 {
					rank = STRAIGHT //rank=STRAIGHT=6 顺子
				}
			} else {
				break
			}
		}
	}

	//判断4.三条-或者-两对
	if len(numMap) == 3 {
		fmt.Println("数字仅三种")
		//len:=numList.Len()
		//numArray:=GetSorttedArray(numList,len)
		for i := 0; i < size-2; i++ {
			if numArray[i] == numArray[i+1] {
				if numArray[i] == numArray[i+2] && rank > 7 {
					rank = TOAK //rank=TOAK=7 三条
					continue
				}
			} else if rank > 8 {
				rank = TP //rank=TP=8 两对
			}
		}
	}
	//判断5 一对
	if len(numMap) == 4 {
		if rank > 9 {
			rank = OP //rank=OP=9 一对
		}
	}
	//判断6 高牌
	if len(numMap) == 5 {
		if rank > 10 {
			rank = HC //rank=HC=10 高牌
		}
	}

	var nameOfRank string
	switch rank {
	case 1:
		nameOfRank = "牌级为皇家同花顺"
	case 2:
		nameOfRank = "牌级为同花顺"
	case 3:
		nameOfRank = "牌级为四条"
	case 4:
		nameOfRank = "牌级为葫芦"
	case 5:
		nameOfRank = "牌级为同花"
	case 6:
		nameOfRank = "牌级为顺子"
	case 7:
		nameOfRank = "牌级为顺子"
	case 8:
		nameOfRank = "牌级为两对"
	case 9:
		nameOfRank = "牌级为一对"
	case 10:
		nameOfRank = "牌级为高牌"
	}
	//fmt.Println()
	//fmt.Println(size)
	//fmt.Println("numLIst:")
	//for i:=numList.Front();i!=nil ;i=i.Next(){
	//	fmt.Print(i.Value,"  ")
	//}
	//fmt.Println()
	//fmt.Println("colorList:")
	//for i := colorList.Front(); i != nil; i = i.Next() {
	//	fmt.Print(i.Value,"  ")
	//}
	//fmt.Println()
	//fmt.Println("numMap",numMap)
	//fmt.Println("colorMap",colorMap)
	//fmt.Println("colorMap的长度",len(colorMap))
	return rank, nameOfRank, numArray
}

func CompareCard(rankA int, rankB int, cardNumAlice []int, cardNumBob []int) string {
	//比大小，如果牌级更高则直接获胜，牌级相等则进行判断
	var winner string
	if rankA < rankB {
		winner = "Alice"
	} else if rankA > rankB {
		winner = "Bob"
	}
	//ranA==ranB，则进行条件比较
	if rankA == rankB {
		switch rankA {
		case 1:
			winner = "双方平手"
		case 2:
			if cardNumBob[len(cardNumBob)-1] > cardNumAlice[len(cardNumAlice)-1] {
				winner = "Bob"
			} else if cardNumBob[len(cardNumBob)-1] < cardNumAlice[len(cardNumAlice)-1] {
				winner = "Alice"
			} else if cardNumBob[len(cardNumBob)-1] == cardNumAlice[len(cardNumAlice)-1] {
				winner = "双方平手"
			}
		case 3:
			if cardNumBob[len(cardNumBob)/2] > cardNumAlice[len(cardNumAlice)/2] {
				winner = "Bob"
			} else if cardNumBob[len(cardNumBob)/2] < cardNumAlice[len(cardNumAlice)/2] {
				winner = "Alice"
			} else if cardNumBob[len(cardNumBob)/2] == cardNumAlice[len(cardNumAlice)/2] {
				a := cardNumBob[len(cardNumBob)/2]
				for i := 0; i < len(cardNumAlice)-1; i++ { //可以改为不遍历，只比较首尾两元素
					if cardNumAlice[i] != a {
						a = cardNumAlice[i]
						break
					}
				}
				b := cardNumBob[len(cardNumBob)/2]
				for i := 0; i < len(cardNumBob)-1; i++ {
					if cardNumBob[i] != b {
						b = cardNumBob[i]
						break
					}
				}
				if a > b {
					winner = "Alice"
				} else if a < b {
					winner = "Bob"
				} else if a == b {
					winner = "双方平手"
				}
			}
		case 4:
			middle := len(cardNumAlice) / 2
			//5牌时中位数一定为三条的重复数字
			aNumOfThree := cardNumAlice[middle]
			bNumOfThree := cardNumBob[middle]
			if aNumOfThree > bNumOfThree {
				winner = "Alice"
			} else if aNumOfThree < bNumOfThree {
				winner = "Bob"
			} else if aNumOfThree == bNumOfThree {
				//a、b存放数组三代二中一对数字的数值
				var a int
				var b int
				if aNumOfThree > cardNumAlice[middle-1] {
					//a不变
					a = cardNumAlice[middle-1]
				} else if aNumOfThree == cardNumAlice[middle-1] {
					//当中位数与上一位相等，则一对存在之后两位
					a = cardNumAlice[middle+1]
				}
				if bNumOfThree > cardNumBob[middle-1] {
					//a不变
					b = cardNumAlice[middle-1]
				} else if bNumOfThree == cardNumBob[middle-1] {
					//当中位数与上一位相等，则一对存在之后两位
					b = cardNumAlice[middle+1]
				}
				if a > b {
					winner = "Alice"
				} else if a < b {
					winner = "Bob"
				} else if a == b {
					winner = "平手"
				}
			}
		case 5:
			asum := 0
			bsum := 0
			for i := 0; i < len(cardNumAlice)-1; i++ {
				asum += cardNumAlice[i]
			}
			for i := 0; i < len(cardNumBob)-1; i++ {
				bsum += cardNumBob[i]
			}
			if asum > bsum {
				winner = "Alice"
			} else if asum < bsum {
				winner = "Bob"
			} else if asum == bsum {
				for i := 0; i < len(cardNumBob)-1; i++ {
					if cardNumAlice[len(cardNumAlice)-1-i] == cardNumBob[len(cardNumBob)-1-i] {
						winner = "平手"
						continue
					}
					if cardNumAlice[len(cardNumAlice)-1-i] > cardNumBob[len(cardNumBob)-1-i] {
						winner = "ALice"
						break
					} else if cardNumAlice[len(cardNumAlice)-1-i] < cardNumBob[len(cardNumBob)-1-i] {
						winner = "Bob"
					}
				}
			}
		case 6:
			if cardNumAlice[0] > cardNumBob[0] {
				winner = "Alice"
			} else if cardNumAlice[0] < cardNumBob[0] {
				winner = "Bob"
			} else if cardNumAlice[0] == cardNumBob[0] {
				winner = "平手"
			}
		case 7:
			middle := len(cardNumAlice) / 2
			//5牌时中位数一定为三条的重复数字
			aNumOfThree := cardNumAlice[middle]
			bNumOfThree := cardNumBob[middle]
			if aNumOfThree > bNumOfThree {
				winner = "Alice"
			} else if aNumOfThree < bNumOfThree {
				winner = "Bob"
			} else if aNumOfThree == bNumOfThree {
				var twoOfA = [2]int{0, 0}
				var twoOfB = [2]int{0, 0}
				for i, j := 0, 0; i < len(cardNumAlice)-1; i++ {
					if cardNumAlice[i] != cardNumAlice[middle] && cardNumAlice[i] != twoOfA[j] {
						twoOfA[j] = cardNumAlice[i]
						j++
					}
				}
				for i, j := 0, 0; i < len(cardNumBob)-1; i++ {
					if cardNumBob[i] != cardNumBob[middle] && cardNumBob[i] != twoOfB[j] {
						twoOfB[j] = cardNumAlice[i]
						j++
					}
				}
				sort.Ints(twoOfA[:])
				sort.Ints(twoOfB[:])
				if twoOfA[1] > twoOfB[1] {
					winner = "ALice"
				} else if twoOfA[1] < twoOfB[1] {
					winner = "Bob"
				} else if twoOfA[1] == twoOfB[1] {
					if twoOfA[0] > twoOfB[0] {
						winner = "Alice"
					} else if twoOfA[0] < twoOfB[0] {
						winner = "Bob"
					} else if twoOfA[0] == twoOfB[0] {
						winner = "平手"
					}
				}
			}
		case 8:
			//五牌的两对中，起脚牌只可能存在于0、2、4，则两对分别一定存在于1、3位
			var twoPairOfA = [2]int{cardNumAlice[1], cardNumAlice[3]}
			var twoPairOfB = [2]int{cardNumBob[1], cardNumAlice[3]}
			if twoPairOfA[1] > twoPairOfB[1] {
				winner = "ALice"
			} else if twoPairOfA[1] < twoPairOfB[1] {
				winner = "Bob"
			} else if twoPairOfA[1] == twoPairOfB[1] {
				if twoPairOfA[0] > twoPairOfA[0] {
					winner = "ALice"
				} else if twoPairOfA[0] < twoPairOfB[0] {
					winner = "Bob"
				} else if twoPairOfA[0] == twoPairOfB[0] {
					a := 0
					b := 0
					for i := 0; i < len(cardNumAlice)-1; i++ {
						if cardNumAlice[i] != twoPairOfA[0] && cardNumAlice[i] != twoPairOfA[1] {
							a = cardNumAlice[i]
							break
						}
					}
					for i := 0; i < len(cardNumAlice)-1; i++ {
						if cardNumBob[i] != twoPairOfB[0] && cardNumBob[i] != twoPairOfB[1] {
							b = cardNumAlice[i]
							break
						}
					}
					if a > b {
						winner = "ALice"
					} else if a < b {
						winner = "Bob"
					} else if a == b {
						winner = "平手"
					}
				}
			}
		case 9:
			pairOfA := 0
			pairOfB := 0
			for i := 1; i < len(cardNumAlice)-1; i++ {
				if cardNumAlice[i] == cardNumAlice[i-1] {
					pairOfA = cardNumAlice[i]
					break
				}
			}
			for i := 1; i < len(cardNumBob)-1; i++ {
				if cardNumBob[i] == cardNumBob[i-1] {
					pairOfB = cardNumBob[i]
					break
				}
			}
			if pairOfA > pairOfB {
				winner = "Alice"
			} else if pairOfA < pairOfB {
				winner = "Bob"
			} else if pairOfA == pairOfB {
				var lastThreeA = [3]int{0, 0, 0}
				var lastThreeB = [3]int{0, 0, 0}
				for i, j := 0, 0; i < len(cardNumAlice)-1; i++ {
					if cardNumAlice[i] != pairOfA {
						lastThreeA[j] = cardNumAlice[i]
						j++
					}
				}
				for i, j := 0, 0; i < len(cardNumBob)-1; i++ {
					if cardNumBob[i] != pairOfB {
						lastThreeB[j] = cardNumBob[i]
						j++
					}
				}
				for i := 0; i < len(lastThreeA)-1; i++ {
					if lastThreeA[len(lastThreeA)-1-i] == lastThreeB[len(lastThreeA)-1-i] {
						winner = "平手"
						continue
					} else if lastThreeA[len(lastThreeA)-1-i] > lastThreeB[len(lastThreeB)-1-i] {
						winner = "ALice"
					} else if lastThreeA[len(lastThreeA)-1-i] < lastThreeB[len(lastThreeA)-1-i] {
						winner = "Bob"
					}
				}
			}
		case 10:
			size := len(cardNumAlice)
			for i := 0; i < size-1; i++ {
				if cardNumAlice[size-1-i] == cardNumBob[size-1-i] {
					winner = "平手"
					continue
				} else if cardNumAlice[size-1-i] > cardNumBob[size-1-i] {
					winner = "Alice"
				} else if cardNumAlice[size-1-i] < cardNumBob[size-1-i] {
					winner = "Bob"
				}
			}
		}
	}
	return winner
}

func GetSorttedArray(list list.List, len int) (numArray []int) {
	numArray = make([]int, len)
	index := 0
	for i := list.Front(); i != nil; i = i.Next() {
		numArray[index] = i.Value.(int)
		index++
	}
	sort.Ints(numArray[:])
	//fmt.Println("after string :numArray",numArray)
	return numArray
}
