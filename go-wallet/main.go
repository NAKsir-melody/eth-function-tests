package main

import (
	"context"
//	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"time"
//	"math/big"
	"os"

	"encoding/binary"
	"log"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

var (
	images   *glutil.Images
	fps      *debug.FPS
	program  gl.Program
	position gl.Attrib
	offset   gl.Uniform
	color    gl.Uniform
	buf      gl.Buffer

	green  float32
	touchX float32
	touchY float32
)

func main() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
		log.SetOutput(os.Stdout)
		//log.SetOutput(os.Stderr)
		log.Printf("===========")
		ScanBlock()
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop(glctx)
					glctx = nil
				}
			case size.Event:
				sz = e
				touchX = float32(sz.WidthPx / 2)
				touchY = float32(sz.HeightPx / 2)
			case paint.Event:
				if glctx == nil || e.External {
					// As we are actively painting as fast as
					// we can (usually 60 FPS), skip any paint
					// events sent by the system.
					continue
				}

				onPaint(glctx, sz)
				a.Publish()
				// Drive the animation by preparing to paint the next frame
				// after this one is shown.
				a.Send(paint.Event{})
			case touch.Event:
				touchX = e.X
				touchY = e.Y
				log.Printf("===========touch %d  %d", e.X, e.Y)
			}
		}
	})
}
// major
func ScanBlock() {
	ctx := context.Background();
	//client, err := ethclient.Dial( "https://mainnet.infura.io/v3/d023760061574ff7b403904d70ee0d55")
	//client, err := ethclient.DialContext(ctx, "https://mainnet.infura.io/v3/d023760061574ff7b403904d70ee0d55")
	client, err := ethclient.DialContext(ctx, "https://ropsten.infura.io/v3/d023760061574ff7b403904d70ee0d55")
	if err != nil {
		fmt.Println("rpc conn error")
		log.Printf("rpc conn error")
		return
	}
	defer client.Close()
/*
/// does not works with infura
	query := ethereum.FilterQuery {
		Addresses: []common.Address{},
	}
	var ch = make(chan types.Log)
	subs, err := client.SubscribeFilterLogs(ctx, query, ch)
	//subscribe := make(chan *types.Header)
	//subs, err := client.SubscribeNewHead(ctx, subscribe)
	if err != nil {
		fmt.Println("Subscribe fail ", err)
		return
	}
	for ; ; {
	select {
	case err = <-subs.Err():
	case _ = <-ch:
		fmt.Println("Subscribe")
	}
*/
	var nBlock uint64
	nBlock = 0
	//for ; ; {
		//b_num := big.NewInt(1)
		//block, err := client.BlockByNumber(ctx,b_num)
		block, err := client.BlockByNumber(ctx,nil)
		if err != nil {
			fmt.Println("block get failed")
			log.Printf("block get failed")
			return
		}
		bn := block.NumberU64();
		if(nBlock != bn) {
			file, _ := os.Create("/sdcard/scan.txt")
			defer file.Close()
			nBlock = bn;
			fmt.Println(nBlock)
			//fmt.Println(block.Time())
			//fmt.Println(block.Difficulty())
			//block_hash := block.Hash()
			block_body := block.Body()
			//eth_unit := big.NewInt(1000000000000000000)

			for no, tx := range block_body.Transactions {
				fmt.Println("==============")
				log.Printf("==============")
				//test write data
				addr := []byte("9f829cadc3424feb1fd7945778a6478d8e19667ada4b0896a58be95a6ae0b036")
				file.Write(addr)
				fmt.Println(no ,tx.Hash().Hex())
				from, err := client.TransactionSender(ctx, tx, block.Hash(),uint(no))
				if(err == nil){
					fmt.Println("from: ",from.Hex())
				}
				fmt.Println("to: ",tx.To().Hex())
				result := tx.Value()
				fmt.Println("value: ",result)
				time.Sleep(50 * time.Millisecond)
			}
			/*
			//Create account - needs test on testnet
			privKey := make([]byte, 32)
			for i := 0; i < 32; i++ {
				privKey[i] = byte(rand.Intn(256))
			}
			priv := convertToPrivateKey(seedPrivKey)
			address := crypto.PubkeyToAddress(priv.PublicKey)
			nonce := client.NonceAt(address,nil)

			//SendTransaciton - needs test on testnet
			var rawTx *types.Transaction
			rawTx = types.NewTransaction(nonce, address, value, gasLimit, gasPrice, input)
			signedTx, err := types.SignTx(rawTx, types.HomesteadSigner{}, privatekey)
			client.SendTransaciton(ctx,signedTx)
			*/

		}
		time.Sleep(1000 * time.Millisecond)
	//}

}

func onStart(glctx gl.Context) {
	var err error
	program, err = glutil.CreateProgram(glctx, vertexShader, fragmentShader)
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}

	buf = glctx.CreateBuffer()
	glctx.BindBuffer(gl.ARRAY_BUFFER, buf)
	glctx.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)

	position = glctx.GetAttribLocation(program, "position")
	color = glctx.GetUniformLocation(program, "color")
	offset = glctx.GetUniformLocation(program, "offset")

	images = glutil.NewImages(glctx)
	fps = debug.NewFPS(images)
}

func onStop(glctx gl.Context) {
	glctx.DeleteProgram(program)
	glctx.DeleteBuffer(buf)
	fps.Release()
	images.Release()
}

func onPaint(glctx gl.Context, sz size.Event) {
	glctx.ClearColor(1, 0, 0, 1)
	glctx.Clear(gl.COLOR_BUFFER_BIT)

	glctx.UseProgram(program)

	green += 0.01
	if green > 1 {
		green = 0
	}
	glctx.Uniform4f(color, 0, green, 0, 1)

	glctx.Uniform2f(offset, touchX/float32(sz.WidthPx), touchY/float32(sz.HeightPx))

	glctx.BindBuffer(gl.ARRAY_BUFFER, buf)
	glctx.EnableVertexAttribArray(position)
	glctx.VertexAttribPointer(position, coordsPerVertex, gl.FLOAT, false, 0, 0)
	glctx.DrawArrays(gl.TRIANGLES, 0, vertexCount)
	glctx.DisableVertexAttribArray(position)

	fps.Draw(sz)
}

var triangleData = f32.Bytes(binary.LittleEndian,
	0.0, 0.4, 0.0, // top left
	0.0, 0.0, 0.0, // bottom left
	0.4, 0.0, 0.0, // bottom right
)

const (
	coordsPerVertex = 3
	vertexCount     = 3
)

const vertexShader = `#version 100
uniform vec2 offset;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position + offset4;
}`

const fragmentShader = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`
