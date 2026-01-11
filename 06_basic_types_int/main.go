package main

import "fmt"

func main(){

	killNormal := 99
	killGhost := 99
	killGhost++
	killNormal++
	totalKills := killGhost+killNormal
	avgKill:=totalKills/2

	fmt.Println(killGhost)
	fmt.Println(killNormal)
	fmt.Println(totalKills)
	fmt.Println(avgKill)

	ratio_normal:=49.0
	ratio_ghost:=50.0
	avgKillRatio:=(ratio_ghost+ratio_normal)/2
	fmt.Println(avgKillRatio)
}