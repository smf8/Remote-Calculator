

# Remote-Calculator
A Server/Client calculator with Golang
## Usage
**To run TCP client/server version:**

    go run cmd/server/main.go

Then run:

    go run cmd/client/main.go
Problems must consist of **2 operands** and a **(*,+,-,/,%)** operation.
Example: `2+3, 4 * 5`

**To run HTTP client/server version:**

    go run cmd/server/http/http-server.go
   If you want to use another client with this server:
   You must send problems as a **POST** request to `localhost:1234/solve` with following JSON format:


    {
	    "first_operand":0,
        "second_operand":0,
        "operator":"",
        "result":0
    }

   **OR**
   You can run client with:


    go run cmd/client/http/http-client.go
