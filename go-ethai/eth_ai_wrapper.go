package main

import (
	"os/exec"
	//tf "github.com/tensorflow/tensorflow/tensorflow/go"
	//"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"fmt"
)

func main() {
	cmd := exec.Command("python", "/home/naksir/WORK/go-ethai/eth_ai_inf.py", "9848288787003200085")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
/*
	fmt.Printf("main\n")
	s := op.NewScope()
	c := op.Const(s, "Hello from TensorFlow version " + tf.Version())
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session.
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	output, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(output[0].Value())
*/
}

