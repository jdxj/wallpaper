fileName=wallpaper.out

build:
	go build -o $(fileName) *.go
clean:
	rm -rvf $(fileName)
