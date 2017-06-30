APP=plasma-mackerel-plugin
MAIN_FILE=main.go
BASE_PACKAGE=github.com/openfresh/$(APP)
SERIAL_PACKAGES= \
		 metrics
TARGET_SERIAL_PACKAGES=$(addprefix test-,$(SERIAL_PACKAGES))

deps-build:
		go get -u github.com/golang/dep/...

deps: deps-build
		dep ensure

deps-update: deps-build
		rm -rf ./vendor
		rm -rf Gopkg.lock
		dep ensure -update

define build-artifact
		GOOS=$(1) GOARCH=$(2) go build -o artifacts/$(APP)
		cd artifacts && tar cvzf $(APP)_$(1)_$(2).tar.gz $(APP)
		rm ./artifacts/$(APP)
		@echo [INFO]build success: $(1)_$(2)
endef

build-all:
		$(call build-artifact,linux,386)
		$(call build-artifact,linux,amd64)
		$(call build-artifact,linux,arm)
		$(call build-artifact,linux,arm64)
		$(call build-artifact,darwin,amd64)

build:
		go build -o bin/$(APP) $(MAIN_FILE)

test: $(TARGET_SERIAL_PACKAGES)

$(TARGET_SERIAL_PACKAGES): test-%:
		go test $(BASE_PACKAGE)/$(*)
