# What?

This program intended to replace `curl` at work as a health-check for the certain docker container like:
```yaml
    healthcheck:
      test: healthcheck --url "http://localhost:8080/x/app" --code "[2,3]\d{2}"
      interval: 10s
      timeout: 5s
      retries: 6
      start_period: 10s
```

# Why?

Because there are some cases when you can't or don't want to have `curl` in your deployment.  
Having web-services in containers often require checking they're alive and didn't crash for any reason.
`Healthcheck` is a 15-lines code that compiles to a relatively small binary that does a single job: query an url and return an error code.

So instead of having something like  
`apt update && apt install curl --no-install-recommends -y && apt clean && rm -rf /var/lib/apt/lists/*`
you just copy the compiled binary to the container and call it from the docker-compose healthcheck. Especially useful in pipelines that being run a thousand times a day to save some time and traffic.
# How?

Just `healthcheck --code "20\d" --url http://example.com:80`  
Or, if used as a healthcheck for docker, just run without args to check `http://localhost:8080` for code `200` (I'm assuming here that you have `8080` as your service port open).  

### Dockerfile.example
Here is an example how to embed `healthcheck` to your Dockerfile.  
Example is based on the standard Nginx container. Note the `HEALTHCHECK` part. [Idented CMD](https://docs.docker.com/engine/reference/builder/#healthcheck) is a mandatory to avoid overriding the "root" CMD.

### Legit args
`--code string` This should be a regexp in "" or clear integer of the expected return code e.g., `"[2,3]0\d"` will accept all 2xx and 3xx codes. Type `"200"` to expect a strict code (default "200")  
`--url string` This should be an URL. Can contain a path e.g., 'http://localhost:8080/x/app'. (default "http://localhost")

## Compilation advisory

Use `GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o healthcheck main.go` to have binary about 25% smaller than without them. Compiling using go v1.17.1+ will make binary even smaller.  
Use then `upx --brute healthcheck` to make it 3 times smaller. Note, that unpacking the binary on launch takes time and CPU. So 150 millisecond CPU spikes at monitoring every 10 seconds will come from this binary.  
Using `upx -9 healthcheck` will leave a slightly bigger binary but startup will be about 55ms.
[Source for these tricks](https://stackoverflow.com/questions/4523920/how-do-i-update-a-formula-with-homebrew)


# Download binary
Download pre-compiled binary from the [Releases page](https://github.com/ay-b/docker-healthcheck/releases)