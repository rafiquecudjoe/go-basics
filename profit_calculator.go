package main

import "fmt"

func main(){

  var revenue float64 

  var expenses float64

  var taxRate float64

  fmt.Print("Total Revenue: ")
  fmt.Scan(&revenue)

  fmt.Print("Total Expenses: ")
  fmt.Scan(&expenses)

  fmt.Print("Tax rate: ")
  fmt.Scan(&taxRate)

  var earningsBeforeTax = revenue - expenses

  var tax = (1 - taxRate/100) * earningsBeforeTax

  var profit = earningsBeforeTax - tax

  var ratio = earningsBeforeTax/profit

  fmt.Println(`Earnings before tax = `,earningsBeforeTax)
  fmt.Println(`Profit = `,profit)
  fmt.Println(`Ratio = `,ratio)


}