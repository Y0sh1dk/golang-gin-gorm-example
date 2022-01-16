package command

import "fmt"

type DB struct {
}

func (db DB) Seed() {
	fmt.Println("Seeding the database...")
}

func (db DB) Reset() {
	fmt.Println("Resetting the database...")
}
