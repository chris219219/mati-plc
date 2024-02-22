## mati-plc

In order to easily configure a controller's interfaces with IP addresses and
serial connections, there should be an app which can send this data to the
controller.

Both the controller and phone will use json as an intermediate data transfer
format to communicate.

The controller sends json data about its existing interfaces to a
microcontroller.

The microcontroller parses this information and sends it to an NFC tag, which
sends it to a phone when read by an app.

The phone app processes this data and creates menus for manipulating the
controller's interfaces.

After the user has interacted with these menus and is ready to change the
interface settings of the controller, they hold the phone up to the NFC tag to
write the data.

The NFC tag sends this data to the microcontroller to parse the data into json
format.

The microcontroller sends this data to the controller, where it is used to
change its interface settings.

## Project Properties

This project is for the PLC (Programmable Logic Controller), AKA the controller.
It is written in Go to make it more maintainable and easy to read.

Entry Point: main.go
Interface Code: iface.go
WIP

## Dependencies

#### <github.com/goccy/go-json>
A lightweight alternative to encoding/json in Go's standard library.

#### <github.com/vishvananda/netlink>
Only for Linux, allows for manipulation of Linux interfaces.