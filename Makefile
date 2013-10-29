# set this to match the base package, ie "github.com/brianm/goskel"
# this autodetection might work for you, or might not
PACKAGE := $(shell git remote -v | grep origin | grep fetch | ruby -pe '$$_ =~ %r{https://(.+)\.git}; $$_ = $$1')

WORKSPACE=$(PWD)/.workspace

# Change this to build your thing appropriately
# in this case it is building the "hello" binary
# yours will probably be different
build: workspace
	GOPATH=$(WORKSPACE) go get $(PACKAGE)
	cp $(WORKSPACE)/bin/goskel .

# Run tests in root, and non {WORKSPACE .git} subdirs
# of the root
test: build
	@GOPATH=$(WORKSPACE) go test $(PACKAGE)
	@GOPATH=$(WORKSPACE) find . -d 1 -type d \
									 -not -name $(shell basename $(WORKSPACE)) \
									 -not -name .git \
		-exec go test $(PACKAGE)/{}  \;

$(WORKSPACE):
	$(eval WORK_BUILD := $(shell mktemp -d /tmp/goskel.XXX))
	mkdir -p $(WORK_BUILD)/src/$(PACKAGE)
	rm -r $(WORK_BUILD)/src/$(PACKAGE)
	cp -pr . $(WORK_BUILD)/src/$(PACKAGE)
	mv $(WORK_BUILD) $(WORKSPACE)
	GOPATH=$(WORKSPACE) go get $(PACKAGE)

# Build the Go workspace and symlink this project into
# it at the correct place.
workspace: $(WORKSPACE)

# Wipes out build artifacts
clean:
	rm -rf $(WORKSPACE)
	rm -f goskel

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
