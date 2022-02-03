echo "[simkid] Create video file from [frameset]"

FRAME_STORAGE="$PWD/frameset"
VIDEO_SRORAGE="$PWD/videoset"
OUTPUT="$VIDEO_SRORAGE/$(date +%Y%m%d%H%M%S)_simkid-simulation.mp4"
echo "[simkid] Process frameset [$FRAME_STORAGE/*.png] to [$OUTPUT]"

ffmpeg -start_number 0 -framerate 60 -i $FRAME_STORAGE/frame-%04d.png -pix_fmt yuv420p -movflags +faststart -vcodec libx264 $OUTPUT
echo "[simkid] All done, thanks!"