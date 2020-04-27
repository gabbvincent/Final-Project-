package main

import "fmt"

func maintenance(mileage int){

  var array2[9]int

 //Reccomended Maintenance per Mileage increments
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

   fmt.Println("At your current mileage of", mileage,"miles. You have: ")

   //Maintenance names
   array1[0] = "(Spark Plugs clean/check)"
   array1[1] = "(engine oil change)"
   array1[2] = "(oil filter change)"
   array1[3] = "(Clean and Lube Chain)"
   array1[4] = "(Check and Tighten nuts/bolts)"
   array1[5] = "(Brake fluid change)"
   array1[6] = "(Tire pressure check)"
   array1[7] = "(Air filter clean or replace)"
   array1[8] = "(Battery check or change)"

   
   //loop that calculates the amount of miles until maintenance
   for i:= 0; i <=8; i++{
    x := (mileage % array2[i])
    y := (x - array2[i])
    until := y * -1
// if there is less than 100 miles until maintenance is due then print this message.
  if until < 100 {
    fmt.Println()
    fmt.Println("[","You have less than 100 until", array1[i],"Don't forget!","]")
    } else {

    fmt.Println()
  //Prints message showing user the their maintenance schedule
    fmt.Println("[",until, "Miles until:", array1[i],"]")
    } 
  }
}

func main() {
//assign a variable for mileage
 var mileage int

 fmt.Println("To check your upcoming maintenance please enter your current vehicle mileage.")
  
  //scan for mileage
  fmt.Scanln(&mileage)

 //call maintenance
  maintenance(mileage)
}
