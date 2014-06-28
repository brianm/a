# per-project info
PACKAGE:=github.com/brianm/a
BINARY:=a

# shouldn't probably change between projects
WORKSPACE:=$(PWD)/_workspace
BRANCH:=$(shell git branch | grep '^* ' | awk '{print $$2'})

$(BINARY): workspace
	GOPATH=$(WORKSPACE) go get $(PACKAGE)
	cp $(WORKSPACE)/bin/$(BINARY) .
	@echo "Built binary at ./$(BINARY)"

workspace: $(WORKSPACE)/src/$(PACKAGE)

$(WORKSPACE)/src/$(PACKAGE):
	mkdir -p $(WORKSPACE)/src/$(PACKAGE)
	git clone . $(WORKSPACE)/src/$(PACKAGE)
	cp .git/config $(WORKSPACE)/src/$(PACKAGE)/.git/
	cd $(WORKSPACE)/src/$(PACKAGE) && git checkout $(BRANCH)

clean:
	rm -rf $(WORKSPACE) $(BINARY)

activate: workspace
	@PROJECT="$(BINARY)" GOSKEL_ACTIVATED="$(BINARY)" PACKAGE="$(PACKAGE)" /bin/bash ./project.bash activate

