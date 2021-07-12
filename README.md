# TileWorldGo

TileWorld this time in GoLang

Docker:

docker build -t tileworldgo .

macOS:

docker run -ti -e DISPLAY=host.docker.internal:0 --rm --init tileworldgo

linux:

docker run -ti -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix:rw --volume="$HOME/.Xauthority:/root/.Xauthority:rw" --network=host --privileged --rm --init tileworldgo