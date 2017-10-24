package helpers

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	CARRIAGE_RETURN           = 0x0D
	LINE_FEED                 = 0x0A
	SPACE                     = 0x20
	DELAY_BETWEEN_CONNECTIONS = time.Second * 10
)

func SwitchInput(address, input, output string) error {
	command := ("x" + input + "AVx" + output)
	conn := getConnection(address)
	command += string(CARRIAGE_RETURN)
	conn.Write([]byte(command))
	resp, err := readUntil('\n', conn, 1)
	if err != nil {
		log.Printf("Maybe it didn't read 1")
		log.Printf("%s", err)
	}
	log.Printf("Feedback: %s", resp)
	conn.Close()
	return nil
}

func getConnection(address string) *net.TCPConn {
	addr, err := net.ResolveTCPAddr("tcp", address+":23")
	if err != nil {
		return nil
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil
	}

	_, err = readUntil('\n', conn, 1)
	if err != nil {
		log.Printf("Maybe it didn't read 2")
		log.Printf("%s", err)
	}
	return conn
}

func readUntil(delimeter byte, conn *net.TCPConn, timeoutInSeconds int) ([]byte, error) {
	conn.SetReadDeadline(time.Now().Add(time.Duration(int64(timeoutInSeconds)) * time.Second))

	buffer := make([]byte, 128)
	message := []byte{}
	found := false
	index := -1

	for !found {
		_, err := conn.Read(buffer)
		if err != nil {
			err = errors.New(fmt.Sprintf("Error reading response: %s", err.Error()))
			log.Printf("%s", err.Error())
			return message, err
		}
		found, index = charInBuffer(delimeter, buffer)
		message = append(message, buffer[:index]...)
	}
	return message, nil
}

func charInBuffer(toCheck byte, buffer []byte) (bool, int) {
	for i, b := range buffer {
		if toCheck == b {
			return true, i
		}
	}
	return false, len(buffer)
}

// func ChangePower(address, command string) string {
//   conn := getConnection(address)
//   command += string(CARRIAGE_RETURN)
//   conn.Write([]byte(command))
//   resp, err := readUntil(CARRIAGE_RETURN, conn, 1)
//   if err != nil {
//      log.Printf("Maybe it didn't read")
//   }
//   s := string(resp)
//   conn.Close()
//   return s
// }

// func removeNil(b []byte) (ret []byte) {
// 	for _, c := range b {
// 		switch c {
// 		case '\x00':
// 			break
// 		default:
// 			ret = append(ret, c)
// 		}
// 	}
// 	return ret
// }
