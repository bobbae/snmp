package myExample

import (
	"github.com/slayercat/gosnmp"
	"github.com/slayercat/GoSNMPServer"
	"fmt"
	"net"
)

var g_MyVal12345 int
var g_MyStringExample, g_IPAddress string

func init() {
	g_Logger = GoSNMPServer.NewDiscardLogger()
	g_MyVal12345 = 12345
	g_MyStringExample = "MyStringExample"
	g_IPAddress = "172.22.22.22"
}

var g_Logger GoSNMPServer.ILogger

//SetupLogger Setups Logger for this mib
func SetupLogger(i GoSNMPServer.ILogger) {
	g_Logger = i
}



func myExmpleOids() []*GoSNMPServer.PDUValueControlItem {
	return []*GoSNMPServer.PDUValueControlItem{
		{
			OID:  "1.3.6.1.4.12345.0.0.0.1",
			Type: gosnmp.Integer,
			OnGet: func() (value interface{}, err error) {
				return GoSNMPServer.Asn1IntegerWrap(g_MyVal12345), nil
			},
			OnSet: func(value interface{}) (err error) {
				val := GoSNMPServer.Asn1IntegerUnwrap(value)
				g_MyVal12345 = val
				return nil
			},
			Document: "My Test 12345",
		},
		{
			OID:  "1.3.6.1.4.12345.0.0.0.2",
			Type: gosnmp.OctetString,
			OnGet: func() (value interface{}, err error) {
				fmt.Println(g_MyStringExample)
				return GoSNMPServer.Asn1OctetStringWrap(g_MyStringExample), nil
			},
			OnSet: func(value interface{}) (err error) {
				//val := GoSNMPServer.Asn1OctetStringUnwrap(value)
				val := fmt.Sprintf("%s", value)
				g_MyStringExample = val
				return nil
			},
			Document: "My String Example",
		},
		{
			OID:  "1.3.6.1.4.12345.0.0.0.3",
			Type: gosnmp.IPAddress,
			OnGet: func() (value interface{}, err error) {
					vat :=net.ParseIP(g_IPAddress)
					
					return GoSNMPServer.Asn1IPAddressWrap(vat), nil
			},
			OnSet: func(value interface{}) (err error) {
					val := GoSNMPServer.Asn1IPAddressUnwrap(value)
					g_IPAddress = val.String()
					return nil
			},
			Document: "My Test IPAddress",
		},
	}
}

// All function provides a list of common used OID in DISMAN-EVENT-MIB
func All() []*GoSNMPServer.PDUValueControlItem {
	return myExmpleOids()
}
