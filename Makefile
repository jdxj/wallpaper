fileName=wallpaper.out

build: clean
	go build -o $(fileName) *.go
clean:
	rm -rvf $(fileName)
