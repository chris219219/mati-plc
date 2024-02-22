package main

import (
	"net"
	//"strconv"
	"github.com/vishvananda/netlink"
	"go.bug.st/serial"
)

type IPBind struct {
	Name      string
	IP        net.IP
	Mask      net.IPMask
	Gateway   net.IP
	Broadcast net.IP
}

type IFace struct {
	Name        string
	IPBinds    []IPBind
}

type SerialConn struct {
	Name     string
	Device   string
	Baudrate uint32
	IsOpen   bool
	port     serial.Port
}

func CreateSerialConn(name string, device string, baudrate uint32) SerialConn {
	var conn SerialConn
	conn.Name = name
	conn.Device = device
	conn.Baudrate = baudrate
	conn.IsOpen = false
	conn.port = nil
	return conn
}

func (conn SerialConn) ConnectSerial() (SerialConn, error) {
	mode := &serial.Mode { BaudRate: int(conn.Baudrate) }
	port, err := serial.Open(conn.Device, mode)
	if (err != nil) {
		return conn, nil
	}
	conn.port = port
	conn.IsOpen = true
	return conn, nil
}

/*
func (conn SerialConn) CloseSerial() (SerialConn, error) {

}
*/

func GetCurrIFaces() ([]IFace, error) {
	// get current interfaces from netlink
	links, err := netlink.LinkList()
	if (err != nil) {
		return nil, err
	}
	var ifaces []IFace

	// create an iface for each link
	for _, link := range links {
		var iface IFace;
		iface.Name = link.Attrs().Name

		netlinkaddrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
		if (err != nil) {
			return nil, err
		}

		// create an ipbind for each address in a link
		for _, netlinkaddr := range netlinkaddrs {
			var ipbind IPBind;
			ipbind.Name = netlinkaddr.String()
			ipbind.IP = netlinkaddr.IP
			ipbind.Mask = netlinkaddr.Mask
			ipbind.Broadcast = netlinkaddr.Broadcast
			iface = iface.AddIPBind(ipbind)
		}

		ifaces = append(ifaces, iface)
	}

	return ifaces, nil
}

func CreateIPBind(name string, ip string, mask string, gateway string, broadcast string) IPBind {
	var res IPBind
	res.Name = name
	res.IP = net.ParseIP(ip)
	res.Mask = net.IPMask(net.ParseIP(mask))
	res.Gateway = net.ParseIP(gateway)
	res.Broadcast = net.ParseIP(broadcast)
	return res
}

func (iface IFace) AddIPBind(bind IPBind) IFace {
	iface.IPBinds = append(iface.IPBinds, bind)
	return iface
}

func (iface IFace) String() string {
	res := "Interface: " + iface.Name + "\n"
	for _, b := range iface.IPBinds {
		res +=
			"  IP Bind: " + b.Name + "\n" +
				"    IP: " + b.IP.String() + "\n" +
				"    Subnet Mask: " + net.IP(b.Mask).String() + "\n" +
				"    Gateway: " + b.Gateway.String() + "\n" +
				"    Broadcast: " + b.Broadcast.String() + "\n"
	}
	return res
}
