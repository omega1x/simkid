echo "[simkid]: run simkid"
FRAME_STORAGE="$PWD/frameset"
VIDEO_SRORAGE="$PWD/videoset"
OUTPUT="$VIDEO_SRORAGE/$(date +%Y%m%d%H%M%S)_simkid-simulation.mp4"

echo "[simkid]: create PNG-frames"
./simkid
echo "[simkid]: process PNG-frames"
ffmpeg -start_number 0 -framerate 60 -i $FRAME_STORAGE/frame-%04d.png -pix_fmt yuv420p -movflags +faststart -vcodec libx264 $OUTPUT
mpv "./videoset/$(ls -Art ./videoset | tail -n 1)" > mpv.log 2>&1 & 
echo "[simkid]: all done, thanks!"
