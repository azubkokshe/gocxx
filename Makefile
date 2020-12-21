all:
	cd src; make dynamic 
	go build -o bin/cgo main.go
	bin/cgo