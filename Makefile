# set this to match the base package, ie "github.com/brianm/goskel"
# this autodetection might work for you, or might not
PACKAGE := $(shell git remote -v | grep push | grep origin \
			 | awk '{print $2}' | cut -d '@' -f 2 | tr ':' '/' \
			 | cut -f 1,2 -d '.')

TMPDIR := $(shell mktemp -d /tmp/go_env.XXX)

# The Go workspace directory
WORKSPACE=$(PWD)/WORKSPACE

# Change this to build your thing appropriately
# in this case it is building the "hello" binary
# yours will probably be different
build: test
	GOPATH=$(WORKSPACE) go install $(PACKAGE)/hello
	@echo "output: $(WORKSPACE)/bin/hello"	

# Run tests in root, and non {WORKSPACE .git} subdirs
# of the root
test: deps
	@GOPATH=$(WORKSPACE) go test $(PACKAGE)
	@GOPATH=$(WORKSPACE) find . -d 1 -type d \
									 -not -name WORKSPACE \
									 -not -name .git \
		-exec go test $(PACKAGE)/{}  \;

# Run "go fmt" on likely packages
fmt: deps
	@GOPATH=$(WORKSPACE) go fmt $(PACKAGE)
	@GOPATH=$(WORKSPACE) find . -d 1 -type d \
									 -not -name WORKSPACE \
									 -not -name .git \
		-exec go fmt $(PACKAGE)/{}  \;

# Fetch dependencies
deps: workspace
	GOPATH=$(WORKSPACE) go get $(PACKAGE)

# Build the Go workspace and symlink this project into
# it at the correct place.
workspace:
	@mkdir -p $(WORKSPACE)/src/$(PACKAGE)
	@rm -r $(WORKSPACE)/src/$(PACKAGE) # remove last dir for symlink
	@ln -s $(PWD) $(WORKSPACE)/src/$(PACKAGE)

# Wipes out build artifacts
clean:
	rm -rf $(WORKSPACE)

# Wipes out the workspace
clean-workspace: clean
	rm -rf $(WORKSPACE)

# Display useful env vars which can be set to enter
# the workspace
env: workspace
	@echo "export GOPATH=$(WORKSPACE)"
	@echo "export PATH=$(WORKSPACE)/bin:$(PATH)"
	@echo "export WORKING_ON=$(PACKAGE)"
	@echo "export project=$(WORKSPACE)/src/$(PACKAGE)"
	@echo "export root=$(PWD)"

# This is, sadly, utterly gross. Godoc will not follow
# symlinks, and we use a symlink for the project in the
# workspace. In order to work around this, we set up a
# copy of the workspace in a temp dir, and run godoc 
# against that workspace.
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

# Convenience to make sure PACKAGE is being picked up correctly
check-sanity:
	@echo "PACKAGE=$(PACKAGE)"

newspace:
	find . -type d ! -path ./$(shell basename $(WORKSPACE))\* -a ! -path ./.git\* -exec mkdir -p "$(WORKSPACE)/src/$(PACKAGE)/{}" \;
	find . -type f ! -path ./$(shell basename $(WORKSPACE))\* -a ! -path ./.git/\* -exec ln {} "$(WORKSPACE)/src/$(PACKAGE)/{}" \;
	ln -s $(PWD)/.git $(WORKSPACE)/src/$(PACKAGE)/.git
