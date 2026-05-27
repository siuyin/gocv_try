# Learning GoCV / OpenCV

## Development environment
```
docker run -it --name gocv -h gocv -v $HOME:/h -u $UID --network host \
  --device /dev/video0:/dev/video0 -v /tmp/.X11-unix:/tmp/.X11-unix \
  -e DISPLAY=$DISPLAY -w /h gocv/opencv bash

docker exec -it -u root gocv bash
within bash session:
    adduser --uid 1000 --disabled-password siuyin
    addgroup siuyin video
```

To restart exited container:
```
docker start -ai gocv
```
