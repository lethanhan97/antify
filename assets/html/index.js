const VIDEO_SOURCE = "http://localhost:8080/video/sample.m3u8";
const VIDEO_IDS = [
  "video-left-1",
  "video-left-2",
  "video-main",
  "video-right-1",
  "video-right-2",
];

if (Hls.isSupported()) {
  VIDEO_IDS.forEach((videoId) => {
    handleHlsVideo(videoId);
  });
} else if (video.canPlayType("application/vnd.apple.mpegurl")) {
  // hls.js is not supported on platforms that do not have Media Source Extensions (MSE) enabled.
  // When the browser has built-in HLS support (check using `canPlayType`), we can provide an HLS manifest (i.e. .m3u8 URL) directly to the video element through the `src` property.
  // This is using the built-in support of the plain video element, without using hls.js.
  handleNativeVideo(video);
}

function handleHlsVideo(videoId) {
  const video = document.getElementById(videoId);

  if (!video) {
    throw new Error("Video not found");
  }

  const hls = new Hls({
    debug: true,
  });
  hls.loadSource(VIDEO_SOURCE);
  hls.attachMedia(video);
  hls.on(Hls.Events.MEDIA_ATTACHED, function () {
    video.muted = true;
    video.loop = true;
    video.play();
  });
}

function handleNativeVideo(video) {
  video.src = VIDEO_SOURCE;
  video.addEventListener("canplay", function () {
    video.play();
  });
}
