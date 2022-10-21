# Set go proxy to able to install go modules in China:

	$ go env -w GO111MODULE=on
	$ go env -w GOPROXY=https://goproxy.cn,direct
	
	Default settings:
	
	GO111MODULE="on"
    GOPROXY="https://proxy.golang.org,direct"

# Go Best Practices

    https://go.dev/talks/2013/bestpractices/

1. Avoid nesting by handling errors first
2. Avoid repetition when possible
3. Important code goes first
4. Document your code
5. Shorter is better
6. Packages with multiple files
7. Make your packages "go get"-able
8. Ask for what you need
9. Keep independent packages independent
10. Avoid concurrency in your API
11. Use goroutines to manage state
12. Avoid goroutine leaks

# Useful command to start:

    $ go mod init goExp

    $ go run start.go

    $ go mod tidy

    $ go build start.go

    $ go test -timeout 30s -run ^Test_customQueue_Front$ goExp/containers

    $ go test goExp/containers

    $ go run -gcflags -m start.go

    -- workspace --

    $ go work init ./hello

    $ go work use ./libs

# Go Talk

- Concurrency Patterns: https://go.dev/talks/2012/concurrency.slide
- Advanced Concurrency Patterns: https://go.dev/talks/2013/advconc.slide

# How to publish your Go module, make it go-get able

1. Give a proper name for your module: github.com/{your github account name}/{your github repo name}
 
        eg: github.com/davidiwu/go-exp

2. After code submitted, and ready to release, give it a version number in the form of a git tag

        $ git tag v1.0.0
        $ git push origin v1.0.0

3. The version number should follow the semantic versioning model:

        eg: v{major}.{minor}.{patch}[-{pre-release identifier}]        
        v1.4.0-beta.2