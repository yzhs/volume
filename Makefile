PREFIX=/usr/local/bin

all: volume

volume: Volume.hs
	ghc -o volume Volume.hs

install: volume
	install -s volume $(PREFIX)

clean:
	-rm Volume.hi Volume.o volume
