
# Remote-Calculator
A Server/Client calculator with Golang 
## Usage
**To run TCP client/server version:**

    go run cmd/server/main.go

then run:

    go run cmd/client/main.go
problems must consist of **2 operand** and a **(*,+,-,/,%)** operation.
example: `2+3, 4 * 5`

**To run HTTP client/server version:**

    go run cmd/server/http/http-server.go
   then:
   

    go run cmd/client/http/http-client.go
