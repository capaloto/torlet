package tor

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

const (
	cmdAuthenticate = "AUTHENTICATE"
	cmdNewCircuit   = "SIGNAL NEWNYM"
	cmdQuit         = "quit"
	cmdOK           = 250
)

// shamelessly stolen from postfix/goControlTor
type controller struct {
	controlConn     net.Conn
	textprotoReader *textproto.Reader
	addr            string
	pass            string
}

func newController(controlAddr, controlPass string) *controller {
	return &controller{addr: controlAddr, pass: controlPass}
}

// Dial handles unix domain sockets and tcp!
func (t *controller) dial() error {
	var err error = nil
	t.controlConn, err = net.Dial("tcp", t.addr)
	if err != nil {
		return fmt.Errorf("dialing tor control port: %s", err)
	}
	reader := bufio.NewReader(t.controlConn)
	t.textprotoReader = textproto.NewReader(reader)
	return nil
}

// Send a command to the controller
func (t *controller) sendCommand(command string) (int, string, error) {
	var code int
	var message string
	var err error

	_, err = t.controlConn.Write([]byte(command))
	if err != nil {
		return 0, "", fmt.Errorf("writing to tor control port: %s", err)
	}
	code, message, err = t.textprotoReader.ReadResponse(cmdOK)
	if err != nil {
		return code, message, fmt.Errorf("reading tor control port command status: %s", err)
	}
	return code, message, nil
}

// reset the tor circuit
func (t *controller) resetCircuit() error {
	resetCommand := fmt.Sprintf("%s \"%s\"\n", cmdAuthenticate, t.pass)
	resetCommand += cmdNewCircuit + "\n"
	resetCommand += cmdQuit
	_, _, err := t.sendCommand(resetCommand)
	return err
}

func (t *controller) close() {
	t.controlConn.Close()
}
