docker: 
	kubectl config view --minify --flatten > shareconfig.txt
	GOOS=linux go build -o shareconfig-server-linux
	docker build -t atakanyenel/shareconfig-server:latest .

compile: main.go
	go build