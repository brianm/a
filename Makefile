# set this to match the base package, ie "github.com/brianm/goskel"
# this autodetection might work for you, or might now
PACKAGE=`git remote -v | grep push | grep origin | awk '{print $2}' | cut -d '@' -f 2 | tr ':' '/' | cut -f 1,2 -d '.'`

# where dependencies, etc, should check out to
WORKSPACE=$(PWD)/WORKSPACE

# change this to build your thing appropriately
# in this case it is building the "hello" binary
# yours will probably be different
build: deps
	GOPATH=$(WORKSPACE) go install $(PACKAGE)/hello

deps: env
	GOPATH=$(WORKSPACE) go get $(PACKAGE)

env:
	mkdir -p $(WORKSPACE)/src/$(PACKAGE)
	rm -r $(WORKSPACE)/src/$(PACKAGE) # remove last dir for symlink
	ln -s $(PWD) $(WORKSPACE)/src/$(PACKAGE)

clean:
	rm -rf $(WORKSPACE)/pkg/*

clean-workspace:
	rm -rf $(WORKSPACE)

