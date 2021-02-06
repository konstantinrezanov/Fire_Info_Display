let timeWorker = new Worker("js/time.js");
let weatherWorker = new Worker("js/weather.js");
let stationWorker = new Worker("js/station.js");
let lastFmWorker = new Worker("js/lastfm.js");
let downWorker = new Worker("js/downdetector.js");

timeWorker.onmessage = function (e) {
  document.getElementById("date").innerHTML = e.data[0];
  document.getElementById("time").innerHTML = e.data[1];
};

weatherWorker.onmessage = function (e) {
  document.getElementById("temp").innerHTML = e.data.main.temp + "&deg;C";
  document.getElementById("feels_like").innerHTML =
    e.data.main.feels_like + "&deg;C";
  document.getElementById("pressure").innerHTML = e.data.main.pressure + "hPa";
};

stationWorker.onmessage = function (e) {
  document.getElementById("trans").innerHTML = "";

  for (let n of e.data.List) {
    let station = document.createElement("div");
    station.id = n.ID;
    let title = document.createElement("p");
    title.innerHTML = n.Status.Title.split(",")[0].slice(1).trim();
    station.appendChild(title);
    let info = document.createElement("div");
    for (let i = 0; i < n.Status.Routes.length; i++) {
      let line = document.createElement("p");
      line.innerHTML =
        n.Status.Routes[i] +
        " &#09;" +
        n.Status.Distances[i] +
        "&nbsp;" +
        n.Status.Arrivals[i] +
        "&nbsp;";
      info.appendChild(line);
    }
    station.append(info);
    document.getElementById("trans").appendChild(station);
  }
};

lastFmWorker.onmessage = function (e) {
  let music = document.getElementById("music");
  music.innerHTML = "";
  for (let n of e.data.recenttracks.track) {
    let trackInfo = document.createElement("div");

    let artist = document.createElement("p");
    let album = document.createElement("p");
    let track = document.createElement("p");

    artist.innerHTML = "Artist: " + n.artist["#text"];
    album.innerHTML = "Album: " + n.album["#text"];
    track.innerHTML = "Track: " + n.name;

    trackInfo.appendChild(artist);
    trackInfo.appendChild(album);
    trackInfo.appendChild(track);
    trackInfo.appendChild(document.createElement("hr"));
    music.appendChild(trackInfo);
  }
};

downWorker.onmessage = function (e) {
  let block = document.getElementById("downdetector");
  block.innerHTML=""
  for (let n of e.data.List) {
    let name = document.createElement("p")
    let status = document.createElement("p")
    name.innerHTML=n.Name
    status.innerHTML=n.Up?"Up":"Down"
    block.appendChild(name)
    block.appendChild(status)
    block.appendChild(document.createElement("hr"))
  }
};
