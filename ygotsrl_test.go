package ygotsros

import (
	"fmt"
	"testing"

	"github.com/openconfig/ygot/ygot"
)

func TestDevice(t *testing.T) {
	device := &Device{}
	device.GetOrCreateConfigure()
	service := device.Configure.GetOrCreateService()
	service.GetOrCreateCustomer("nokia")
	service.GetOrCreateVprn("1")
	vprn := service.GetVprn("1")
	vprn.AdminState = E_NokiaTypesSros_AdminState(2)
	vprn.Ecmp = ygot.Uint32(64)
	vprn.GetOrCreateAggregates()
	vprn.GetOrCreateAggregates().GetOrCreateAggregate("192.0.2.0/24")
	port := device.Configure.GetOrCreatePort("1/1/1")
	port.AdminState = E_NokiaTypesSros_AdminState(1)
	port.Description = ygot.String("test description")
	port.GetOrCreateEthernet()
	port.Ethernet.Mode = E_NokiaTypesPort_Mode(3)
	port.Ethernet.EncapType = E_NokiaTypesPort_EthernetEncapType(3)

	j, err := ygot.EmitJSON(device, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
		// debug
		SkipValidation: false,
	})
	if err != nil {
		return
	}
	fmt.Println(j)
}
