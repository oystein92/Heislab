package main

import(
	
	"fmt"
	"time"
	"net"
	"strconv"

)
var number int=0;
var timeout = make(chan bool,1);
var quit = make(chan bool,1);

func master(connection *net.UDPConn){

	
		
	for{
		number=number+1;
		time.Sleep(500*time.Millisecond);
		connection.Write([]byte(strconv.Itoa(number)));		
		numbertemp, _ :=strconv.Atoi(string([]byte(strconv.Itoa(number))));
		fmt.Println(numbertemp);
	}
}


func slave(connection *net.UDPConn){
	msg := make([]byte,1024)	
	check := false
	for{
		connection.SetReadDeadline(time.Now().Add(1500*time.Millisecond));
		msgSize,_,err := connection.ReadFromUDP(msg);
		if err != nil {
			timeout <- true;
		
		}else{
			select{
		
			case <- quit:
				check = true;
			default:
				
		
			}
			if(check){
				break;
			}
			fmt.Println("Received:",string(msg[0:msgSize]))
			numbertemp, _ :=strconv.Atoi(string(msg[0:msgSize]));
			number=numbertemp;
	
		}
		
	
	}
	
}

func main(){
	workspace := 15;
	serverIP := "129.241.187.147";	
	serverPort := 20000 + 16;	
	serverAddr,_ := net.ResolveUDPAddr("udp",serverIP+ ":" +strconv.Itoa(serverPort));	
	readPort := 20000 + workspace;
	readAddr,_ := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(readPort));		
	readConn,_ := net.ListenUDP("udp",readAddr);
	sendConn,_ := net.DialUDP("udp",nil,serverAddr);
	go slave(readConn);
	switch {
	

        case <-timeout:
        	go master(sendConn);
        	quit <- true; 
        default:
                
        }	
	deadChan :=make(chan bool,1);
	<- deadChan;
}
