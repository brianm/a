# set this to match the base package, ie "github.com/brianm/goskel"
# this autodetection might work for you, or might not
PACKAGE := $(shell git remote -v | grep push | grep origin \
			 | awk '{print $2}' | cut -d '@' -f 2 | tr ':' '/' \
			 | cut -f 1,2 -d '.')

WORK_BUILD := $(shell mktemp -d /tmp/go_env.XXX)

WORKSPACE=$(PWD)/WORKSPACE

# Change this to build your thing appropriately
# in this case it is building the "hello" binary
# yours will probably be different
build: workspace
	GOPATH=$(WORKSPACE) go get $(PACKAGE)/hello

# Run tests in root, and non {WORKSPACE .git} subdirs
# of the root
test: build
	@GOPATH=$(WORKSPACE) go test $(PACKAGE)
	@GOPATH=$(WORKSPACE) find . -d 1 -type d \
									 -not -name WORKSPACE \
									 -not -name .git \
		-exec go test $(PACKAGE)/{}  \;

# Run "go fmt" on likely packages
fmt: workspace
	@GOPATH=$(WORKSPACE) go fmt $(PACKAGE)
	@GOPATH=$(WORKSPACE) find . -d 1 -type d \
									 -not -name WORKSPACE \
									 -not -name .git \
		-exec go fmt $(PACKAGE)/{}  \;

WORKSPACE:
	mkdir -p $(WORK_BUILD)/src/$(PACKAGE)
	@echo "copying!"
	rm -r $(WORK_BUILD)/src/$(PACKAGE)
	cp -pr . $(WORK_BUILD)/src/$(PACKAGE)
	@echo "moving!"
	mv $(WORK_BUILD) $(WORKSPACE)
	@echo "getting"
	GOPATH=$(WORKSPACE) go get $(PACKAGE)

# Build the Go workspace and symlink this project into
# it at the correct place.
workspace: WORKSPACE

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

docserver: workspace
	GOPATH=$(WORKSPACE) godoc -http=:5050

# Convenience to make sure PACKAGE is being picked up correctly
check-sanity:
	@echo "PACKAGE=$(PACKAGE)"
