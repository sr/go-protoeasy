all: test

deps:
	go get -d -v ./...

updatedeps:
	go get -d -v -u -f ./...

testdeps:
	go get -d -v -t ./...

updatetestdeps:
	go get -d -v -t -u -f ./...

build: deps
	go build ./...

install: deps
	go install ./...

launch: install
	protoeasy -d

proto: install
	protoeasy --go --grpc --go_import_path go.pedge.io/protoeasy .
	find . -name *\.pb\*\.go | xargs strip-package-comments

docker-build:
	docker build -t pedge/protoeasy .

docker-launch: docker-build
	docker run -d -p 6789:6789 pedge/protoeasy make launch

lint: testdeps
	go get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go' | grep -v '\.pb\.log\.go'); do \
		golint $$file; \
		if [ -n "$$(golint $$file)" ]; then \
			exit 1; \
		fi; \
	done

vet: testdeps
	go vet ./...

errcheck: testdeps
	go get -v github.com/kisielk/errcheck
	errcheck ./...

#pretest: lint vet errcheck
pretest: vet errcheck

test: testdeps pretest
	go test ./...

clean:
	go clean ./...
	rm -rf example/out


example: install
	rm -rf _example-out
	protoeasy \
		--out=_example-out \
		--cpp --cpp_rel_out=cpp \
		--csharp --csharp_rel_out=csharp \
		--objectivec --objectivec_rel_out=objectivec \
		--python --python_rel_out=python \
		--ruby --ruby_rel_out=ruby \
		--go --go_rel_out=go --go_import_path=go.pedge.io/protoeasy/example/example-out/go \
		--grpc \
		--grpc-gateway \
		example

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	install \
	launch \
	proto \
	docker-build \
	docker-launch\
	lint \
	vet \
	errcheck \
	pretest \
	test \
	clean \
	example
