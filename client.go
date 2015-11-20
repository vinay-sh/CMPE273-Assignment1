package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)
type Response struct{
TradeId int
Stocks string
UninvestedAmount float32
}
type Records struct{
Stocks string
CurrentMarketValue float32
UninvestedAmount float32
}
type CheckPort struct{
TradeId int
}
type StockParam struct{
	StockSymbolAndPercentage string
	Budget float32
}
func main() {
fmt.Println("Welcome User!!!")
fmt.Println("Enter Request String : ")
var input1 string
fmt.Scanln(&input1)
fmt.Println("Enter Budget : ")
var input2 float32
fmt.Scanln(&input2)

	client, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply Response
	var response Records
	sp:=&StockParam{input1,input2}
	c := jsonrpc.NewClient(client)
	err = c.Call("Stock.BuyStock", sp, &reply)
	fmt.Printf("Result: %d %s %f \n",reply.TradeId,reply.Stocks,reply.UninvestedAmount)
	fmt.Println("Enter TradeID : ")
	var input3 int
	fmt.Scanln(&input3)
	cp:=&CheckPort{input3}
	err = c.Call("Listenerx.CheckPortfolio", cp,&response)
	fmt.Printf("Result: %s  Current Market Value :%f Univested Amt: %f \n",response.Stocks,response.CurrentMarketValue,response.UninvestedAmount)
	if err != nil {
		log.Fatal("error:", err)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply1 bool
		err = c.Call("Listener.GetLine", line, &reply1)
		if err != nil {
			log.Fatal(err)
		}
	}
}