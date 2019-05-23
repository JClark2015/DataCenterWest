package datacenter


type network struct{Name string; Vlan string}

type dataStore struct{Name string; Capacity int64; Freespace int64; Maintainance bool}

type cluster struct
{
	Name string
	Hosts []string
	DataStores []dataStore
	Networks []network
}

type DataCenter struct
{	
	Name string
	Clusters []cluster
	

}

//return a pointer to an instance of a hard coded data center
func NewDataCenter() *DataCenter {
		dataCenterName := "DC-West"
		
		//define cluster data
		linuxClusterName := "cluster-dcw-lnx-dev"
		linuxClusterHost := []string{"host-dcw-lnx-dev-1", "host-dcw-lnx-dev-2", "host-dcw-lnx-dev-3", "host-dcw-lnx-dev-4", "host-dcw-lnx-dev-5"}
		linuxClusterDataStore1 := dataStore{"ds-lnx-dev-1", 8796093022208, 1854907772928, false}
		linuxClusterDataStore2 := dataStore{"ds-lnx-dev-2", 5717460467712, 2794586464256, false}
		linuxClusterDataStore3 := dataStore{"ds-lnx-dev-3", 4398046511104, 3610875461632, false}
		linuxClusterDataStore4 := dataStore{"ds-lnx-dev-4", 4398046511104, 3820875461632, true}
		linuxClusterDataStores := []dataStore{linuxClusterDataStore1, linuxClusterDataStore2, linuxClusterDataStore3 ,linuxClusterDataStore4}
		linuxClusterNetwork1 := network{"153_dev_pg_1", "5.4.1.0"}
		linuxClusterNetwork2 := network{"154_dev_pg_2", "5.4.5.0"}
		linuxClusterNetwork3 := network{"179_dev_pg_3", "5.4.9.0"}
		linuxClusterNetwork4 := network{"180_dev_pg_4", "5.4.18.0"}
		linuxClusterNetworks := []network{linuxClusterNetwork1, linuxClusterNetwork2, linuxClusterNetwork3 ,linuxClusterNetwork4}
		linuxCluster := cluster{linuxClusterName, linuxClusterHost, linuxClusterDataStores, linuxClusterNetworks}
		
		windowsClusterName := "cluster-dcw-win-qa"
		windowsClusterHost := []string{"host-dcw-win-qa-1", "host-dcw-win-qa-2", "host-dcw-win-qa-3", "host-dcw-win-qa-4", "host-dcw-win-qa-5"}
		windowsClusterDataStore1 := dataStore{"ds-win-qa-1", 7311752327168, 3441869148160, false}
		windowsClusterDataStore2 := dataStore{"ds-win-qa-2", 9400824418304, 3651155738624, false}
		windowsClusterDataStore3 := dataStore{"ds-win-qa-3", 4398046511104, 3610875461632, true}
		windowsClusterDataStore4 := dataStore{"ds-win-qa-4", 5277655814144, 4044135927808, false}
		windowsClusterDataStores := []dataStore{windowsClusterDataStore1, windowsClusterDataStore2, windowsClusterDataStore3 ,windowsClusterDataStore4}
		windowsClusterNetwork1 := network{"8_qa_pg_1", "2.3.8.0"}
		windowsClusterNetwork2 := network{"52_qa_pg_2", "2.3.2.0"}
		windowsClusterNetwork3 := network{"53_qa_pg_3", "2.3.3.0"}
		windowsClusterNetwork4 := network{"129_qa_pg_4", "2.3.7.0"}
		windowsClusterNetworks := []network{windowsClusterNetwork1, windowsClusterNetwork2, windowsClusterNetwork3 ,windowsClusterNetwork4}
		windowsCluster := cluster{windowsClusterName, windowsClusterHost, windowsClusterDataStores, windowsClusterNetworks}
		
		
		//initialize clusters with defined cluster data
		dataCenterClusters := []cluster{linuxCluster, windowsCluster}
		
		//initialize DataCenter with clusters
		dataCenter := DataCenter{dataCenterName, dataCenterClusters}
		
		return &dataCenter
	}
