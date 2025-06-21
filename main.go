package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Current Date & Time</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                text-align: center;
                margin-top: 20%;
                background-color: #f4f4f4;
            }
            h1 {
                font-size: 2.5rem;
                color: #333;
            }
        </style>
    </head>
    <body>
        <h1 id="clock">Loading...</h1>
        <script>
            function updateTime() {
                const now = new Date();
                const formatted = now.getFullYear() + "-" +
                                  String(now.getMonth() + 1).padStart(2, '0') + "-" +
                                  String(now.getDate()).padStart(2, '0') + " " +
                                  String(now.getHours()).padStart(2, '0') + ":" +
                                  String(now.getMinutes()).padStart(2, '0') + ":" +
                                  String(now.getSeconds()).padStart(2, '0');
                document.getElementById('clock').textContent = "Current Date & Time: " + formatted;
            }

            setInterval(updateTime, 1000);
            updateTime(); // initial call
        </script>
    </body>
    </html>
    `
    fmt.Fprint(w, html)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at http://localhost:9090")
    http.ListenAndServe(":9090", nil)
}
