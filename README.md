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

    $ go test -timeout 30s -run ^Test_customQueue_Front$ goExp/containers

    $ go test goExp/containers

    $ go run -gcflags -m start.go