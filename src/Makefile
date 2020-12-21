all: dynamic
	
dynamic:
	g++ -fPIC -c mylib.cpp
	g++ -shared -o libmylib.so mylib.o
	mv -f libmylib.so ../lib
	rm -f mylib.o

static:
	g++ -c mylib.cpp
	ar rc libmylib.a mylib.o
	ranlib libmylib.a
	mv -f libmylib.a ../lib
	rm -f mylib.o