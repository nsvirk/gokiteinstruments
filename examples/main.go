package main

import (
	"fmt"

	kiteinstruments "github.com/nsvirk/gokiteinstruments"
)

func main() {

	instruments := kiteinstruments.NewManager()

	fmt.Println("------------------------------------------------------")
	fmt.Println("Count of instruments:", instruments.Count())
	fmt.Println("------------------------------------------------------")
	fmt.Println("GetByID: 'NSE:IOC'")
	fmt.Println("------------------------------------------------------")
	instDetails, err := instruments.GetByID("NSE:IOC")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Instrument Details: %+v\n", instDetails)
	fmt.Println("------------------------------------------------------")
	fmt.Println("GetByExchTradingsymbol: 'NSE' 'INFY'")
	fmt.Println("------------------------------------------------------")
	instDetails, err = instruments.GetByExchTradingsymbol("NSE", "INFY")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Instrument Details: %+v\n", instDetails)
	fmt.Println("------------------------------------------------------")
	fmt.Println("GetbyInstToken:'408065'")
	fmt.Println("------------------------------------------------------")
	instDetails, err = instruments.GetByExchTradingsymbol("NSE", "INFY")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Instrument Details: %+v\n", instDetails)
	fmt.Println("------------------------------------------------------")
}
