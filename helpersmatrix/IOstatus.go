package helpersmatrix

import (
	"log"
)

func GetInput(address, output string) (string, error) {
	command := ("Statusx" + output)
	conn := getConnection(address)
	command += string(CARRIAGE_RETURN)
	conn.Write([]byte(command))
	resp, err := readUntil(CARRIAGE_RETURN, conn, 1)
	if err != nil {
		log.Printf("%s", err)
	}
	log.Printf("Feedback: %s", resp)
	s := string(resp[1])
	conn.Close()
	return s, err
}
