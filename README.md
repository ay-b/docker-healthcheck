# What?

This program intended to replace `curl` at work as a health-check for the certain docker container like:
```yaml
    healthcheck:
      test: healthcheck --url http://localhost:8080
      interval: 10s
      timeout: 5s
      retries: 6
```

# Why?

Because there are some cases when you can't or don't want to have `curl` in your deployment.  
Having web-services in containers often require checking they're alive and didn't crash for any reason.
`Healthcheck` is a 15-lines code that compiles to a relatively small binary that does a single job: query an url and return an error code.

So instead of having something like  
`apt update && apt install curl no-install-recommends -y && apt clean && rm -rf /var/lib/apt/lists/*`
You just copy compiled binary to the container and call it from compose healthcheck. Especially useful in pipelines that being run a thousand times a day to save some time and traffic.
# How?

Just `healthcheck --code "20\d" --url http://example.com:80`  
Or, if used as a healthcheck for docker, just run without args to check `http://localhost:8080` for code `200` (I'm assuming here that you have `8080` as your service port open).  

### Dockerfile.example
Here is an example how to embed `healthcheck` to your Dockerfile.  
Example is based on the standard Nginx container. Note the `HEALTHCHECK` part. [Idented CMD](https://docs.docker.com/engine/reference/builder/#healthcheck) is a mandatory to avoid overriding the "root" CMD.

### Legit args
`--code` accepts integers and regexps in quotes, e.g. "[2,3]\d\d" will accept all 2xx and 3xx codes.  
`--url` accepts url with protocol and port (optionally), e.g. `http://example.com:8080`, `https://domain.com:8443`

## Compilation advisory

Use `GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o healthcheck main.go` to have binary about 25% smaller than without them. Compiling using go v1.17.1+ will make binary even smaller.  
Use then `upx --brute healthcheck` to make it 3 times smaller. Note, that unpacking the binary on launch takes time and CPU. So 150 millissecond CPU spikes at monitoring every 10 seconds will come from this binary.
Using `upx -9 healthcheck` will leave a little bigger binary but startup will be about 55ms.
[Source for these tricks](https://stackoverflow.com/questions/4523920/how-do-i-update-a-formula-with-homebrew)

The binary added to the repo just for convenience.