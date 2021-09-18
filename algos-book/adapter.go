package algosbook

import (
	"fmt"
	"math/rand"
)

type IEnemyAttacker interface {
	fireWeapon() bool
	driveForward()
	assignDriver()
}

type EnemyTank struct{}

func (et EnemyTank) fireWeapon() bool {
	success := rand.Int()%2 == 0
	fmt.Printf("Fired Weapon. Hit? %v\n", success)
	return success
}

func (et EnemyTank) driveForward() {
	fmt.Println("Driving Tank Forward")
}

func (et EnemyTank) assignDriver() {
	fmt.Println("Tank Driver assigned")
}

type EnemyRobot struct{}

func (er EnemyRobot) smashTarget() bool {
	success := rand.Int()%2 == 0
	fmt.Printf("Smashing target. Worked? %v\n", success)
	return success
}

func (er EnemyRobot) walkForward() {
	fmt.Println("Robot is walking forward")
}

func (er EnemyRobot) assignPilot() {
	fmt.Println("Robot pilot assigned")
}

type EnemyRobotAdapter struct {
	enemyRobot EnemyRobot
}

func (era EnemyRobotAdapter) fireWeapon() bool {
	return era.enemyRobot.smashTarget()
}

func (era EnemyRobotAdapter) driveForward() {
	era.enemyRobot.walkForward()
}

func (era EnemyRobotAdapter) assignDriver() {
	era.enemyRobot.assignPilot()
}

func RunAdapterDemo() {
	fmt.Println("------------------------------\nRunning adapter demo...")

	var ea1, ea2 IEnemyAttacker

	ea1 = EnemyTank{}
	er := EnemyRobot{}
	ea2 = EnemyRobotAdapter{enemyRobot: er}

	ea1.assignDriver()
	ea1.driveForward()
	ea1.fireWeapon()

	ea2.assignDriver()
	ea2.driveForward()
	ea2.fireWeapon()

	ea1, ea2 = ea2, ea1

	ea1.assignDriver()
	ea1.driveForward()
	ea1.fireWeapon()

	ea2.assignDriver()
	ea2.driveForward()
	ea2.fireWeapon()

	fmt.Print("Adapter demo completed\n------------------------------\n")
}
