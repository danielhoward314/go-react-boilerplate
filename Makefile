DIR := ${CURDIR}
init:
	go get & cd client && npm install

run separate:
	go run main.go -env=dev -port=8080 & cd client && npm run start-client

test prod:
	go run main.go -env=prod -port=8080 -html=$(DIR)/client/dist/index.html -webpack=$(DIR)/client/dist/js/
