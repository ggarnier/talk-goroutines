## Tracing the execution

<div style="display: inline-block; font-size: 70%">
  <div style="display: flex; align-items: center">
    <span style="width: 50px; height: 50px; background-color: #889e89"></span>
    <span style="margin-left: 20px">Runnable goroutines</span>
  </div>
  <div style="display: flex; align-items: center">
    <span style="width: 50px; height: 50px; background-color: #807f8a"></span>
    <span style="margin-left: 20px">Running goroutines</span>
  </div>
</div>

`go run main.go` (multithread)

<img src="static/trace-multithread.png" style="width:1200px; height: 80px" />

<br />

`GOMAXPROCS=1 go run main.go` (single thread)

<img src="static/trace-singlethread-nosleep.png" style="width:1200px; height: 120px" />
