CC=gcc
ARCHIVE=ar
ARCHIVE_INDEX=ranlib
CFLAGS=-W -Wall -ansi -pedantic -I../include
LDFLAGS=-lm
PKG_FOLDER=$(shell cat ../go.mod | grep ^module | sed -e 's/^module//' | xargs)
SHARED_LIB=../shared/libfoo.so
STATIC_LIB=../static/libfoo.a
SRC= $(wildcard *.c)
OBJ= $(SRC:.c=.o)

.PHONY: clean mrproper shared static

shared: $(SHARED_LIB)

static: $(STATIC_LIB)

$(SHARED_LIB): $(OBJ)
	$(CC) -shared -o $@ $^ $(LDFLAGS)

$(STATIC_LIB): $(OBJ)
	$(ARCHIVE) rc $@ $^ && $(ARCHIVE_INDEX) $(STATIC_LIB)

%.o: %.c
	$(CC) -fPIC -o $@ -c $< $(CFLAGS)

clean:
	rm -rf *.o