package virtualmachine

import( 
	"fmt"
	"bufio"
	"os"
)

//used to hold VM data after user input	
type VirtualMachine struct{
	Name string
	OSType string
	Ip string
	Enviornment string
	
}

func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

//read VM data from input and return a pointer to it
func NewVirtualMachine() *VirtualMachine {
	reader := bufio.NewReader(os.Stdin)
	
	//get VM name from user input	
	fmt.Print("Enter VM Name: ")
	vmName, _, err := reader.ReadLine()
	checkErr(err)
	virtualMachineName := string(vmName)

	//get VM OSType from user input		
	fmt.Print("Enter VM OSType: ")
	vmOSType, _, err := reader.ReadLine()
	checkErr(err)
	virtualMachineOSType := string(vmOSType)
	
	//get VM Ip from user input	
	fmt.Print("Enter VM Ip: ")
	vmIp, _, err := reader.ReadLine()
	checkErr(err)
	virtualMachineIp := string(vmIp)
	
	//get VM Enviornment from user input	
	fmt.Print("Enter VM Enviornment: ")
	vmEnviornment, _, err := reader.ReadLine()
	checkErr(err)
	virtualMachineEnviornment := string(vmEnviornment)
	
	//initialize VM with user input values
	virtualMachine := VirtualMachine{virtualMachineName, virtualMachineOSType, virtualMachineIp, virtualMachineEnviornment}
	
	return &virtualMachine

}
