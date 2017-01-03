package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  "net/http"
  "time"
)

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
    return
  }
  defer file.Close()

  outFile, err := os.Create(os.Args[1] + ".out.csv")
  if err != nil {
    log.Fatal(err)
    return
  }
  defer outFile.Close()

  out := bufio.NewWriter(outFile)

  _, err = out.WriteString("url,status\n")
  if err != nil {
    log.Fatal(err)
    return
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    checkUrl(scanner.Text(), out)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
    return
  }

}

func checkUrl(url string, out *bufio.Writer) {

  retry := true
  exp_backoff := time.Duration(time.Second)
  for retry {
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(out, "%v,%v\n", url, err)
      log.Println(err)
      return
    }

    fmt.Println(url, resp.StatusCode)

    if resp.StatusCode == 429 {
      time.Sleep(exp_backoff)
      fmt.Printf("slept %v\n", exp_backoff)
      exp_backoff = exp_backoff * 2
    } else {
      fmt.Fprintf(out, "%v,%v\n", url, resp.StatusCode)
      retry = false
    }
  }

  out.Flush()

}