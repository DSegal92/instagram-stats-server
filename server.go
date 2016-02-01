package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	title := "Statistics"
	var body bytes.Buffer

	body.WriteString(`<div>
											<table>
												<thead>
													<tr>
														<th> Date </th>
														<th> Follows </th>
														<th> Followers </th>
									 `)

	stats := getStatistics()

	for i := 0; i < len(stats); i++ {
		row := fmt.Sprintf("<tr><td>%v</td><td>%v</td><td>%v</td>", stats[i].Date, stats[i].Follows, stats[i].Followers)
		body.WriteString(row)
	}

	body.WriteString(`			</tr>
												</thead>
											</table>
										</div>`)

	fmt.Fprintf(w, "<h1>%s</h1>%s", title, body.String())
}

func main() {
	fmt.Println()
	http.HandleFunc("/instagram/stats/", statsHandler)
	http.ListenAndServe(":8080", nil)
}
