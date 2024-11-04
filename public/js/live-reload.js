document.addEventListener("DOMContentLoaded", function(event) {
    let readyForReload = false;
    let conn ;

    function connect() {
        conn = new WebSocket("ws://" + document.location.host + "/ws/1");
        conn.onopen = function() {
            if (readyForReload) {
                location.reload();
            }
        }

        conn.onclose = function (evt) {
            console.log("Server restart, wait for up")
            readyForReload = true;

            setInterval(() => {
                connect();
            }, 1000);
        };

        conn.onerror = function (evt) {
            readyForReload = true;

            setInterval(() => {
                connect();
            }, 1000);
        }
    }

    connect();
});
