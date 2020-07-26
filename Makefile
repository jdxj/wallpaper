fileName=wallpaper.out

build: clean
	go build -o $(fileName) *.go
	go install
clean:
	rm -rvf $(fileName)
