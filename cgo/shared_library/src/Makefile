# Places a static or shared (dynamic) library into ../lib

all: static
	
dynamic:
	rm -f lib* || true
	gcc -fPIC -c mylib.c
	gcc -shared -o libmylib.so mylib.o

static:
	rm -f lib* || true
	gcc -c mylib.c
	ar rc libmylib.a mylib.o
	ranlib libmylib.a

clean:
	rm -f mylib.o || true
	rm -f lib* || true
