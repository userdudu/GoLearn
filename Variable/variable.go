package main

import(
	"fmt"
)

type UserInfo struct{
	strName string
	nAge	int
}

var g_id int

func setid( nid int ){
	g_id = nid
}

func gcd( x, y int )int{
	for y != 0 {
		x, y = y, x % y
	}
	return x
}

func main(){
	var nNumb int = 100
	fmt.Println( nNumb )

	var user UserInfo = UserInfo{"shane",31}
	fmt.Println( user )

	setid( 200 )
	fmt.Println( g_id )

	nNumb = gcd( 13,10 )
	fmt.Println( nNumb )
}
