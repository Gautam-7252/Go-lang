//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printstatus(list map[string]int) {
	status := map[int]int{}
	for _, value := range list {
		switch value {
		case Online:
			status[Online] += 1
		case Offline:
			status[Offline] += 1
		case Maintenance:
			status[Maintenance] += 1
		case Retired:
			status[Retired] += 1
		}
	}
	fmt.Println("\nTotal servers :", len(list))
	fmt.Println("Online servers :", status[Online])
	fmt.Println("Offline servers :", status[Offline])
	fmt.Println("Maintenance servers :", status[Maintenance])
	fmt.Println("Retired servers :", status[Retired])
}
func main() {
	ServerList := map[string]int{
		"flipkart": Online,
		"darkstar": Online,
		"aiur":     Online,
		"amazon":   Online,
		"google":   Online,
	}
	printstatus(ServerList)
	// change server status of `darkstar` to `Retired`
	ServerList["darkstar"] = Retired
	printstatus(ServerList)
	// change server status of `aiur` to `Offline`
	ServerList["aiur"] = Offline
	printstatus(ServerList)
}
