test:
	go test ./...

format:
	go fmt ./...

test_a:
	go run main.go https://github.com/1Shubham7/demo-repo-for-templ8

test_b:
	go run main.go -b=dev -d=new-proj https://github.com/1Shubham7/demo-repo-for-templ8

.PHONY: test format test_a test_b