package main

import (
//	"flag"
//	"json"
    "os"
    "./file"
    "fmt"
    "./network"
    "container/vector"
)

// var network_file = flag.



func main() {
	f, err := file.Open("./network")
    if f == nil {
        fmt.Printf("cant't open file; err=%s\n", err.String())
        os.Exit(1)
    }
    fmt.Printf("opened network file\n");
    ann :=  network.CreateNetwork(2, 1)
    fmt.Printf("\twe have a network %v\n", ann)

    fmt.Printf("\tcreating neuron1\n")
    neuron1 := ann.CreateNeuron("neuron1", 1, 1.0)
    ann.AddInput(0, neuron1, 1)

    fmt.Printf("\tcreating neuron2\n")
    neuron2 := ann.CreateNeuron("neuron2", 1, 1.0)
    ann.AddInput(1, neuron2, 1)
    
    fmt.Printf("\tcreating neuron3\n")
    neuron3 := ann.CreateNeuron("neuron3", 2, 1.0)
    neuron3.AddInput(neuron1, 1)
    neuron3.AddInput(neuron2, -1)

    fmt.Printf("\tcreating neuron4\n")
    neuron4 := ann.CreateNeuron("neuron4", 2, 1.0)
    neuron4.AddInput(neuron1, -1)
    neuron4.AddInput(neuron2, 1)

    fmt.Printf("\tcreating neuron5\n")
    neuron5 := ann.CreateNeuron("neuron5", 2, 1.0)
    neuron5.AddInput(neuron3, 1)
    neuron5.AddInput(neuron4, 1)

    output1 := ann.Output(0)
    neuron5.AddOutput(output1)

    pattern1 := network.MakePattern(0,0)
    var (
        output *vector.Vector
        e os.Error
    )
    output, e = ann.RunPattern(pattern1)
    if(e == os.EINVAL) {
        os.Exit(1)
    }
    for i := 0; i < output.Len(); i++ {
        fmt.Printf("output%v = %v\n", i+1, output.At(i))
    }

    pattern2 := network.MakePattern(0,1)
    output, e = ann.RunPattern(pattern2)
    if(e == os.EINVAL) {
        os.Exit(1)
    }
    for i := 0; i < output.Len(); i++ {
        fmt.Printf("output%v = %v\n", i+1, output.At(i))
    }

    pattern3 := network.MakePattern(1,0)
    output, e = ann.RunPattern(pattern3)
    if(e == os.EINVAL) {
        os.Exit(1)
    }
    for i := 0; i < output.Len(); i++ {
        fmt.Printf("output%v = %v\n", i+1, output.At(i))
    }

    pattern4 := network.MakePattern(1,1)
    output, e = ann.RunPattern(pattern4)
    if(e == os.EINVAL) {
        os.Exit(1)
    }
    for i := 0; i < output.Len(); i++ {
        fmt.Printf("output%v = %v\n", i+1, output.At(i))
    }
}
