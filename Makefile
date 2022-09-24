mongo:
	docker run --name mongo -d -p 127.0.0.1:27017:27017 mongo:4.4
stopmongo:
	docker stop mongo
	docker rm mongo
run:mongo
	go run main.go server