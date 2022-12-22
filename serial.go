package main

import (
	"fmt"
	"log"
	"strconv"

	"go.bug.st/serial"
)

var port serial.Port

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

func getSerialInput() {
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
			//case 48: // 0
			default:
				if started {
					b = append(b, buff[i])
				}
			}

		}
		if b != nil {
			serialOutput, err := strconv.Atoi(fmt.Sprintf("%v", string(b)))
			if err == nil && !started {
				//return serialOutput
				serialInput <- serialOutput
			} else if err != nil {
				log.Fatal(err)
				return
			} else {
				continue
			}
		}
	}
}
