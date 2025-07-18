package memgo

import (
	"bufio"
	"net"
)

const (
	STOP      = "STOP"
	GET       = "GET"
	GETALL    = "GETALL"
	SET       = "SET"
	DELETE    = "DELETE"
	DELETEALL = "DELETEALL"
	GETBYPATH = "GETBYPATH"
)

type Memgo struct {
	addr *net.TCPAddr
	conn *net.TCPConn
}

func NewMemgo(host, port string) (*Memgo, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		// TODO: Log
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		// TODO: Log
		return nil, err
	}

	return &Memgo{
		addr: tcpAddr,
		conn: conn,
	}, nil
}

func (m *Memgo) writeMsg(msg string) bool {
	res, err := m.conn.Write([]byte(msg + "\n"))
	if err != nil {
		// TODO: Log
		return false
	}

	if res == 0 {
		// TODO: Log
		return false
	}

	return true
}

func (m *Memgo) receiveMsg() string {
	buffer := bufio.NewReader(m.conn)
	bytes, err := buffer.ReadBytes('\n')
	if err != nil {
		// TODO: Log
		return ""
	}

	return string(bytes)
}

func (m *Memgo) Dispose() bool {
	msg := STOP
	w := m.writeMsg(msg)

	if w == false {
		return false
	}

	r := m.receiveMsg()
	if r != "Success\n" {
		m.conn.Close()
		return false
	}

	return true
}

func (m *Memgo) Set(key, value string) bool {
	msg := SET + " " + key + " " + value

	w := m.writeMsg(msg)
	if w == true {
		r := m.receiveMsg()
		return r == "Success\n"
	}

	return false
}

func (m *Memgo) GetAll() string {
	msg := GETALL

	w := m.writeMsg(msg)
	if w == true {
		r := m.receiveMsg()
		return r
	}

	return ""
}
