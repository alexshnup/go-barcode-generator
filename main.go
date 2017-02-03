package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"

	"image/color"

	"image/draw"

	"golang.org/x/image/bmp"

	"syscall"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"
)

func main() {

	switch string(os.Args[2]) {
	case "qr":
		// base64 := "IAV19ysYSl0HUuG5QiCDvdHkowqdGXb0HbaUAWUzHw=="
		base64 := os.Args[3]
		log.Println("Original data:", base64)

		code1pixel, err := qr.Encode(base64, qr.L, qr.Unicode)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Encoded data: ", code1pixel.Content())

		if base64 != code1pixel.Content() {
			log.Fatal("data differs")
		}
		log.Println("Encoded data: ", code1pixel.Content())

		if base64 != code1pixel.Content() {
			log.Fatal("data differs")
		}

		codeScalled, err := barcode.Scale(code1pixel, 300, 200)
		if err != nil {
			log.Fatal(err)
		}
		drtest(codeScalled)

	case "ean":
		// code, err := ean.Encode("123456789012")
		code1pixel, err := ean.Encode(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Encoded data: ", code1pixel.Content())

		// if base64 != code.Content() {
		// 	log.Fatal("data differs")
		// }

		codeScalled, err := barcode.Scale(code1pixel, 300, 200)
		if err != nil {
			log.Fatal(err)
		}
		drtest(codeScalled)

	}
}

// func writePng(filename string, img image.Image) {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = png.Encode(file, img)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	file.Close()
// }

func drtest(imgSrc image.Image) {

	// pngImgFile, err := os.Open("./test.png")
	//
	// if err != nil {
	// 	fmt.Println("PNG-file.png file not found!")
	// 	os.Exit(1)
	// }
	//
	// defer pngImgFile.Close()
	//
	// // create image from PNG file
	// imgSrc, err := png.Decode(pngImgFile)
	//
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// create a new Image with the same dimension of PNG image
	newImg := image.NewGray(image.Rect(0, 0, 600, 800))

	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds().Add(image.Point{150, 300}), imgSrc, imgSrc.Bounds().Min, draw.Over)
	// draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min.Add(image.Point{85, 85}), draw.Over)

	// // create new out JPEG file
	// // jpgImgFile, err := os.Create("./test.jpg")
	// jpgImgFile, err := os.Create("/dev/sde")
	//
	// if err != nil {
	// 	fmt.Println("Cannot create JPEG-file.jpg !")
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//
	// defer jpgImgFile.Close()
	//
	// var opt jpeg.Options
	// opt.Quality = 80
	//
	// // convert newImage to JPEG encoded byte and save to jpgImgFile
	// // with quality = 80
	// // err = jpeg.Encode(jpgImgFile, newImg, &opt)
	// err = jpeg.Encode(jpgImgFile, newImg, &opt)
	//
	// //err = jpeg.Encode(jpgImgFile, newImg, nil) -- use nil if ignore quality options
	//
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// ImgOut, err := os.Create("/dev/sde")
	//
	// if err != nil {
	// 	fmt.Println("Cannot create out !")
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer ImgOut.Close()

	buf := new(bytes.Buffer)

	err := bmp.Encode(buf, newImg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	send_s3 := buf.Bytes()

	fmt.Println("OK")

	// title := "title"
	// message := "body 1111111111111111111"
	// mailCommand := exec.Command("mail", "-s", title, "root@localhost")
	// stdin, err := mailCommand.StdinPipe()
	// if err != nil {
	//   fmt.Fprintf(os.Stderr, "StdinPipe failed to perform: %s (Command: %s, Arguments: %s)", err, mailCommand.Path, mailCommand.Args)
	//   return
	// }
	// stdin.Write([]byte(message))
	// stdin.Close()
	// mailCommand.Output()
	// fmt.Println(mailCommand.ProcessState.String())

	// buffer := make([]byte, 480000)

	devout(send_s3[1078:])

}

func devout(buffer []byte) {
	// disk := "/dev/sde"
	disk := os.Args[1]
	var fd, numread int
	var err error

	fd, err = syscall.Open(disk, syscall.O_RDWR, 0777)

	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}

	// buffer := make([]byte, 480000)

	// //READ
	// numread, err = syscall.Read(fd, buffer)
	//
	// if err != nil {
	// 	fmt.Print(err.Error(), "\n")
	// }
	//
	// fmt.Printf("Numbytes read: %d\n", numread)
	// fmt.Printf("Buffer: %b\n", buffer[])

	//WRITE
	numread, err = syscall.Write(fd, buffer)

	if err != nil {
		fmt.Print(err.Error(), "\n")
	}

	fmt.Printf("Numbytes write: %d\n", numread)
	// fmt.Printf("Buffer: %x\n", buffer[:1000])

	err = syscall.Close(fd)

	if err != nil {
		fmt.Print(err.Error(), "\n")
	}
}
