package main

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}
	header = append(header, 'P', 'N', 'G')
	println(png, header)
}
