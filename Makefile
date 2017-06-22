APP=hateblo2hugo
BASE_PACKAGE=github.com/stormcat24/$(APP)
SERIAL_PACKAGES= \
		 service
TARGET_SERIAL_PACKAGES=$(addprefix test-,$(SERIAL_PACKAGES))

deps-build:
		go get -u github.com/golang/dep/...
		go get github.com/golang/lint/golint

deps: deps-build
		dep ensure

deps-update: deps-build
		rm -rf ./vendor
		rm -rf Gopkg.lock
		dep ensure -update

build:
		go build -ldflags="-w -s" -o bin/$(APP) main.go

test: $(TARGET_SERIAL_PACKAGES)

$(TARGET_SERIAL_PACKAGES): test-%:
		go test $(BASE_PACKAGE)/$(*)

mock:
	go get github.com/golang/mock/mockgen
	mockgen -source service/movabletype.go -package service -destination service/movabletype_mock.go
