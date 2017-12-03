package dynaport

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

const (
	lowPort   = 10000
	maxPorts  = 65535
	blockSize = 1024
	maxBlocks = 16
	attempts  = 10
)

var (
	port      int
	firstPort int
	once      sync.Once
	mu        sync.Mutex
	lnLock    sync.Mutex
	lockLn    net.Listener
)

// MustGet returns n ports that are free to use, panicing if it can't succeed.
func MustGet(n int) []int {
	ports, err := Get(n)
	if err != nil {
		panic(err)
	}
	return ports
}

// Get returns n ports that are free to use.
func Get(n int) (ports []int, err error) {
	mu.Lock()
	defer mu.Unlock()

	if n > blockSize-1 {
		return nil, fmt.Errorf("dynaportt: block size is too small for ports requested")
	}

	once.Do(initialize)

	for len(ports) < n {
		port++

		if port < firstPort+1 || port >= firstPort+blockSize {
			port = firstPort + 1
		}

		ln, err := listen(port)
		if err != nil {
			continue
		}
		ln.Close()

		ports = append(ports, port)
	}

	return ports, nil
}

func initialize() {
	if lowPort+maxBlocks*blockSize > maxPorts {
		panic("dynaport: block size is too big or too many blocks requested")
	}
	rand.Seed(time.Now().UnixNano())
	var err error
	for i := 0; i < attempts; i++ {
		block := int(rand.Int31n(int32(maxBlocks)))
		firstPort = lowPort + block*blockSize
		lockLn, err = listen(firstPort)
		if err != nil {
			continue
		}
		return
	}
	panic("dynaport: can't allocated port block")
}

func listen(port int) (*net.TCPListener, error) {
	return net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: port})
}
