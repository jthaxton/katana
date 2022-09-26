# Katana

## About
Katana is a service that slices your mp4 files into png, and returns them in a .zip.

## Set up
```
docker pull 14jthaxton/katana:1.0
docker run -p 8080:8080 14jthaxton/katana:1.0
```
This image is ~3GB... Grab a coffee. Take a walk. This part might take a few minutes.

## Talking to the service
```
curl -X 'POST' \
  'http://localhost:8080/parse' \
  -F file=@full/path/to/video.mp4 \
  -H 'Content-Type: multipart/form-data' \
  --output ./output.zip
```
This will slice a video into png screenshots.
Grab a coffee. Take a walk. This part might take a few minutes.

```
curl -X 'GET' \
  'http://0.0.0.0:8080/video?name=<some_youtube_link>' \
  -H 'Content-Type: application/json'
```
This will download a video from youtube.
Grab a coffee. Take a walk. This part might take a few minutes.

## Downsides
For a random 10 second video, `/parse` returned ~700MB (compressed) of png files. In the future, I will include a body param to configure frequency of slices.