package main

import "fmt"

func main(){

  var revenue float64 

  var expenses float64

  var taxRate float64


 revenue =  userInfo("Total Revenue: ")
 expenses = userInfo("Total Expenses: ")
 taxRate = userInfo("Tax rate: ")

  var earningsBeforeTax,profit,ratio = calculateRevenue(revenue,expenses,taxRate)

  fmt.Println(`Earnings before tax = `,earningsBeforeTax)
  fmt.Println(`Profit = `,profit)
  fmt.Println(`Ratio = `,ratio)


}

func userInfo(infoText string) float64 {

	var userInput float64

	  fmt.Print(infoText)
      fmt.Scan(&userInput)

	  return userInput

}

func calculateRevenue(revenue float64, expenses float64, taxRate float64)(float64,float64,float64){
  var earningsBeforeTax = revenue - expenses

  var tax = (1 - taxRate/100) * earningsBeforeTax

  var profit = earningsBeforeTax - tax

  var ratio = earningsBeforeTax/profit

  return earningsBeforeTax,profit,ratio
}