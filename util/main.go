package main

import (
	"fmt"
	"strings"
)

func main() {
	//hao := formatKuoHao("[[45,23],[404,3],[735,24],[856,3],[249,11],[326,7],[589,8],[91,17],[126,24],[713,21],[606,13],[585,5],[861,30],[604,4],[822,27],[231,6],[507,6],[265,5],[912,12],[878,29],[46,6],[421,20],[941,26],[151,25],[490,4],[315,14],[630,16],[292,24],[214,2],[432,18],[520,21],[88,26],[4,21],[337,28],[780,9],[220,29],[721,13],[927,25],[67,25],[835,21],[646,19],[973,26],[235,12],[427,9],[471,21],[267,16],[388,8],[788,13],[937,27],[810,26],[288,5],[966,28],[698,15],[343,12],[648,20],[238,4],[436,15],[588,4],[373,15],[100,12],[180,19],[904,15],[854,21],[107,2],[822,12],[485,24],[196,10],[978,24],[178,10],[642,29],[455,16],[490,17],[455,25],[77,7],[456,15],[102,25],[767,19],[169,5],[461,11],[385,25],[896,5],[185,26],[885,14],[948,26],[907,6],[877,18],[421,24],[783,15],[999,20],[756,10],[308,12],[34,12],[339,17],[613,15],[270,10],[681,3],[385,15],[123,4],[10,20],[799,28],[506,23],[265,24],[193,4],[638,23],[144,29],[874,19],[470,14],[195,16],[77,23],[573,28],[559,12],[146,10],[538,10],[705,4],[592,26],[258,24],[900,25],[836,8],[353,5],[197,26],[572,21],[347,17],[763,21],[67,20],[927,6],[135,4],[392,30],[131,15],[650,23],[100,30],[848,9],[858,27],[203,15],[249,4],[884,15],[465,18],[316,30],[730,15],[310,19],[823,21],[785,21],[15,16]]\n772\n502")
	//fmt.Println(hao)

	kuoHao := formatKuoHao("6[[1,4],[3,4],[5,4],[2,4],[3,1],[5,3],[2,1],[0,5],[0,2],[5,1],[3,2],[0,3]]")
	fmt.Println(kuoHao)

}

func formatKuoHao(s string) string {
	s = strings.ReplaceAll(s, "[", "{")
	s = strings.ReplaceAll(s, "]", "}")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}
