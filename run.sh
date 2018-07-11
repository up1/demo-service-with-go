export GOPATH=`pwd`
go get gopkg.in/mgo.v2
gofmt -w src/
go test -v -cover -coverprofile=coverage.out  ./...
go tool cover -html=coverage.out
go install main

#sh run.sh
