package main

import (
  // "encoding/binary"
  "fmt"
  "math/rand"
  "net"
  "os"
  "github.com/go-vgo/robotgo"
  "image/jpeg"
  // "reflect"
  // "strconv"
  // "strings"
  "image"
  "time"
  "unsafe"
  "bytes"
)

func random(min, max int) int {
  return rand.Intn(max-min) + min
}

var packetsRec int  = 0
// var  ImageBytes = []byte{}

func PacketsPerTrack() {
}

func val(p *uint8, n int) uint8 {
  addr := uintptr(unsafe.Pointer(p))
  addr += uintptr(n)
  p1 := (*uint8)(unsafe.Pointer(addr))
  return *p1
}



// serveWs handles websocket requests from the peer.
func copyToVUint8A(dst []uint8, src *uint8) {
	for i := 0; i < len(dst)-4; i += 4 {
		dst[i] = val(src, i+2)
		dst[i+1] = val(src, i+1)
		dst[i+2] = val(src, i)
		dst[i+3] = val(src, i+3)
	}
}

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("Please provide a port number!")
    return
  }
  PORT := ":" + arguments[1]

  s, err := net.ResolveUDPAddr("udp4", PORT)
  if err != nil {
    fmt.Println(err)
    return
  }

  connection, err := net.ListenUDP("udp4", s)
  if err != nil {
    fmt.Println(err)
    return
  }

  defer connection.Close()
  // buffer := make([]byte, 1024)
  // rand.Seed(time.Now().Unix())

  // _, addr, err := connection.ReadFromUDP(buffer)

  start := time.Now()
  t := time.Now()

  for {
    start = time.Now()
    elapsed := float64(start.Sub(t)) / float64(time.Millisecond)
    if (elapsed > 1000 / 60) {
      // for i := 0; i < 10; i++ {
        bitmap := robotgo.CaptureScreen(0, 0, 1000, 1000)
        fmt.Println(elapsed)
        // ImageBytes := robotgo.ToBitmapBytes(bitmap)
        // ft :=robotgo.ToImage(bitmap)
        // fmt.Println(ft.ColorModel())
        bmp1 := robotgo.ToBitmap(bitmap)
        img1 := image.NewRGBA(image.Rect(0, 0, bmp1.Width, bmp1.Height))
        img1.Pix = make([]uint8, bmp1.Bytewidth*bmp1.Height)

        copyToVUint8A(img1.Pix, bmp1.ImgBuf)
        img1.Stride = bmp1.Bytewidth

        opt := jpeg.Options{
          Quality: 10,
        }

        f := new(bytes.Buffer)

        jpeg.Encode(f, img1, &opt)
      // }
      fmt.Println(len(img1.Pix), len(f.Bytes()))

      // _ = img1
      // fmt.Println(img1.Stride)
      // fmt.Println(bmp1.Height, bmp1.Bytewidth)

      // fmt.Println(len(ImageBytes))
      // for i := 0; i < 54 ; i++ {
      //   fmt.Println(ImageBytes[i])
      // }

      // fmt.Println(elapsed, " === ", packetsRec)
      // for i := 0; i < 100; i++ {
      //   data := img1.Pix[(i * 2000):((i + 1) * 2000)]
      //   var h, l uint8 = byte(uint16(i)>>8), byte(uint16(i)&0xff)
      //   b := make([]uint8, 2)
      //   b[0] = h
      //   b[1] = l
      //   // fmt.Println(b)
      //
      //
      //   _, err = connection.WriteToUDP(append(data, b...), addr)
      //   if err != nil {
      //     fmt.Println(err)
      //     return
      //   }
      // }
      // break
      //
      // f, _ := os.Create("bob.jpg")
      // defer f.Close()
      // opt := jpeg.Options{
      //   Quality: 90,
      // }
      //
      // jpeg.Encode(f, img1, &opt)
      t = start
    }
  }
}


