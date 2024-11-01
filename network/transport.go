package network

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte
}

type Transport interface {
	// 方法返回一个只读通道 类型为RPC
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
