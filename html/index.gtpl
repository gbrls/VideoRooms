<!DOCTYPE html>

<html>

<head>
    <script>

        document.addEventListener("DOMContentLoaded", function () {initialiseMediaPlayer();}, false);

        var mediaPlayer;
        var webSocket;
        var latestTime;

        function initialiseMediaPlayer() {
            mediaPlayer = document.getElementById('media-video');
            mediaPlayer.controls = true;

            //syncVideo();

            setSocket();
            mediaPlayer.play();

        }

        function setSocket() {
            if (webSocket) {
                return false;
            }

            webSocket = new WebSocket("{{.}}");

            webSocket.onopen = function (evt) {
                console.log("OPEN");
            }
            webSocket.onclose = function (evt) {
                console.log("CLOSE");
                webSocket = null;
            }
            webSocket.onmessage = function (evt) {
                console.log("RESPONSE: " + evt.data);
                latestTime = evt.data;
                console.log(latestTime);
                mediaPlayer.currentTime = parseFloat(latestTime);

                //syncVideo();
            }
        }

        function syncVideo() {
            setSocket();

            if (webSocket) {
                webSocket.send(mediaPlayer.currentTime);
            }

            //mediaPlayer.currentTime = parseFloat(latestTime);

            var btn = document.getElementById('play-pause-button');

            btn.title = 'sync';
            btn.innerHTML = 'sync';
            btn.className = 'sync';




            //if (mediaPlayer.paused || mediaPlayer.ended) {
            //    btn.title = 'pause';
            //    btn.innerHTML = 'pause';
            //    btn.className = 'pause';
            //    mediaPlayer.play();
            //
            //}
            //else {
            //    btn.title = 'play';
            //    btn.innerHTML = 'play';
            //    btn.className = 'play';
            //    mediaPlayer.pause();
            //}

        }


    </script>

    <head>

    <body>
        <h1>Test</h1>

        <div id='media-player'>
            <video id='media-video' width="800" controls>

                <source src='/static/video.mp4' type='video/mp4'>

            </video>
            <div id='media-controls'>

                <button id='play-pause-button' class='play' title='play' onclick='syncVideo();'>Play</button>

            </div>
        </div>
    </body>

</html>