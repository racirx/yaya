PROJ=yaya
ORG_PATH=github.com/racirx
REPO_PATH=$(ORG_PATH)/$(PROJ)
export PATH := $(PWD)/bin:$(PATH)

VERSION ?= $(shell ./scripts/git-version)

LD_FLAGS="-w -X $(REPO_PATH)/version.Version=$(VERSION)"

build: bin/yaya

bin/yaya:
	@mkdir -p bin/
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)/cmd