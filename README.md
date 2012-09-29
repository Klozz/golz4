golz4
=====

Basic lz4 bindings for C LZ4.

Just lets you call the methods on byte slices.

To build, get LZ4 from http://code.google.com/p/lz4, and install:  
`svn checkout http://lz4.googlecode.com/svn/trunk/ lz4-read-only`  
`cd lz4-read-only`  
`gcc -O3 -I. -std=c99 -Wall -W -Wundef -Wno-implicit-function-declaration lz4hc.c lz4.c bench.c -shared -o liblz4.so`  
`cp *.h /usr/local/include`  
`cp liblz4.so /usr/local/include`
