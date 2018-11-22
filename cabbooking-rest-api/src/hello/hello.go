package main

import (
    "fmt"
    "strconv"
)

func main() {
	
	
//	 To initialize a map with some data, use a map literal:

    // Location point
	commits := map[string]string{
	    "Mehrauli": "40.722,-73.989",
	    "Vasant_Kunj":   "40.750,-73.996",
	    "Vasant_Vihar": "41.250,-74.290",
	    "Munirka": "41.352, -74.495",
     }
	
	i := commits["Mehrauli"]
	
	fmt.Println("\n" )
	fmt.Println("Value of Key Mehrauli :", i )	
	fmt.Println("\n" )
	
	for key, value := range commits {
      fmt.Println("Key:", key, ", ","Value:", value)
    }
	
	
	//Location distance
	// Location point
	distance := map[string]string{
	    "Mehrauli": "1km",
	    "Vasant_Kunj":   "2km",
	    "Vasant_Vihar": "3km",
	    "Munirka": "4km",
     }
	
	j := distance["Mehrauli"]
	
	fmt.Println("\n" )
	fmt.Println("Value of Key Mehrauli :",j )	
	fmt.Println("\n" )
	
	for key, value := range distance {
      fmt.Println("Key:", key, ", ","Value:", value)
    }
	
	
//  -----------------------------------------------------------------------------
//  -----------------------------------------------------------------------------

//	    "location_name": "Mehrauli",
//	    "location": "40.722,-73.989"
	    
//	    "location_name": "Vasant_Kunj",
//	    "location": "40.750,-73.996"
	
//	    "location_name": "Vasant_Vihar",
//	    "location": "41.250,-74.290"
	    
//	    "location_name": "Munirka",
//	    "location": "41.352, -74.495"


		type CustOrder struct {
		    customer_name        string   `json:"customer_name,omitempty"`
		    customer_mobile_no   int64    `json:"customer_mobile_no,omitempty"`
		    source_loc_address   string   `json:"source_loc_address,omitempty"`
		    dest_loc_address     string   `json:"dest_loc_address,omitempty"`
		}
		
		
		
		type CabLocation struct {
		    driver_name           string   `json:"driver_name,omitempty"`
		    distance              string   `json:"distance,omitempty"`
		    location_name         string   `json:"location_name,omitempty"`
		    location              string   `json:"location,omitempty"`
		}
		
		
		
		type CabBookingDetails struct {
		    customer_name           string   `json:"customer_name,omitempty"`
		    customer_mobile_no      int64    `json:"customer_mobile_no,omitempty"`
		    source_loc_address      string   `json:"source_loc_address,omitempty"`
		    dest_loc_address        string   `json:"dest_loc_address,omitempty"`
		    driver_name             string   `json:"driver_name,omitempty"`
		    status                  string   `json:"status,omitempty"`
		}


        var custOrder []CustOrder        
        custOrder =append(custOrder,CustOrder{customer_name: "Albert",customer_mobile_no: 89786331,source_loc_address: "Mehrauli",dest_loc_address: "Karol Bagh"})        	
        fmt.Println("\n" )
	    fmt.Println("custOrder :", custOrder)	         
        
        
        var cabLocation []CabLocation        
        cabLocation =append(cabLocation,CabLocation{driver_name: "Alex",distance: "1KM",location_name: "Mehrauli",location: "40.722,-73.989"})        	
        fmt.Println("\n" )
	    fmt.Println("cabLocation :", cabLocation)	         
        
        
        var cabBookingDetails []CabBookingDetails        
        cabBookingDetails =append(cabBookingDetails,CabBookingDetails{customer_name: "Albert",customer_mobile_no: 89786331,source_loc_address: "Mehrauli",dest_loc_address: "Karol Bagh",driver_name: "Alex",status:"Booked"})        	
        fmt.Println("\n" )
	    fmt.Println("cabBookingDetails :", cabBookingDetails)	
        
	
        	
         str := "40.722,-73.989"

         // Step 1: Convert it to a rune
         latStr := string(str[0:6])    
         latInt, _ := strconv.ParseFloat(latStr,64)
         fmt.Println(latInt)         
         
         lonStr := string(str[7:14])
         lonInt, _ := strconv.ParseFloat(lonStr,64) 
         fmt.Println(lonInt)         
         
         fmt.Println("lat :" + latStr + " , " +  "lon :" + lonStr)
	
}