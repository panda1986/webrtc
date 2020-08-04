var name, connectedUser;
var connection = new WebSocket('ws://localhost:8888/signal');
connection.onopen = function (ev) {
    console.log("Connected");
};

//通过回调函数处理所有的消息
connection.onmessage = function (msg) {
    console.log("Got message", msg.data);
    var data = JSON.parse(msg.data);
    switch (data.type) {
        case "login":
            onLogin(data.success);
            break;
        case "offer":
            onOffer(data.offer, data.name);
            break;
        case "answer":
            onAnswer(data.answer);
            break;
        case "candidate":
            onCandidate(data.candidate);
            break;
        case "leave":
            onLeave();
            break;
        default:
            break;
    }
};

connection.onerror = function (err) {
    console.log("Got error", err)
};

// Alias 以Json格式发送消息
function send(msg) {
    if (connectedUser) {
        msg.name = connectedUser
    }
    connection.send(JSON.stringify(msg))
}

// 登录到应用程序，与服务器进行交互
var loginPage = document.querySelector("#login-page");
var usernameInput = document.querySelector("#username");
var loginButton = document.querySelector('#login');

var callPage = document.querySelector('#call-page');
var theirUsernameInput = document.querySelector('#their-username');
var callButton = document.querySelector('#call');
var hangButton = document.querySelector('#hang-up');
var sendButton = document.querySelector("#send");
var messageInput  = document.querySelector("#message");
var received = document.querySelector("#received");

callPage.style.display = "none";
loginButton.addEventListener("click", function (event) {
    name = usernameInput.value;
    if (name.length > 0) {
        send({
            type: "login",
            name: name
        })
    }
});

function onLogin(success) {
    if (success  === false) {
        alert("login failed, please try a different name")
    } else {
        loginPage.style.display = "none";
        callPage.style.display = "block";
        //准备好通话的通道
        startConnection();
    }
}

/*
开始一个对等连接, 步骤如下：
从相机中获取视频流
验证用户的浏览器是否支持webrtc
创建RTCPeerConnection
*/

var yourVideo = document.querySelector("#yours");
var theirVideo = document.querySelector("#theirs");
var yourConnection, theirConnection, stream;

function startConnection() {
    if (hasUserMedia()) {
        navigator.getUserMedia({
            video: true,
            audio: false
        }, function (myStream) {
            stream = myStream;
            if ('srcObject' in yourVideo) {
                yourVideo.srcObject = stream
            } else {
                yourVideo.src = URL.createObjectURL(stream)
            }
            if (hasRTCPeerConnection()) {

                setupPeerConnection(stream);
            } else {
                alert("sorry, your browser doesn't has peer connection")
            }
        }, function (error) {
            alert("sorry, we failed to capture your camera, please try again, error is:" + error)
        })
    } else {
        alert("sorry, your browser doesn't has user media")
    }
}

function setupPeerConnection(stream) {
    var configuration = {
        //添加自定义iceServers
        //"iceServers": [{"urls": "stun:127.0.0.1:9876"}]
    };
    yourConnection = new webkitRTCPeerConnection(configuration);
    openDataChannel();
    // 设置流的监听
    yourConnection.addStream(stream);
    yourConnection.onaddstream = function (ev) {
        theirVideo.srcObject = ev.stream
    };
    // 处理ice
    yourConnection.onicecandidate = function (event) {
        if (event.candidate) {
            send({
                type : "candidate",
                body: {
                    candidate: event.candidate
                }
            })
        }
    };
}

function hasUserMedia() {
    navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMdia;
    return !!navigator.getUserMedia;
}

function hasRTCPeerConnection() {
    window.RTCPeerConnection = window.RTCPeerConnection || window.webkitRTCPeerConnection || window.mozRTCPeerConnection;
    return !! window.RTCPeerConnection;
}

/*
 发起通话
 和远程用户发起通话，首先发送offer给另一个用户来开始整个过程，一旦用户得到这个offer，他将创建一个响应并开始交换ICE候选，直到成功连接服务器。
*/

callButton.addEventListener("click", function () {
    var theirUsername = theirUsernameInput.value;
    if (theirUsername.length > 0) {
        startPeerConnection(theirUsername)
    }
});

function startPeerConnection(user) {
    connectedUser = user;
    // 开始创建offer
    yourConnection.createOffer(function (sdp) {
        console.log("send offer sdp:", sdp);
        send({
            type: "offer",
            body: {
                offer: sdp
            }
        });
        yourConnection.setLocalDescription(sdp);
    }, function (error) {
        console.log("create offer failed", error);
        alert("create offer failed")
    })
}

function onOffer(sdp, name) {
    connectedUser = name;
    yourConnection.setRemoteDescription(new RTCSessionDescription(sdp));
    yourConnection.createAnswer(function (sdp2) {
        yourConnection.setLocalDescription(sdp2);
        send({
            type: "answer",
            body: {
                answer:  sdp2
            }
        })
    }, function (error) {
        console.log("create answer failed");
        alert("create answer failed")
    })
}

function onAnswer(sdp) {
    yourConnection.setRemoteDescription(new RTCSessionDescription(sdp))
}

function onCandidate(candidate) {
    yourConnection.addIceCandidate(new RTCIceCandidate(candidate))
}

hangButton.addEventListener("click", function () {
    send({
        type: "leave"
    });
    onLeave()
});

function onLeave() {
    connectedUser = null;
    theirVideo.srcObject = null;
    yourConnection.close();
    yourConnection.onicecandidate = null;
    yourConnection.onaddstream = null;
    setupPeerConnection(stream);
}

function openDataChannel() {
    var dataChannelOptions = {
        reliable: true,
        negotiated: true,
        id: 0
    };
    dataChannel = yourConnection.createDataChannel("chat", dataChannelOptions);
    dataChannel.onerror = function (error) {
        console.log("data channel error", error)
    };
    dataChannel.onmessage = function (event) {
        console.log("got data channel message", event.data);
        received.innerHTML += "recv " + event.data + "<br />";
        received.scrollTop = received.scrollHeight;
    };
    dataChannel.onopen = function () {
        dataChannel.send(name + " has connected")
    };
    dataChannel.onclose = function () {
        console.log("the data channel is closed")
    }
}

sendButton.addEventListener("click", function (event) {
    var val = messageInput.value;
    received.innerHTML += "send: " + val + "<br />";
    received.scrollTop = received.scrollHeight;
    dataChannel.send(val);
});


