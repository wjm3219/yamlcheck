# yamlcheck

- A command line tool to check yaml file, in Golang.
- Based on [go-yaml/yaml](https://github.com/go-yaml/yaml) library, which I clone into this repo.

### Usage

- Add **$GOPATH/bin** to your **$PATH**
- go get github.com/wjm3219/yamlcheck
- yamlcheck [file] 

### Test

- test.yml is a simple yaml test file
- run 'yamlcheck test.yml', or other yaml file. Result will print from stderr.