package main

import (
	"net"
	"strconv"
	"github.com/vishvananda/netlink"
)

type IPBind struct {
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
	IPBinds    []IPBind
	SerialBinds []SerialBind
}

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

func CreateSerialBind(name string, device string, baudrate uint32) SerialBind {
	var res SerialBind
	res.Name = name
	res.Device = device
	res.Baudrate = baudrate
	return res
}

func (iface IFace) AddIPBind(bind IPBind) IFace {
	iface.IPBinds = append(iface.IPBinds, bind)
	return iface
}

func (iface IFace) AddSerialBind(bind SerialBind) IFace {
	iface.SerialBinds = append(iface.SerialBinds, bind)
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
	for _, b := range iface.SerialBinds {
		res +=
			"  Serial Bind: " + b.Name + "\n" +
				"    Device: " + b.Device + "\n" +
				"    Baudrate: " + strconv.FormatUint(uint64(b.Baudrate), 10) + "\n"
	}
	return res
}
