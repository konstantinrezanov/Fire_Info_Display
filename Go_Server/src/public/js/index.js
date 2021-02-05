var timeWorker = new Worker("js/time.js");
var update = new Worker("js/update.js");

timeWorker.onmessage = function (e) {
  document.getElementById("date").innerHTML = e.data[0];
  document.getElementById("time").innerHTML = e.data[1];
};

update.onmessage = function (e) {
  document.getElementById("trans").innerHTML = "";

  document.getElementById("temp").innerHTML = e.data.Weather.Temp + "&deg;C";
  document.getElementById("feels_like").innerHTML =
    e.data.Weather.FeelsLike + "&deg;C";
  document.getElementById("pressure").innerHTML =e.data.Weather.Pressure + "hPa";

  for (let n of e.data.Stations) {
    let station=document.createElement('div')
    station.id=n.ID
    let title=document.createElement("p")
    title.innerHTML=(n.Status.Title).split(",")[0].slice(1).trim()
    station.appendChild(title)
    let info=document.createElement("div")
    for (let i=0; i<n.Status.Routes.length; i++) {
        let line=document.createElement("p")
        line.innerHTML =
          n.Status.Routes[i] +
          " &#09;" +
          n.Status.Distances[i] +
          "&nbsp;" +
          n.Status.Arrivals[i] +
          "&nbsp;";
        info.appendChild(line)
    }
    station.append(info)
    document.getElementById("trans").appendChild(station)
  }
}