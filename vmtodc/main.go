package main

import (
	"strings"
	"time"
    "fmt"
	"errors"
	"datacenter"
	"virtualmachine"
	"math/rand"
	
)

//used to hold data values after VM has been mapped to a DC
type vmToDataCenter struct
{
	vm *virtualmachine.VirtualMachine
	cluster string
	host string
	datastore string
	network string
}

//maps VM to a DC cluster
func determineCluster(v *virtualmachine.VirtualMachine, d *datacenter.DataCenter) (string, error) {
	
	switch v.OSType {
		case "linux": {
			if (v.Enviornment == "dev") {
					return d.Clusters[0].Name, nil
					} else {
						return " ", errors.New(fmt.Sprintf("No linux cluster found for VM in %s environment",  v.Enviornment))
					}
				
		}
		case "windows": {
			if(v.Enviornment == "qa") {
				return d.Clusters[1].Name, nil
			} else {
				return " ", errors.New(fmt.Sprintf("No windows cluster found for VM in %s environment",  v.Enviornment))
			}
		}
		default:
		 return " ", errors.New(fmt.Sprintf("No cluster found for VM with %s OStype in %s environment ",  v.OSType, v.Enviornment))
		
	}	
}

//maps VM to a DC host
func determineHost(v *virtualmachine.VirtualMachine, d *datacenter.DataCenter) (string, error) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5)
	
	switch v.OSType {
		case "linux": {
					return d.Clusters[0].Hosts[r], nil
				
		}
		case "windows": {
				return d.Clusters[1].Hosts[r], nil

		}
		default:
		 return " ", errors.New(fmt.Sprintf("Unable to find Host for VM %s", v.Name))
	}
	
	
}

//maps VM to a DC datastore
func determineDatastore(v *virtualmachine.VirtualMachine, d *datacenter.DataCenter) (string, error) {
	
	temp := make([]int64, 4)
	
	switch v.OSType {
		case "linux": {
					for i := 0; i < len(d.Clusters[0].DataStores); i++ {
						if(d.Clusters[0].DataStores[i].Maintainance != true) {
							temp[i] = d.Clusters[0].DataStores[i].Freespace
						}
					}
					maxElem := maxElement(temp)
					return d.Clusters[0].DataStores[maxElem].Name, nil

				
		}
		case "windows": {
				for i := 0; i < len(d.Clusters[1].DataStores); i++ {
						if(d.Clusters[1].DataStores[i].Maintainance != true) {
							temp[i] = d.Clusters[1].DataStores[i].Freespace
						}
					}
					maxElem := maxElement(temp)
					return d.Clusters[1].DataStores[maxElem].Name, nil
		}
		default:
		 return " ", errors.New(fmt.Sprintf("Unable to find ",  v.OSType, v.Enviornment))
	}
}

//maps VM to a DC network
func determineNetwork(v *virtualmachine.VirtualMachine, d *datacenter.DataCenter) (string, error) {

	tempIp := truncateString(v.Ip)	

	switch v.OSType {
		case "linux": {
					switch tempIp {
						case truncateString(d.Clusters[0].Networks[0].Vlan): {
							return d.Clusters[0].Networks[0].Name, nil
						}
						case truncateString(d.Clusters[0].Networks[1].Vlan): {
							return d.Clusters[0].Networks[1].Name, nil
						}
						case truncateString(d.Clusters[0].Networks[2].Vlan): {
							return d.Clusters[0].Networks[2].Name, nil
						}
						case truncateString(d.Clusters[0].Networks[3].Vlan): {
							return d.Clusters[0].Networks[3].Name, nil
						}									
						default: {
							return " ",errors.New(fmt.Sprintf("No VLAN network found for given IP %s for VM.",  v.Ip)) 
						}
					}
		}
		case "windows": {
			
					switch tempIp {
						case truncateString(d.Clusters[1].Networks[0].Vlan): {
							return d.Clusters[1].Networks[0].Name, nil
						}
						case truncateString(d.Clusters[1].Networks[1].Vlan): {
							return d.Clusters[1].Networks[1].Name, nil
						}
						case truncateString(d.Clusters[1].Networks[2].Vlan): {
							return d.Clusters[1].Networks[2].Name, nil
						}
						case truncateString(d.Clusters[1].Networks[3].Vlan): {
							return d.Clusters[1].Networks[3].Name, nil
						}									
						default: {
							return " ",errors.New(fmt.Sprintf("No VLAN network found for given IP %s for VM.",  v.Ip)) 
						}
					}
		}
		default:
		 return " ", errors.New(fmt.Sprintf("Unable to find ",  v.OSType, v.Enviornment))
	}
}

//helper function for determineDatastore, returns minmum element index of an array
func maxElement(s []int64) (int){
	max := s[0]
	maxElem := 0
	
	for i := 0; i < len(s); i++ {
		if(s[i] > max){
			max = s[i]
			maxElem = i
		}
	}
	
	return maxElem
}

//helper function for determineNetwork, truncates an ip string from 4 to 3 octets
func truncateString(s string) string{
	temp1 := strings.Split(s, ".")
	
	trunc := 2
	if( len(temp1[3]) > 1) {
		trunc = 3
	}
	
	temp2 := []byte(s)
	temp3 := make([]byte, len(s)-trunc)
	
	for i := 0; i < len(s)-trunc; i++ {
		temp3[i] = temp2[i]
	}
	
	return string(temp3);
}

//print assigned vm values
func vmAssignmentPrint(v vmToDataCenter) {

	fmt.Println("-++++++++-")
	fmt.Printf("\nName: %s\n", v.vm.Name)
	fmt.Printf("OSType: %s\n", v.vm.OSType)
	fmt.Printf("Ip: %s\n", v.vm.Ip)
	fmt.Printf("Enviornment: %s\n", v.vm.Enviornment)
	fmt.Printf("Cluster: %s\n", v.cluster)
	fmt.Printf("Host: %s\n", v.host)
	fmt.Printf("Datastore: %s\n", v.datastore)
	fmt.Printf("Network: %s\n\n", v.network)
	fmt.Println("-++++++++-")
}

//check panic errors
func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

//pick a cluster,  datastore, host and network and assign it to the following VMs according to the data associated with the VMs.
func main() {
	//initialize an instance of the datacenter with set data values
	dcWest := datacenter.NewDataCenter()
	
	//initialize a VM from user input
	vmInput := virtualmachine.NewVirtualMachine()
	
	//determine VMs assigned Cluster, Host, DataStore, and Network	
	vmCluster, err := determineCluster(vmInput, dcWest)
	checkErr(err)
	vmHost, err := determineHost(vmInput, dcWest)
	checkErr(err)
	vmDatastore, err := determineDatastore(vmInput, dcWest)
	checkErr(err)
	vmNetwork, err := determineNetwork(vmInput, dcWest)
	checkErr(err)
	
	//initialize VM assignment with VMinput, Cluster, Host, DataStore, and Network		
	vmAssignment := vmToDataCenter{vmInput, vmCluster, vmHost, vmDatastore, vmNetwork}
	
	//print VMs assigned Cluster, Host, DataStore, and Network
	vmAssignmentPrint(vmAssignment)
	
}