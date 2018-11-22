package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "fmt"
    "strconv"
    "github.com/olivere/elastic"
    "context"
)

// The CustOrder, CabLocation and CabBookingDetails Type
type CustOrder struct {
    CustomerName         string    `json:"customername,omitempty"`
    CustomerMobileNo     int64     `json:"customermobileno,omitempty"`
    SourceLocAddress     string    `json:"sourcelocaddress,omitempty"`
    DestLocAddress       string    `json:"destlocaddress,omitempty"`
}
type Location struct {
    Lat                   float64  `json:"lat,omitempty"`
    Lon                   float64  `json:"lon,omitempty"`
}
type CabLocation struct {
    DriverName           string    `json:"drivername,omitempty"`
    Distance             string    `json:"distance,omitempty"`
    LocationName         string    `json:"locationname,omitempty"`
    Location             *Location `json:"location,omitempty"`
}
type CabBookingDetails struct {
    CustomerName          string   `json:"customername,omitempty"`
    CustomerMobileNo      int64    `json:"customermobileno,omitempty"`
    SourceLocAddress      string   `json:"sourcelocaddress,omitempty"`
    DestLocAddress        string   `json:"destlocaddress,omitempty"`
    DriverName            string   `json:"drivername,omitempty"`
    Status                string   `json:"status,omitempty"`
}


// Struct Variable declaration
var custorder []CustOrder
var cablocation []CabLocation
var cabbookingdetails []CabBookingDetails 


/**
 ** Display all exsiting Cab locations 
 **/

func GetCabLocation(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    location := params["locationname"] 
    fmt.Println(" \nlocation " + location + " \n")
    //getLocationPoint method called
    getLocationPoint(location)
	
    json.NewEncoder(w).Encode(cablocation)
}



/**
 ** Book a cab ride 
 **/

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	
    //params := mux.Vars(r)
    var bookingetails CabBookingDetails
    _ = json.NewDecoder(r.Body).Decode(&bookingetails)
    
    cabbookingdetails = append(cabbookingdetails, bookingetails)
    
    fmt.Printf("\n")
    fmt.Printf("Customer booking for cab ride is accepted - ", cabbookingdetails)
    fmt.Printf("\n")
    
    json.NewEncoder(w).Encode(cabbookingdetails)
    
}


/**
 ** Below function get list of drivers by given location
 **/

func getDriverListByGivenlocation(ctx context.Context, client *elastic.Client,lat float64 ,lon float64 ,distance string){
	 
	 //ELASTICSEARCH Query 
	//distanceQuery := elastic.NewGeoDistanceQuery("location").GeoPoint(elastic.GeoPointFromLatLon(40.722, -73.989)).Distance("1km")	 
	distanceQuery := elastic.NewGeoDistanceQuery("location").GeoPoint(elastic.GeoPointFromLatLon(lat, lon)).Distance(distance)
	
	boolQuery := elastic.NewBoolQuery().Must(elastic.NewMatchAllQuery()).Filter(distanceQuery)

    searchResult, err := client.Search().Index("cablocation").Type("cab").Query(boolQuery).Do(ctx)
    CheckErr(err)

    if searchResult.Hits.TotalHits > 0 {
        fmt.Printf("Found a total of %d Driver \n", searchResult.Hits.TotalHits)

        for _, hit := range searchResult.Hits.Hits {
            var s CabLocation
            err := json.Unmarshal(*hit.Source, &s)
            if err != nil {
                // deserialization failed
            }
            
            cablocation = append(cablocation,s)
            fmt.Printf("\n")
            fmt.Printf("cablocation" , cablocation)
            fmt.Printf("\n")
        }
        
    } else {
        fmt.Print("Found no cabs for ride \n")
    }
}


/**
 ** Get location lattitue , longitude and distance from map
 **/

func getLocationPoint(location string){
	
	// Location point map
	pointsMap := map[string]string{
	    "Mehrauli": "40.722,-73.989",
	    "Vasant Kunj":   "40.750,-73.996",
	    "Vasant Vihar": "41.250,-74.290",
	    "Munirka": "41.352, -74.495",
     }
	
	//Location distance map
	distanceMap := map[string]string{
	    "Mehrauli": "1km",
	    "Vasant Kunj":   "2km",
	    "Vasant Vihar": "3km",
	    "Munirka": "4km",
     }
	
	// Get lattitue and longitude from pointsMap according to corresponding location key
	pointKey := pointsMap[location]
	fmt.Println("Value of  pointKey :" + pointKey + " \n")		
	
	// Get lattitue and longitude from distanceMap according to corresponding location key
	distanceKey := distanceMap[location]
	fmt.Println("Value of distanceKey :" + distanceKey + " \n" )	
	
    // Step 1: Convert it to a rune
    str := []rune(pointKey)

    // Step 2: get substring from pointKey string
    latStr := string(str[0:6])    
    latInt, _ := strconv.ParseFloat(latStr,64)
         
    lonStr := string(str[7:14])
    lonInt, _ := strconv.ParseFloat(lonStr,64) 
         
    /**
     ** ELASTICSEARCH context and client initialized
     **/
    
	ctx := context.Background()

	client, err := elastic.NewClient()
	CheckErr(err)
	
	/**
     ** getDriverListByGivenlocation(ctx, client)  method is called
     **/
	getDriverListByGivenlocation(ctx, client,latInt,lonInt,distanceKey)
}


/**
 ** Exception handling
 **/

func CheckErr(err error) {
    if err != nil {
        fmt.Println(err)
    }
}



/**
 ** main function to boot up everything
 **/

func main() {
	
	//RESTAPI Routes
    router := mux.NewRouter()
    router.HandleFunc("/cabride/{locationname}", GetCabLocation).Methods("GET")
    router.HandleFunc("/cabride/createbooking", CreateBooking).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
    
}


