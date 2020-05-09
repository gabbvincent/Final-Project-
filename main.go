//Programmers Name: Vincent G.
//Description: Final project

package main

import "fmt"

//Calculates Miles until 
func milesUntil(x int, y int) int {

 a := x % y
 
 return y - a

}

func main() {

var array2[9]int

//Reccomended Maintenance per # of Miles
array2[0] = 4000.0 //(Spark Plugs: clean/check)
array2[1] = 3500.0 //(engine oil change)
array2[2] = 7000.0 //(oil filter change)
array2[3] = 500.0 //(Clean/Lube Chain)
array2[4] = 6000.0 //(Check/Tighten nuts/bolts)
array2[5] = 15000.0 //(Brake fluid change)
array2[6] = 500.0 //(Tire pressure)
array2[7] = 10000.0 //(air filter)
array2[8] = 25000.0 //(battery check/change)
 
var array1[9]string

//Maintenance
array1[0] = "(Spark Plugs clean/check)"
array1[1] = "(engine oil change)"
array1[2] = "(oil filter change)"
array1[3] = "(Clean and Lube Chain)"
array1[4] = "(Check and Tighten nuts/bolts)"
array1[5] = "(Brake fluid change)"
array1[6] = "(Tire pressure check)"
array1[7] = "(Air filter clean or replace)"
array1[8] = "(Battery check or change)"

var x int

//Get Vehicle Mileage
fmt.Println("To check your upcoming reccomended maintenance, Enter your current mileage?")

fmt.Scanln(&x)

fmt.Println()
fmt.Println("At your current Mileage of", x, "you have: ")

//Print Miles until maintenance.
 for i:=0;i<=8;i++{
   
 fmt.Println()
 fmt.Println("[",milesUntil(x, array2[i]), "Miles until: ", array1[i], "is due.","]")

 if milesUntil(x, array2[i]) <= 100{

   fmt.Println("[ DON'T FORGET! ]")

   }

 }
 
}
