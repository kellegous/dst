ALL: main

main: main.cc
	g++ -Wall -std=c++17 -o $@ $<

clean:
	rm -f main