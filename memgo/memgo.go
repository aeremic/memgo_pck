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

const (
	EMPTY_KEY = "EMPTY_KEY"
	NOT_FOUND = "NOT_FOUND"
	EMPTY     = "EMPTY"
	SUCCESS   = "SUCCESS"
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

func (m *Memgo) sendCommand(msg string) string {
	w := m.writeMsg(msg)
	if w == true {
		r := m.receiveMsg()
		return r
	}

	return ""
}

func (m *Memgo) Dispose() bool {
	msg := STOP
	w := m.writeMsg(msg)

	if w == false {
		return false
	}

	r := m.receiveMsg()
	if r != SUCCESS+"\n" {
		m.conn.Close()
		return false
	}

	return true
}

func (m *Memgo) Set(key, value string) bool {
	return m.sendCommand(SET+" "+key+" "+value) == SUCCESS+"\n"
}

func (m *Memgo) GetAll() string {
	return m.sendCommand(GETALL)
}

func (m *Memgo) Get(key string) string {
	return m.sendCommand(GET + " " + key)
}

func (m *Memgo) DeleteAll() bool {
	return m.sendCommand(DELETEALL) == SUCCESS+"\n"
}

func (m *Memgo) Delete(key string) bool {
	return m.sendCommand(DELETE+" "+key) == SUCCESS+"\n"
}

func (m *Memgo) GetByPath(key, path string) string {
	return m.sendCommand(GETBYPATH + " " + key + " " + path)
}
