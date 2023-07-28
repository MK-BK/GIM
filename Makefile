.PHONY build:
build: 
	cd web && npm run build

txt = "hello world"
.PHONY test:
test: 
	@echo $$PWD