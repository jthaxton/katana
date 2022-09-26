FROM golang
RUN apt-get update && apt-get install -y yasm pkg-config
RUN git clone https://github.com/FFmpeg/FFmpeg.git
RUN apt-get install zlib1g-dev
RUN cd FFmpeg && ./configure --prefix=/usr/local/ffmpeg --enable-shared --enable-decoder=mjpeg,png --enable-demuxer=image2 --enable-protocol=file && make && make install
ENV PKG_CONFIG_PATH=/usr/local/ffmpeg/lib/pkgconfig/
RUN echo "include ld.so.conf.d/*.conf \n /usr/lib \n /usr/local/lib \n /usr/local/ffmpeg/lib" >> /etc/ld.so.conf
RUN ldconfig

# RUN youtube-dl

WORKDIR /src
ADD . .
RUN go mod download
RUN go get github.com/3d0c/gmf
RUN go build .

ENTRYPOINT [ "./m" ]
EXPOSE 8080