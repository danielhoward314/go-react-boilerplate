DIR := ${CURDIR}
init:
	go get & cd client && npm install

run separate:
	go run main.go & cd client && npm run start-client

test prod:
	go run main.go -html=$(DIR)/client/dist/index.html -webpack=$(DIR)/client/dist/js/
