package main  

import (
	"fmt"        //표준 패키지
	"golang/study/4week/4week_2_mod"    // 모듈 내 패키지

	"github.com/guptarohit/asciigraph"   //깃허브 외부 저장소 패키지
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main () {
	4week_2_mod.PrinCustom()
	expkg.PrintSample()

	data := []float64{3, 6, 9, 1, 2, 3}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
	g
}