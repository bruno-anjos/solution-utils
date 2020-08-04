module github.com/bruno-anjos/solution-utils

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect

)

replace github.com/bruno-anjos/archimedesHTTPClient v0.0.0-20200731165616-9aa4edba78b5 => ./../archimedesHTTPClient
