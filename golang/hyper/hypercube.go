package main

import "fmt"
import "math/rand"
import "time"
import "math"
import "strconv"

type packet struct {
    srce int
    dest int
}

func make_bundle(width int) (bs []chan packet) {
    bs = make([]chan packet,width)
    for i := 0; i < width; i++ {
        bs[i] = make(chan packet)
    }
    return
}

func make_injector(chs []chan packet) {
    n := len(chs)
    go func() {
        for {
            s := rand.Intn(n)
            t := rand.Intn(n)
            fmt.Printf("Injection of packet {%v --> %v}.\n", s, t)
            chs[s] <- packet{s,t}
            time.Sleep(time.Duration(2)*time.Second)
        }
    } ()
}


func node(index int, in chan packet, outs [](chan packet), report chan packet) {
	for {
		p := <- in
		if p.dest == index {
			fmt.Printf("Node %v received packet from %v.\n",index,p.srce)
			report <- p

		} else {
			fmt.Printf("Node %v forwarding packet {%v --> %v}.\n",index,p.srce,p.dest)
			k := int(math.Log2(float64(len(outs))))
			i := 0
			for i < k {
				s := uint(i)
				z := index & (1 << s)
				x := p.dest & (1 << s)
				if (z != x){
					y := index ^ (1 << s)
					outs[y] <- p
					break
				}
				i = i + 1
			}
		}
	}
}

func make_hypcube(chs []chan packet, rep chan packet) {
	n := len(chs)
	i := 0
	for i < n {
		go node(i, chs[i], chs, rep)
		i = i + 1
	}
}

func main() {
	var s string = ""
	var k int = 0

	fmt.Println("Enter hypercube dimensions: ")
	fmt.Scanln(&s)
	k, _ = strconv.Atoi(s) 
	n := int(math.Pow(2, float64(k)))

	chs := make_bundle(n)
	rep := make(chan packet)
	make_injector(chs)
	make_hypcube(chs,rep)
	for x := range rep {
		fmt.Printf("Packet {%v --> %v} routed.\n", x.srce, x.dest)
	}
}
