[![CodeFactor](https://www.codefactor.io/repository/github/omega1x/simkid/badge)](https://www.codefactor.io/repository/github/omega1x/simkid)



# SIMKID: Toy motion simulator of simple physical systems for teaching kids

Teach kids physics and programming in parallel. This is home-based project for practice [Go](https://go.dev/) with my younger son.


## Details

In `simkid` the simulation space is a 1920&#10005;1080 pixel area which is updated in 60 fps rate. All coordinates are measured in pixels. Velocity and acceleration should be given in _pixel per frame_ [_p/f_] and _pixel per frame per frame_ [_p/f&sup2;_] accordingly.

The main component of `simkid` is `sim.go` - simulation logic: put here individual equations of motion for maximum 10 balls and possibly add interactions between them.

Use `simkid-ffmpeg.sh` to compile PNG-frames into final MP4-video.


## Usage

Clone the repo and run:

```
go run .
./simkid-ffmpgeg.sh
mpv "./videoset/$(ls -Art ./videoset | tail -n 1)"
```

