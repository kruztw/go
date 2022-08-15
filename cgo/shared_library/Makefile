all: static

static:
	cd src; make static
	go run simple.go

dynamic:
	cd src; make dynamic
	#go run simple.go
	go run -exec "env DYLD_LIBRARY_PATH=./src" simple.go # LDFLAGS: -L./src not work in MAC

clean:
	rm simple
	cd src; make clean

