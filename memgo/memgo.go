package memgo

import "net"

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

func (m *Memgo) Dispose() bool {
	msg := "STOP"
	res := m.writeMsg(msg)

	if res != false {
		m.conn.Close()
	}

	return res
}

func (m *Memgo) Set(key, value string) bool {
	msg := "SET " + key + " " + value
	return m.writeMsg(msg)
}
