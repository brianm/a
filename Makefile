# set this to match the base package
PACKAGE=github.com/brianm/goskel

# where dependencies, etc, should check out to
WORKSPACE=$(PWD)/build

build: env
	GOPATH=$(WORKSPACE) go get $(PACKAGE)
	GOPATH=$(WORKSPACE) go install $(PACKAGE)/goskel

env:
	mkdir -p $(WORKSPACE)/src/$(PACKAGE)
	rm -r $(WORKSPACE)/src/$(PACKAGE) # remove last dir for symlink
	ln -s $(PWD) $(WORKSPACE)/src/$(PACKAGE)

clean:
	rm -rf build
