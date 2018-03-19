package main

import "fmt"

func main(){

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	
	for country := range countryCapitalMap{    //this same as part2
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	for country,city := range countryCapitalMap{ // this same as part1
                fmt.Println("Capital of",country,"is",city)
        }

	captial,ok := countryCapitalMap["United States"]
	if (ok){ // <==> if ok
		fmt.Println("Capital of United States is", captial)  
	}else {
		fmt.Println("Capital of United States is not present") 
	}

}
