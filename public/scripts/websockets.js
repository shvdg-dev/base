let socket = new WebSocket("ws://localhost:8080/websocket");

socket.onopen = function(event){
    console.log("Connection established");
}

socket.onmessage = function (event) {
    console.log("Received data" + event.data);
}

socket.onclose = function(event){
    console.log("Connection closed");
}