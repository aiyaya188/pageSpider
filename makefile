# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME= spider 
BINARY_UNIX=$(BINARY_NAME)_linux

all:  build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	 $(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/tebeka/selenium/log
	$(GOGET) github.com/jinzhu/gorm        
#	$(GOGET) github.com/markbates/pop
#	$(GOGET) github.com/go-redis/redis
#	$(GOGET) github.com/garyburd/redigo/redis
	$(GOGET) github.com/go-ini/ini
	$(GOGET) github.com/jinzhu/gorm/dialects/mysql  
	$(GOGET) github.com/garyburd/redigo/redis
	$(GOGET) github.com/cheekybits/genny/generic
	$(GOGET) github.com/go-vgo/robotgo 
	$(GOGET) github.com/advancedlogic/GoOse 






# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
