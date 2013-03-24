# set this to match the base package, ie "github.com/brianm/goskel"
# this autodetection might work for you, or might now
PACKAGE := $(shell git remote -v | grep push | grep origin | awk '{print $2}' | cut -d '@' -f 2 | tr ':' '/' | cut -f 1,2 -d '.')

# where dependencies, etc, should check out to
WORKSPACE=$(PWD)/WORKSPACE

# change this to build your thing appropriately
# in this case it is building the "hello" binary
# yours will probably be different
build: deps
	GOPATH=${WORKSPACE} go install $(PACKAGE)/hello

deps: workspace
	GOPATH=$(WORKSPACE) go get $(PACKAGE)

workspace:
	@mkdir -p $(WORKSPACE)/src/$(PACKAGE)
	@rm -r $(WORKSPACE)/src/$(PACKAGE) # remove last dir for symlink
	@ln -s $(PWD) $(WORKSPACE)/src/$(PACKAGE)

clean:
	rm -rf $(WORKSPACE)/pkg/*
	rm -rf $(WORKSPACE)/bin/*

clean-workspace: clean
	rm -rf $(WORKSPACE)

env: workspace
	@echo "export GOPATH=$(WORKSPACE)"
	@echo "export PATH=$(WORKSPACE)/bin:$(PATH)"
	@echo "export WORKING_ON=$(PACKAGE)"
	@echo "export project=$(WORKSPACE)/src/$(PACKAGE)"
	@echo "export root=$(PWD)"

docserver: deps
	@echo "starting docserver on http://localhost:5050"; 
	$(eval TMPD := $(shell mktemp -d /tmp/godoc.XXX))
	$(eval DST="$(TMPD)/src/$(PACKAGE)")
	@mkdir -p $(DST)
	@rm -r $(DST) # remove final dir
	@cp -RH $(WORKSPACE)/src/* $(TMPD)/src/ #don't follow symlinks
	@rm $(DST) # remove the symlink
	@mkdir $(DST) # make dir to hold our stuff
	@cp -r *.go $(DST)
	@find . -d 1 -type d -not -name WORKSPACE -not -name .git \
		-exec cp -r {} $(DST) \;
	@echo "#!/bin/bash" > $(TMPD)/run
	@echo "trap 'rm -rf $(TMPD); exit' SIGHUP SIGINT SIGTERM" >> $(TMPD)/run
	@echo "GOPATH=$(TMPD) godoc -http=:5050" >> $(TMPD)/run
	@bash $(TMPD)/run

check-sanity:
	@echo "PACKAGE=$(PACKAGE)"
