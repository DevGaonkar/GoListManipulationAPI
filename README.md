Go Programming - List Manipulation with API

-To run the program
go mod tidy (install gin)
go run main.go (running the server)

Example Usage
-Run this in Terminal or Command Prompt

Examples:
curl -X POST "http://localhost:8080/add" -H "Content-Type: application/json" -d '{"number": 5}'

OR 

Invoke-WebRequest -Uri http://localhost:8081/add -Method POST -Headers @{"Content-Type"="application/json"} -Body '{"number": 5}'