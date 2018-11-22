# Finaccel-rest-api
Code - Elasticsearch, Go language

Finaccel-rest-api being developed using Go language and Elasticsearch.Below mentioned API is created.

1). Search all driver for a given location.

    URL:- http://localhost:8000/cabride/{location}
    
2). Book the customer trip and send them the booking acknowledgment message.
    
    URL:- http://localhost:8000/cabride/createbooking
          {
            "customername": "Albert",
            "customermobileno": 89786331,
            "sourcelocaddress": "Mehrauli",
            "destlocaddress": "Karol Bagh",
            "drivername": "Alex",
            "status": "Booked"
           }
           
     - Here Customer booking accepted and add into array and printed it on console, 
       due to complexity of schema dependency and time contraint.
       
  
