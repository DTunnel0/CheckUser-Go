package connection

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type AUXOpenVPNConnection interface {
	Connect()
	Send(data string)
	Receive(size int) string
	Close()
}

type openVPNConnection struct {
	host   string
	port   int
	socket net.Conn
}

func NewAUXOpenVPNConnection(host string, port int) AUXOpenVPNConnection {
	return &openVPNConnection{
		host: host,
		port: port,
	}
}

func (vpn *openVPNConnection) Connect() {
	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", vpn.host, vpn.port))
	vpn.socket = conn
}

func (vpn *openVPNConnection) Send(data string) {
	if vpn.socket == nil {
		return
	}
	vpn.socket.Write([]byte(data))
}

func (vpn *openVPNConnection) Receive(size int) string {
	if vpn.socket == nil {
		return ""
	}
	data := make([]byte, size)
	vpn.socket.Read(data)
	return string(data)
}

func (vpn *openVPNConnection) Close() {
	if vpn.socket == nil {
		return
	}
	vpn.socket.Close()
}

type OpenVPNConnection struct {
	connection AUXOpenVPNConnection
	next       contract.Connection
}

func NewOpenVPNConnection(connection AUXOpenVPNConnection) contract.Connection {
	return &OpenVPNConnection{connection: connection}
}

func (vpn *OpenVPNConnection) CountByUsername(ctx context.Context, username string) int {
	vpn.connection.Connect()
	defer vpn.connection.Close()

	vpn.connection.Send("status\n")
	data := vpn.connection.Receive(1024)
	count := strings.Count(data, username) / 2
	if vpn.next != nil {
		count += vpn.next.CountByUsername(ctx, username)
	}
	return count
}

func (vpn *OpenVPNConnection) Count(ctx context.Context) int {
	vpn.connection.Connect()
	defer vpn.connection.Close()

	vpn.connection.Send("status\n")
	data := vpn.connection.Receive(1024)
	regex := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3},\w+,)`)
	all := len(regex.FindAll([]byte(data), -1))
	if vpn.next != nil {
		all += vpn.next.Count(ctx)
	}
	return all
}

func (vpn *OpenVPNConnection) SetNext(next contract.Connection) {
	vpn.next = next
}
