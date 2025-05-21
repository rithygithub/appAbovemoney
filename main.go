package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "time"
)

func aspirationsHandler(w http.ResponseWriter, r *http.Request) {
    var aspirations []string

    file, err := os.Open("aspirations.json")
    if err != nil {
        http.Error(w, "Failed to read aspirations", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    if err := json.NewDecoder(file).Decode(&aspirations); err != nil {
        http.Error(w, "Failed to parse aspirations", http.StatusInternalServerError)
        return
    }

    now := time.Now()
    dayOfYear := now.YearDay()
    hour := now.Hour()

    // Seed random generator with day and hour for deterministic results
    seed := int64(dayOfYear*100 + hour)
    rng := rand.New(rand.NewSource(seed))

    // Shuffle and pick first 2 aspirations
    rng.Shuffle(len(aspirations), func(i, j int) {
        aspirations[i], aspirations[j] = aspirations[j], aspirations[i]
    })

    // Prepare HTML output with magenta background and larger text
    fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>Daily Aspirations</title>
    <style>
        body {
            background-color: magenta;
            color: white;
            font-family: Arial, sans-serif;
            text-align: center;
        }
        h1 {
            font-size: 32px;
        }
        ul {
            list-style-type: none;
            padding: 0;
        }
        li {
            font-size: 24px;
            margin: 10px 0;
        }
    </style>
</head>
<body>
    <h1>Aspirations for Day %03d, Hour %02d</h1>
    <ul>`, dayOfYear, hour)

    count := 2
    if len(aspirations) < 2 {
        count = len(aspirations)
    }
    for i := 0; i < count; i++ {
        fmt.Fprintf(w, "<li>%s</li>", aspirations[i])
    }

    fmt.Fprintf(w, `
    </ul>
</body>
</html>`)
}

func main() {
    http.HandleFunc("/", aspirationsHandler)
    fmt.Println("Listening on :8080")
    http.ListenAndServe(":8080", nil)
}
