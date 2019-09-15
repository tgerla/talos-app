UI = ./ui

all: build

build:
	cd $(UI) && npm run $@
	go run .

serve:
	cd $(UI) && npm run $@

lint:
	cd $(UI) && npm run $@

install:
	cd $(UI) && npm $@


clean:
	-rm -rf .gotron
	-rm -rf $(UI)/dist
	-rm -rf $(UI)/node_modules
