package day03

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_practise_01(t *testing.T) {
	/*
	   剪刀石头布：
	       系统产生1-3的随机数，分别代表剪刀，石头和布
	       玩家键盘输入1-3数字，分别代表剪刀，石头和布

	*/
	choose := ""
	sysNum := 0 //系统产生的随机数
	num := 0    //玩家键盘输入
	name1 := "" //系统的：剪刀，石头，布
	name2 := "" // 玩家的：剪刀，石头，布
	for {
		fmt.Println("请输入z代表退出，任意键代表继续游戏：")
		fmt.Scanln(&choose)
		if choose == "z" {
			fmt.Println("欢迎下次再来。。")
			break
		} else {
			rand.Seed(time.Now().UnixNano())
			sysNum = rand.Intn(3) + 1
			switch sysNum {
			case 1:
				name1 = "剪刀"
			case 2:
				name1 = "石头"
			case 3:
				name1 = "布"
			}
			//玩家输入
			fmt.Println("请输入：1代表剪刀、2代表石头、3代表布：")
			fmt.Scanln(&num)
			switch num {
			case 1:
				name2 = "剪刀"
			case 2:
				name2 = "石头"
			case 3:
				name2 = "布"
			}
			//判定输赢
			if num == sysNum {
				fmt.Printf("系统出：%s，玩家出：%s,平局\n", name1, name2)
			} else {
				switch sysNum {
				case 1: //系统剪刀
					if num == 2 { //玩家石头
						fmt.Printf("系统出：%s，玩家出：%s,玩家胜\n", name1, name2)
					} else if num == 3 { //玩家布
						fmt.Printf("系统出：%s，玩家出：%s,玩家输\n", name1, name2)
					}

				case 2: //系统石头
					if num == 1 {
						//玩家剪刀
						fmt.Printf("系统出：%s，玩家出：%s,玩家输\n", name1, name2)
					} else if num == 3 {
						//玩家布
						fmt.Printf("系统出：%s，玩家出：%s,玩家胜\n", name1, name2)
					}
				case 3: //系统布
					if num == 1 {
						fmt.Printf("系统出：%s，玩家出：%s,玩家胜\n", name1, name2)
					} else if num == 2 {
						fmt.Printf("系统出：%s，玩家出：%s,玩家输\n", name1, name2)
					}
				}
			}

		}
	}

}
