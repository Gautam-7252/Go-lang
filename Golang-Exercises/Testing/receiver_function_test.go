// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func newPlayer() Player {
	return Player{
		Name:      "test",
		Health:    100,
		Maxhealth: 100,
		Energy:    500,
		Maxenergy: 500,
	}
}
func TestModhealth(t *testing.T) {
	player := newPlayer()
	player.ModHealth(100, "add")

	if player.Health > player.Maxhealth {
		t.Errorf("Player's Health %v is greater than its limit i.e., %v .", player.Health, player.Maxhealth)
	}
	player.ModEnergy(500, "add")
	if player.Energy > player.Energy {
		t.Errorf("Player's Energy %v is greater than its limit i.e., %v .", player.Energy, player.Maxenergy)
	}
	player.ModHealth(101, "reduce")
	if player.Health < 0 {
		t.Errorf("Player's Health %v is Lower than 0.", player.Health)
	}
	player.ModEnergy(600, "reduce")
	if player.Energy < 0 {
		t.Errorf("Player's Energy %v is Lower than 0.", player.Energy)
	}
}
