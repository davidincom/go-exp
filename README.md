# Set go proxy to able to install go modules in China:

	$ go env -w GO111MODULE=on
	$ go env -w GOPROXY=https://goproxy.cn,direct
	
	Default settings:
	
	GO111MODULE="on"
    GOPROXY="https://proxy.golang.org,direct"

# Useful command to start:

    $ go mod init goExp

    $ go run start.go

    $ go mod tidy

    $ go build start.go