package main

import "net"
import "strconv"

type IP4Bind struct {
	Name      string
	IP        net.IP
	Mask      net.IPMask
	Gateway   net.IP
	Broadcast net.IP
}

type SerialBind struct {
	Name     string
	Device   string
	Baudrate uint32
}

type IFace struct {
	Name        string
	IP4Binds    []IP4Bind
	SerialBinds []SerialBind
}

func CreateIP4Bind(name string, ip string, mask string, gateway string, broadcast string) IP4Bind {
	var res IP4Bind
	res.Name = name
	res.IP = net.ParseIP(ip)
	res.Mask = net.IPMask(net.ParseIP(mask))
	res.Gateway = net.ParseIP(gateway)
	res.Broadcast = net.ParseIP(broadcast)
	return res
}

func CreateSerialBind(name string, device string, baudrate uint32) SerialBind {
	var res SerialBind
	res.Name = name
	res.Device = device
	res.Baudrate = baudrate
	return res
}

func (iface IFace) AddIP4Bind(bind IP4Bind) IFace {
	iface.IP4Binds = append(iface.IP4Binds, bind)
	return iface
}

func (iface IFace) AddSerialBind(bind SerialBind) IFace {
	iface.SerialBinds = append(iface.SerialBinds, bind)
	return iface
}

func (iface IFace) String() string {
	res := "Interface: " + iface.Name + "\n"
	for _, b := range iface.IP4Binds {
		res +=
			"  IP Bind: " + b.Name + "\n" +
				"    IP: " + b.IP.String() + "\n" +
				"    Subnet Mask: " + net.IP(b.Mask).String() + "\n" +
				"    Gateway: " + b.Gateway.String() + "\n" +
				"    Broadcast: " + b.Broadcast.String() + "\n"
	}
	for _, b := range iface.SerialBinds {
		res +=
			"  Serial Bind: " + b.Name + "\n" +
				"    Device: " + b.Device + "\n" +
				"    Baudrate: " + strconv.FormatUint(uint64(b.Baudrate), 10) + "\n"
	}
	return res
}
