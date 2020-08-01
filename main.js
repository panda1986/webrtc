function hasUserMedia() {
    navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMdia;
    return !!navigator.getUserMedia;
}

function hasRTCPeerConnection() {
    window.RTCPeerConnection = window.RTCPeerConnection || window.webkitRTCPeerConnection || window.mozRTCPeerConnection;
    return !! window.RTCPeerConnection;
}

function startPeerConnection(stream) {
    var configuration = {
        //添加自定义iceServers
        //"iceServers": [{"urls": "stun:127.0.0.1:9876"}]
    };
    yourConnection = new webkitRTCPeerConnection(configuration);
    theirConnection = new webkitRTCPeerConnection(configuration);

    // 处理ice
    yourConnection.onicecandidate = function (event) {
        if (event.candidate) {
            theirConnection.addIceCandidate(new RTCIceCandidate(event.candidate))
        }
    };
    theirConnection.onicecandidate = function (event) {
        if (event.candidate) {
            yourConnection.addIceCandidate(new RTCIceCandidate(event.candidate))
        }
    };

    // 开始offer
    yourConnection.createOffer(function (offer) {
        yourConnection.setLocalDescription(offer);
        theirConnection.setRemoteDescription(offer);

        theirConnection.createAnswer(function (offer) {
            theirConnection.setLocalDescription(offer);
            yourConnection.setRemoteDescription(offer);
        })
    })
}

var yourVideo = document.querySelector("#yours");
var theirVideo = document.querySelector("#theirs");
var yourConnection, theirConnection;
if (hasUserMedia()) {
    navigator.getUserMedia({
        video: true,
        audio: false
    }, function (stream) {
        if ('srcObject' in yourVideo) {
            yourVideo.srcObject = stream
        } else {
            yourVideo.src = URL.createObjectURL(stream)
        }
        if (hasRTCPeerConnection()) {
            startPeerConnection(stream);
        } else {
            alert("sorry, your browser doesn't has peer connection")
        }
    }, function (error) {
        alert("sorry, we failed to capture your camera, please try again, error is:" + error)
    })
} else {
    alert("sorry, your browser doesn't has user media")
}