package serialcon

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"go.bug.st/serial"
)

var port serial.Port
var PrevSer int = 100
var SerialInput chan int = make(chan int) // Channel to read the serial input (seperate thread)

func InitSerial() {
	var err error
	mode := &serial.Mode{
		BaudRate: 9600,
	}

	port, err = serial.Open(getPorts(), mode)
	if err != nil {
		log.Fatal(err)
	}
	go GetSerialInput() // start up the serial channel

	time.Sleep(time.Second * 2)
}

// Returns a list of all available ports found (to find the port that the Arduino is connected to)
func getPorts() string {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
		return ""
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
		//fmt.Println(port.Read())
	}
	return ports[0]
}

// Reads the serial input send by the Arduino.
func GetSerialInput() {
	var buff []byte = make([]byte, 10)
	var b []byte
	started := false
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			return
		}
		if n == 0 {
			fmt.Println("\nEOF")
			return
		}
		//fmt.Printf("buff %v\n", buff[:n])
		for i := 0; i < n; i++ {

			switch buff[i] {
			case 10: // LF  acts as final signal
			case 13: // CR - ignored
				started = false
			case 88: // X acts as start signal
				started = true
				b = nil
			default:
				if started {
					b = append(b, buff[i])
				}
			}

		}

		// If serial data came in, it is loaded into the SerialInput queue
		if b != nil {
			serialOutput, err := strconv.Atoi(fmt.Sprintf("%v", string(b))) //converts ascii value to int
			if err == nil && !started {
				SerialInput <- serialOutput
			} else if err != nil {
				log.Fatal(err)
				return
			} else {
				continue
			}
		}
	}
}
