package filebeatUdpWriter

import (
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func TestWriter_Write(t *testing.T) {
	r := gin.New()

	//make a chan for udp server
	chan1 := make(chan string)
	quitChan := make(chan bool)
	//start a dummy udp server
	go startUdpServer(chan1, quitChan)

	udpLogger, _ := CreateLogger("localhost:18125") //you can catch the error here
	//asset the logger is not nil
	//sleep 1 sec
	time.Sleep(1 * time.Second)

	r.Use(GinHandle("gin", udpLogger))
	//r.Use(GinHandle("gin", log.Logger))
	testMsg := "test udp logger"
	udpLogger.Info().Msg(testMsg)

	//send quit signal to udp server
	quitChan <- true

	time.Sleep(1 * time.Second)
	//read from chan
	msg := <-chan1
	//check if msg contains testMsg
	if strings.Contains(msg, testMsg) == false {
		fmt.Println(msg)
		t.Errorf("expected %s, got %s", testMsg, msg)
	}

}

func startUdpServer(
	//output chan
	chan1 chan string,
	quitChan chan bool,
) {
	//start a dummy udp server

	//hostName := "localhost"
	portNum := "18125"

	//listen to incoming udp packets
	pc, err := net.ListenPacket("udp", "localhost:"+portNum)
	if err != nil {
		log.Fatal().Err(err).Msg("error listening to udp packets")
	}

	log.Info().Msg("listening on " + pc.LocalAddr().String())
	defer pc.Close()

	//loop
	for {
		//check if quitChan has a value
		select {
		case <-quitChan:

			//read incoming udp packets
			buffer := make([]byte, 1024)
			n, _, err := pc.ReadFrom(buffer)
			if err != nil {
				log.Fatal().Err(err).Msg("error reading from udp packets")
			}

			//log.Info().Str("address", addr.String()).Msg("received: " + string(buffer[:n]))
			//push to chan1el
			chan1 <- string(buffer[:n])

			log.Info().Msg("quitting udp server")
			return
		default:
			// avoid spinning
			time.Sleep(100 * time.Millisecond)
		}

	}

}
