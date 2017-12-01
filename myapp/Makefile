LASTFILENAME = "mainsevr"
package : goBuild
	docker build -t myapp:3 .  
goBuild:
	go build -o $(LASTFILENAME) main.go && \
	cd ./echoapp/ && \
	go build -o myappcli Myclient.go
	
clean:
		--rm -rf $(LASTFILENAME) myappcli \
		docker rm -f myapp && \
		docker rmi myapp:3