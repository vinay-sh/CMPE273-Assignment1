package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"strings"
	"math/rand"
	"strconv"

)

type StockParam struct{
	StockSymbolAndPercentage string
	Budget float32
}
type List struct {
stockName string
bid int
}
type CheckPort struct{
TradeId int
}
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

type Listener int
func (l *Listener) GetLine(line []byte, ack *bool) error {
	fmt.Println(string(line))
	return nil
}
var tradeId [10]int
var symbols [10]string
var querystring [10]string
var amtmain [10]float32
var budget [10]float32
var bid[10]float32
var countmain int=0
type Listenerx struct{}

func (t *Listenerx) CheckPortfolio(stocks *CheckPort, rec *Records) error {
var inxf int

findtid:=stocks.TradeId
for k:=0;k<len(tradeId);k++{
	fmt.Println(tradeId[k])

if tradeId[k]==findtid{
inxf=k
break;
}
}
fmt.Println("String"+symbols[inxf])
fmt.Println(amtmain[inxf])
var err error
req, err := http.NewRequest("GET", "https://query.yahooapis.com/v1/public/yql?q=select%20symbol%2CBid%20from%20yahoo.finance.quotes%20where%20symbol%20in%20("+querystring[inxf]+")%0A%09%09&format=json&diagnostics=true&env=http%3A%2F%2Fdatatables.org%2Falltables.env&callback=", nil)
if err != nil {
	log.Fatal(err)
}
req.SetBasicAuth("<token>", "x-oauth-basic")

client := http.Client{}
res, err := client.Do(req)
if err != nil {
	log.Fatal(err)
}

log.Println("StatusCode:", res.StatusCode)
var dat map[string]interface{}
// read body
body, err := ioutil.ReadAll(res.Body)
res.Body.Close()
if err != nil {
	log.Fatal(err)
}

err = json.Unmarshal(body, &dat);if err != nil {
panic(err)
}
v:=dat["query"].(map[string]interface{})["results"].(map[string]interface{})["quote"]
var w [5]string
var x [5]string
var vy [5]string
var cmp [5]string
var len1 [5]string

sym:=strings.Split(symbols[inxf],",")
var stockNames []string
var stockNames1 []string

for k:=0;k<len(sym);k++{
	stockNames=strings.Split(sym[k],"$")
	vy[k]=stockNames[0]
	stockNames1=strings.Split(vy[k],":")
	len1[k]=stockNames1[1]

	cmp[k]=stockNames[1]
}

var senddetails string
var maxsum float32
maxsum=0
for k:=0;k<len(v.([]interface{}));k++{
w[k]=v.([]interface{})[k].(map[string]interface{})["symbol"].(string)
x[k]=v.([]interface{})[k].(map[string]interface{})["Bid"].(string)

value1,err:=strconv.ParseFloat(cmp[k],32)
value2,err:=strconv.ParseFloat(x[k],32)
value3,err:=strconv.ParseFloat(len1[k],32)

maxsum=maxsum+float32(value1*value3)

if(value1>value2){
	x[k]="+"+cmp[k]
}
if(value1<value2){
	x[k]="-"+x[k]
}
t := strings.TrimSpace(x[k])
if err != nil {
	log.Fatal(err)
}
if(k==len(v.([]interface{}))-1){
senddetails=senddetails+vy[k]+"$"+t}else
{senddetails=senddetails+vy[k]+"$"+t+","}

}
fmt.Println(senddetails)
rec.Stocks=senddetails
rec.CurrentMarketValue=maxsum
rec.UninvestedAmount=amtmain[inxf]
fmt.Println(rec.Stocks)
fmt.Println(rec.CurrentMarketValue)
fmt.Println(rec.UninvestedAmount)
return nil
}
type Stock struct{}
func (t *Stock) BuyStock(stocks *StockParam, reply *Response) error {
symbol:=stocks.StockSymbolAndPercentage
	sym:=strings.Split(symbol,",")
	stockArray:=""
	var hasstockpercent [5]string
	var mystockarray [5]string
	y:=0
	for k:=0;k<len(sym);k++{
		stockNames:=strings.Split(sym[k],":")
if k==len(sym)-1{
stockArray=stockArray+"%22"+stockNames[0]+"%22"
}else{
stockArray=stockArray+"%22"+stockNames[0]+"%22%2C"
}
hasstockpercent[y]=stockNames[1]
mystockarray[y]=stockNames[0]
fmt.Println(hasstockpercent[k])
fmt.Println(mystockarray[k])

y++
}

	reply.TradeId = rand.Int()
tradeId[countmain]=reply.TradeId
querystring[countmain]=stockArray

	var err error
	req, err := http.NewRequest("GET", "https://query.yahooapis.com/v1/public/yql?q=select%20symbol%2CBid%20from%20yahoo.finance.quotes%20where%20symbol%20in%20("+stockArray+")%0A%09%09&format=json&diagnostics=true&env=http%3A%2F%2Fdatatables.org%2Falltables.env&callback=", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("<token>", "x-oauth-basic")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("StatusCode:", res.StatusCode)
	var dat map[string]interface{}
	// read body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &dat);if err != nil {
	panic(err)
}
v:=dat["query"].(map[string]interface{})["results"].(map[string]interface{})["quote"]
var w [5]string
var x [5]string
var finalResponse string=""
var val1 float32
var uinv float32
for k:=0;k<len(v.([]interface{}));k++{
	val1=stocks.Budget

varallval:=strings.Split(hasstockpercent[k],"%")
var pq float32
p, err := strconv.ParseFloat(varallval[0],32)
pq=float32(p/100)
pq=val1*pq;
fmt.Println(pq)
amt:=float32(pq);
fmt.Println(amt)
w[k]=v.([]interface{})[k].(map[string]interface{})["symbol"].(string)
x[k]=v.([]interface{})[k].(map[string]interface{})["Bid"].(string)
newx, err := strconv.ParseFloat(x[k],32)
fmt.Println(newx)
var newpqx float32
newpqx=float32(newx)
var noofshares float32
if newpqx<amt{
noofshares=amt/newpqx
}else{
noofshares=0
}
nnof:=int(noofshares)
strshares := strconv.Itoa(nnof)
//fmt.Println(strshares)
//fmt.Println("Symbol: "+w[k]+" Bid : "+x[k])
if k==len(v.([]interface{}))-1{
finalResponse=finalResponse+w[k]+":"+strshares+":$"+x[k]
}else{
finalResponse=finalResponse+w[k]+":"+strshares+":$"+x[k]+", "
}
if err != nil {
	log.Fatal(err)
}
fmt.Println(finalResponse)
uinv=uinv+(float32(nnof)*newpqx)
}
reply.Stocks= finalResponse
symbols[countmain]=reply.Stocks
budget[countmain]=stocks.Budget
reply.UninvestedAmount = stocks.Budget-uinv
amtmain[countmain]=reply.UninvestedAmount
countmain++
fmt.Println(reply.UninvestedAmount)
return nil
}


func main() {
	list := new(Listener)
	buystock := new(Stock)
	list2 := new(Listenerx)

	server := rpc.NewServer()
	server.Register(list)
	server.Register(buystock)
	server.Register(list2)

	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}

	}

}