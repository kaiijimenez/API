FROM golang:1.11.2
ADD . /go/src/github.com/kaiijimenez/API
WORKDIR /go/src/github.com/kaiijimenez/API
EXPOSE 8080
RUN  go get -u github.com/go-sql-driver/mysql && go get -u github.com/beego/bee 
CMD ["bee", "run", "-downdoc=true", "-gendoc=true"]
##check not working and delete the images created