package main

import (
  // "bufio"
  // "github.com/go-vgo/robotgo"
  "image"
  "encoding/binary"
  "time"
  "fmt"
  "net"
  "os"
  "image/jpeg"
  // "github.com/vcaesar/imgo"

  // "strings"
)


var packetsRec int  = 0
var bytesRec int  = 0

var ImageBytes = make([]byte, 1000000)

func PacketsPerTrack() {
  start := time.Now()
  t := time.Now()

  for {
    start = time.Now()
    elapsed := float64(start.Sub(t)) / float64(time.Millisecond)
    if (elapsed > 1000 / 10) {
      fmt.Println(elapsed, " === ", packetsRec, " byres === ", bytesRec)
      packetsRec = 0
      bytesRec = 0
      t = start
      // robotgo.SaveImg(ImageBytes, "test2.png")
    }
  }
}

func ReadUdp(c *net.UDPConn) {
  img1 := image.NewRGBA(image.Rect(0, 0, 500, 500))
  img1.Pix = make([]uint8, 2000 * 500)
  img1.Stride = 2000



  for {
    buffer := make([]uint8, 2002)
    n, _, err := c.ReadFromUDP(buffer)
    if err != nil {
      fmt.Println(err)
      return
    }
    // if (n == 54) {
      // copy(ImageBytes[:54], buffer[0:n])
    // } else {
    newData := []uint8 {buffer[n - 2], buffer[n - 1]}
    i := binary.BigEndian.Uint16(newData)
    fmt.Println(buffer[n - 1])
    copy(img1.Pix[(int(i) * 2000):((int(i) + 1) * 2000)], buffer[0:n -2])


    // }
    bytesRec += n
    packetsRec++
    // uint16(newData)
    // fmt.Println(binary.BigEndian.Uint16(newData))
    // fmt.Printf("Reply: %s\n", string(buffer[0:n]))
    if int(i) == 499 {
      break
    }
  }
    f, _ := os.Create("outimage.jpg")
    defer f.Close()
    opt := jpeg.Options{
      Quality: 90,
    }
    jpeg.Encode(f, img1, &opt)
  // imgo.SaveToPNG("test_IMG.png", img1)
}

func main() {
  arguments := os.Args
  if len(arguments) == 1 {
    fmt.Println("Please provide a host:port string")
    return
  }
  CONNECT := arguments[1]

  s, err := net.ResolveUDPAddr("udp4", CONNECT)
  c, err := net.DialUDP("udp4", nil, s)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
  defer c.Close()
  go ReadUdp(c)
  arr := make([]byte, 1024, 1024)
  c.Write(arr)
  PacketsPerTrack()

  // for {
  //   // reader := bufio.NewReader(os.Stdin)
  //   // fmt.Print(">> ")
  //   // text, _ := reader.ReadString('\n')
  //   // data := []byte("stress test")
  //   if strings.TrimSpace(string(data)) == "STOP" {
  //     fmt.Println("Exiting UDP client!")
  //     return
  //   }
  //
  //   if err != nil {
  //     fmt.Println(err)
  //     return
  //   }
  // }
}


