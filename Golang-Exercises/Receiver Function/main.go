//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	Name              string
	Health, Maxhealth uint
	Energy, Maxenergy uint
}

func (player *Player) ModHealth(amount uint, action string) {
	if action == "add" {
		fmt.Println("Adding", amount, "of health to", player.Name)
		player.Health += amount
		if player.Health > player.Maxhealth {
			player.Health = player.Maxhealth
		}
	} else if action == "reduce" {
		fmt.Println("Reducing", amount, "of health of", player.Name)
		if player.Health-amount > player.Health {
			player.Health = 0
		} else {
			player.Health -= amount
		}
	} else {
		fmt.Println("Can't Perform any action!")
	}
	fmt.Println("After ModHealth :", player.Health)
}
func (player Player) ModEnergy(amount uint, action string) {
	if action == "add" {
		fmt.Println("Adding", amount, "of Energy to", player.Name)
		player.Energy += amount
		if player.Energy > player.Maxenergy {
			player.Energy = player.Maxenergy
		}
	} else if action == "reduce" {
		fmt.Println("Reducing", amount, "of Energy of", player.Name)
		if player.Energy-amount > player.Energy {
			player.Energy = 0
		} else {
			player.Energy -= amount
		}
	} else {
		fmt.Println("Can't Perform any action!")
	}
	fmt.Println("After ModEnergy :", player.Energy)
}
func main() {
	player := Player{"Soldier", 10, 100, 80, 100}
	fmt.Println("Player Stats:", player)
	fmt.Println("\nPlayer Earlier health:", player.Health)
	player.ModHealth(20, "reduce")
	fmt.Println("\nPlayer Earlier energy:", player.Energy)
	player.ModEnergy(20, "reduce")
}
