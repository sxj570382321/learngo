
package main

import "fmt"

func main(){

        countryCapitalMap := map[string]string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
	fmt.Println("origin map")

        for country,_ := range countryCapitalMap{   // <==> country:= range countryCapitalMap
                fmt.Println("Capital of",country,"is",countryCapitalMap[country])
        }

	delete(countryCapitalMap,"France") // not france
	fmt.Println("Entry for France is deleted")
	fmt.Println("after delete entry France")

	
        for country := range countryCapitalMap{   // <==> country,_ := range countryCapitalMap   
                fmt.Println("Capital of",country,"is",countryCapitalMap[country])
        }
}
