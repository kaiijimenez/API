# API

<<<<<<< HEAD
#Version 5 (Final Version) Completed
=======
# Version 4 (Final Version) 
>>>>>>> 3bbd0b215ec1ab011ade2dd3469cccb00dea75f9
Second provider (from files)
They can be choose by two different configuration variables:
    prov := beego.AppConfig.String("weatherprovider") or prov := beego.AppConfig.String("fileprovider")
Added three files (mexico, mx), (bogota, co), (paris, fr) to get data from these

# Commands to use 
Get dependencies need in the project with:
$dep init 

# Database migration:
$docker exec -it weatherapi bee migrate -conn="root:root@tcp(weatherdb:3306)/weatherapidb"

# Run tests:
go test github.com/kaiijimenez/API/... -v

Unit test and connection to DB 

# Running APP
docker-compose up -d --build
docker-compose up

# ENDPOINTS
[GET] localhost:8080/v1/weather/r?city=$City&country=$Country where country is a two code characters in lowercase and city as string
[PUT] localhost:8080/v1/scheduler/weather/r?city=$City&country=$Country same as above

# ENPOINT Swagger Docs
localhost:8080/swagger/


<<<<<<< HEAD
#ERRORS
Issue resolved. 
=======
# ERRORS
Had an issue with version 5 (worker pool) everytime that it was running accepted only one connection with the pool and the channel used got closed so had an error trying to run the testcases and testing the endpoints without exited the application. 
Will try to fin a solution even if it is out of time.
>>>>>>> 3bbd0b215ec1ab011ade2dd3469cccb00dea75f9

