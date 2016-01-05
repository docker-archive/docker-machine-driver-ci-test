build: docker-machine-driver-ci-test
	
docker-machine-driver-ci-test: driver.go main.go
	GO15VENDOREXPERIMENT=1 go build .
	
clean:
	rm docker-machine-driver-ci-test