package main

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight float32 = 800, 600

type gameState int

const (
	start gameState = iota
	play
	pause
	gameover
)

const ballStartSpeed float32 = 300
const paddleStartSpeed float32 = 300

var state = start

var nums = [][]byte{
	{
		1, 1, 1,
		1, 0, 1,
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 0,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		0, 1, 1,
		0, 0, 1,
		1, 1, 1,
	},
}

var pingpong = []byte{
	0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1,
	1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
	1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1,
	1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1,
}

var gameOver = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0,
	0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0,
	1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

var pressSpaceToPlay = []byte{
	1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1,
	1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1,
	1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0,
	1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0,
	1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0,
}

func showMessage(message []byte, size int, pos pos, pixels []byte) {
	startX := int(pos.x) - len(message)*size/10 - size*3/2
	startY := int(pos.y) - 2 - size*5/2
	for i, v := range message {
		if v == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(x, y, color{255, 255, 255}, pixels)
				}
			}
		}
		startX += size
		if (i+1)%(len(message)/5) == 0 {
			startY += size
			startX -= size * len(message) / 5
		}
	}
}

func drawNumber(pos pos, color color, size int, num int, pixels []byte) {
	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size*5)/2

	for i, v := range nums[num] {
		if v == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(x, y, color, pixels)
				}
			}
		}
		startX += size
		if (i+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

type ball struct {
	pos
	radius float32
	xv     float32
	yv     float32
	color  color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x+x), int(ball.y+y), color{255, 255, 255}, pixels)
			}
		}
	}
}

func (ball *ball) update(leftPaddle, rightPaddle *paddle, elapsedTime float32) {
	ball.x += ball.xv * elapsedTime
	ball.y += ball.yv * elapsedTime

	//handle collisions
	if ball.y-ball.radius < 0 || ball.y+ball.radius > winHeight {
		ball.yv = -ball.yv
	}

	if ball.x < 0 {
		rightPaddle.score++
		ball.pos = getCenter()
		if rightPaddle.score == 3 {
			state = gameover
		} else {
			state = pause
		}
	} else if ball.x > winWidth {
		leftPaddle.score++
		ball.pos = getCenter()
		if leftPaddle.score == 3 {
			state = gameover
		} else {
			state = pause
		}
	}

	if ball.x-ball.radius < leftPaddle.x+leftPaddle.w/2 {
		if ball.y-ball.radius > leftPaddle.y-leftPaddle.h/2 && ball.y+ball.radius < leftPaddle.y+leftPaddle.h/2 {
			ball.xv = -ball.xv + 100
			leftPaddle.speed += 75
			ball.yv += (ball.y - leftPaddle.y) * 12
			ball.x = leftPaddle.x + leftPaddle.w/2 + ball.radius
		}
	}

	if ball.x+ball.radius > rightPaddle.x-rightPaddle.w/2 {
		if ball.y-ball.radius > rightPaddle.y-rightPaddle.h/2 && ball.y+ball.radius < rightPaddle.y+rightPaddle.h/2 {
			ball.xv = -ball.xv - 100
			rightPaddle.speed += 50
			ball.yv += (ball.y - rightPaddle.y) * 12
			ball.x = rightPaddle.x - rightPaddle.w/2 - ball.radius
		}
	}
}

type paddle struct {
	pos
	w     float32
	h     float32
	speed float32
	score int
	color color
}

func lerp(a, b, pct float32) float32 {
	return a + pct*(a-b)
}

func (paddle *paddle) draw(pixels []byte) {
	startX := paddle.x - float32(paddle.w)/2
	startY := paddle.y - float32(paddle.h)/2
	for y := 0; y < int(paddle.h); y++ {
		for x := 0; x < int(paddle.w); x++ {
			setPixel(x+int(startX), y+int(startY), color{255, 255, 255}, pixels)
		}
	}
	numX := lerp(paddle.x, getCenter().x, 0.5)
	drawNumber(pos{numX, 35}, color{255, 255, 255}, 10, paddle.score, pixels)
}

func (paddle *paddle) update(keyState []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= paddle.speed * elapsedTime
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += paddle.speed * elapsedTime
	}
}

func (paddle *paddle) aiUpdate(ball *ball, elapsedTime float32) {
	if ball.y > paddle.y {
		paddle.y += paddle.speed * elapsedTime
	} else if ball.y < paddle.y {
		paddle.y -= paddle.speed * elapsedTime
	}
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*int(winWidth) + x) * 4
	if index <= len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

func getCenter() pos {
	return pos{winWidth / 2, winHeight / 2}
}

func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		log.Fatal(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("testing sdl2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(uint32(sdl.PIXELFORMAT_ABGR8888), sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		log.Fatal(err)
	}
	defer tex.Destroy()

	pixels := make([]byte, int(winWidth*winHeight)*4)

	player1 := paddle{pos{50, 100}, 20, 100, paddleStartSpeed, 0, color{255, 255, 255}}
	player2 := paddle{pos{float32(winWidth) - 50, 100}, 20, 100, paddleStartSpeed, 0, color{255, 255, 255}}
	ball := ball{pos{300, 300}, 20, ballStartSpeed, ballStartSpeed, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()
	var frameStart time.Time
	var elapsedTime float32

	for {

		frameStart = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		switch state {

		case play:
			player1.update(keyState, elapsedTime)
			player2.aiUpdate(&ball, elapsedTime)
			ball.update(&player1, &player2, elapsedTime)

		case pause:
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				player1.speed = paddleStartSpeed
				player2.speed = paddleStartSpeed
				ball.xv = ballStartSpeed
				ball.yv = ballStartSpeed
				state = play
			}

		case start:
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				state = pause
			}

		case gameover:
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				player1.speed = paddleStartSpeed
				player2.speed = paddleStartSpeed
				ball.xv = ballStartSpeed
				ball.yv = ballStartSpeed
				player1.score = 0
				player2.score = 0
				state = play
			}

		}

		clear(pixels)
		if state == start {
			showMessage(pingpong, 14, getCenter(), pixels)
			showMessage(pressSpaceToPlay, 6, pos{winWidth / 2, winHeight/2 + 100}, pixels)
		} else if state == gameover {
			showMessage(gameOver, 14, getCenter(), pixels)
			showMessage(pressSpaceToPlay, 6, pos{winWidth / 2, winHeight/2 + 100}, pixels)
		} else {
			if state == pause {
				showMessage(pressSpaceToPlay, 6, pos{winWidth / 2, winHeight/2 + 100}, pixels)
			}

			player1.draw(pixels)
			player2.draw(pixels)
			ball.draw(pixels)
		}

		tex.Update(nil, pixels, int(winWidth)*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())
		if elapsedTime < 0.005 {
			sdl.Delay(5 - uint32(elapsedTime/1000.0))
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}
