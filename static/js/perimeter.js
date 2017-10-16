var eventSource = null;

var lastEvent = 0;
var eventMonitor = null;

if (!Date.now) {
    Date.now = function() { return new Date().getTime(); }
}

function initSSE(channel) {
    if (!window.EventSource) {
        alert("EventSource is not enabled in this browser");
        return;
    }

    eventSource = new EventSource('/stream/' + channel);

    eventSource.onopen = onSSEStatusChange;
    eventSource.onerror = onSSEStatusChange;

    eventSource.addEventListener("notification", onSSENotification);
    eventSource.addEventListener("ping", onSSEPing);
    eventSource.addEventListener("callback", onSSECallBack);
}

function onSSEStatusChange(e) {
    clearInterval(eventMonitor);

    if (eventSource.readyState === EventSource.OPEN) {
        eventMonitor = setInterval(checkSSEState, 2000);
    }

    console.log("State: " + eventSource.readyState);
}

function onSSENotification(e) {
    var data = JSON.parse(e.data);

    Materialize.toast(data.message, 5000, data.class);
}

function onSSEPing(e) {
    lastEvent = Date.now();
}

function onSSECallBack(e) {

}

function checkSSEState() {
    if ((lastEvent > 0) && (lastEvent < Date.now() - 30000)) {
        Materialize.toast("No PING for " + (Date.now() - lastEvent) + "ms", 5000, "red");
    }
}