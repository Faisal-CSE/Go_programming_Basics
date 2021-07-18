package main

import (
	"fmt"
	"log"
	"net"
	"project1/StringUtility"
)

// GetOutboundIP : Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}


func main()  {
	fmt.Println("Hello World!")

	s := StringUtility.Upper("faisal porag")

	x := 15
	y := 20

	result := StringUtility.Summation(x , y)

	fmt.Println(s)

	fmt.Println("Result is ", result)

	var a []int = []int{7,5,3,10,74,1,62,6,4}

	for i, element := range a{
		for j := i+1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element{
				fmt.Println(element)
			}//else{
				//fmt.Println("No duplicates ...")
			//}
		}
	}

	fmt.Println("IP Address: ", GetOutboundIP())

}
