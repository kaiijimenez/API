# API
## Version 5 (Final Version) Completed
```
Second provider (from files)
They can be choose by two different configuration variables:
prov := beego.AppConfig.String("weatherprovider") or prov := beego.AppConfig.String("fileprovider")
Added three files:
(mexico, mx), (bogota, co), (paris, fr) to get data from these
```

## Commands to use 
Get dependencies need in the project with:
```
$dep init 
```

## Database migration:
$docker exec -it weatherapi bee migrate -conn="root:root@tcp(weatherdb:3306)/weatherapidb"

## Run tests:
Unit test and connection to DB 
```
go test github.com/kaiijimenez/API/... -v
```

## Running APP
```
docker-compose up -d --build
docker-compose up
```

## ENDPOINTS
- [GET] localhost:8080/v1/weather/r?city=$City&country=$Country where country is a two code characters in lowercase and city as string
- [PUT] localhost:8080/v1/scheduler/weather/r?city=$City&country=$Country same as above


## ENPOINT Swagger Docs
```
localhost:8080/swagger/
```

## ERRORS
Issue resolved. 

