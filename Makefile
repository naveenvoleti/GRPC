URL := http://localhost:8080/chat

input := main.go

output := main.exe

hello:
	curl $(URl)/hello

build:
	go build $(input)

runServer:
	echo "Running"
	go run server/$(input)

runClient:
	echo "Running"
	go run client/$(input)

clean:
	rm $(output)
	echo "Deleted"