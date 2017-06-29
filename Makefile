BASE_PACKAGE=github.com/openfresh/plasma-mackerel-plugin
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

build:
		go build -ldflags="-w -s" -o bin/mackerel-plugin-plasma main.go

test: $(TARGET_SERIAL_PACKAGES)

$(TARGET_SERIAL_PACKAGES): test-%:
		go test $(BASE_PACKAGE)/$(*)
