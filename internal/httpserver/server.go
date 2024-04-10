package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
)

func indexRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, `
<a href="/metrics">metrics</a>
<br />
<a href="/healthz">healthz</a>
<br />
<a href="/items">items</a>
`)
}

func healthzRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "ok")
}

func metricsRoute(w http.ResponseWriter, req *http.Request) {
	items, err := fmipcore.ReadItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %q", err)
		return
	}

	gaugeDesc := func(name, desc string) {
		fmt.Fprintf(w, "# HELP %s %s\n", name, desc)
		fmt.Fprintf(w, "# TYPE %s gauge\n", name)
	}
	gaugeValue := func(name string, labels map[string]string, val float64) {
		fmt.Fprintf(w, "%s{", name)
		var labelKeys []string
		for k, v := range labels {
			labelKeys = append(labelKeys, fmt.Sprintf("%s=\"%s\"", k, v))
		}
		sort.Strings(labelKeys)
		fmt.Fprintf(w, strings.Join(labelKeys, ", "))
		fmt.Fprintf(w, "} %f\n", val)
	}
	itemKeys := func(item fmipcore.FmItem) map[string]string {
		return map[string]string{
			"name":   item.Name,
			"serial": item.SerialNumber,
		}
	}

	gaugeDesc("findmy_item_location_timestamp", "The epoch timestamp that the item was last seen")
	for _, item := range items {
		gaugeValue("findmy_item_location_timestamp", itemKeys(item), float64(item.Location.Timestamp))
	}

	gaugeDesc("findmy_item_location_latitude", "The latitude of each item")
	for _, item := range items {
		gaugeValue("findmy_item_location_lat", itemKeys(item), float64(item.Location.Latitude))
	}

	gaugeDesc("findmy_item_location_longitude", "The longitude of each item")
	for _, item := range items {
		gaugeValue("findmy_item_location_lat", itemKeys(item), float64(item.Location.Longitude))
	}

	gaugeDesc("findmy_item_location_old", "If the item is old")
	for _, item := range items {
		v := 0.0
		if item.Location.IsOld {
			v = 1.0
		}
		gaugeValue("findmy_item_location_old", itemKeys(item), v)
	}

	gaugeDesc("findmy_item_location_inaccurate", "If the item is inaccurate")
	for _, item := range items {
		v := 0.0
		if item.Location.IsInaccurate {
			v = 1.0
		}
		gaugeValue("findmy_item_location_inaccurate", itemKeys(item), v)
	}

	gaugeDesc("findmy_last_fetch_timestamp", "Last metric fetch timestamp from macos-findmy-mqtt")
	gaugeValue("findmy_last_fetch_timestamp", map[string]string{}, float64(time.Now().UnixMilli()))
}

func itemsRoute(w http.ResponseWriter, req *http.Request) {
	items, err := fmipcore.ReadItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %q", err)
		return
	}

	out, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %q", err)
		return
	}

	fmt.Fprintf(w, "%s", string(out))
}

func Run(port int) {
	if port == 0 {
		return
	}

	fmt.Printf("webserver listening on port :%d\n", port)

	http.HandleFunc("/", indexRoute)
	http.HandleFunc("/healthz", healthzRoute)
	http.HandleFunc("/metrics", metricsRoute)
	http.HandleFunc("/items", itemsRoute)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
