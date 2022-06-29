package myExample

import "github.com/slayercat/gosnmp"
import "github.com/slayercat/GoSNMPServer"


var g_MyVal12345 int

func init() {
	g_Logger = GoSNMPServer.NewDiscardLogger()
	g_MyVal12345 = 12345
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
			Document: "MyTest12345",
		},
	}
}

// All function provides a list of common used OID in DISMAN-EVENT-MIB
func All() []*GoSNMPServer.PDUValueControlItem {
	return myExmpleOids()
}
