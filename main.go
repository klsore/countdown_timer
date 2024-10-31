<!DOCTYPE html>
<html>
<head>
    <title>Bitcoin Halving Countdown</title>
    <link rel="stylesheet" type="text/css" href="/bitcoinhalving/static/style.css">
</head>
<body>
    <div class="container">
        <h1 class="futuristic-text">Bitcoin Halving Countdown</h1>
        <div id="timer" class="futuristic-timer">Loading...</div>
    </div>

    <script>
        let eventSource = new EventSource('/bitcoinhalving/events');

        eventSource.onmessage = function (event) {
            document.getElementById("timer").innerText = event.data;
        };

        eventSource.onerror = function () {
            document.getElementById("timer").innerText = "Error";
        };

        // Fetch targetTime from the file
        fetch('/bitcoinhalving/static/targetTime.txt')
            .then(response => response.text())
            .then(targetTime => {
                // Update the HTML content with the fetched targetTime
                document.getElementById("targetTime").innerText = `Reward-Drop ETA date: ${targetTime}UTC`;
            })
            .catch(error => {
                console.error('Error fetching targetTime:', error);
            });
    </script>
   
    <div class="container2">
        <!-- The targetTime will be inserted here -->
        <h1 id="targetTime" class="text">Reward-Drop ETA date: Loading...</h1>
    </div>
</body>
</html>
