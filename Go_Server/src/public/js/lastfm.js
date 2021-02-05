var xmlhttp = new XMLHttpRequest();
xmlhttp.onreadystatechange = function () {
  if (this.readyState == 4 && this.status == 200) {
    var myObj = JSON.parse(this.responseText);

    postMessage(myObj);
  }
};
xmlhttp.open(
  "GET",
  "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=Petruccinator69&api_key=ad57f38806de299b643b55377483af01&format=json&limit=10",
  true
);
xmlhttp.send();
setInterval(() => {
  var xmlhttp = new XMLHttpRequest();
  xmlhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      var myObj = JSON.parse(this.responseText);

      postMessage(myObj);
    }
  };
  xmlhttp.open(
    "GET",
    "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=Petruccinator69&api_key=ad57f38806de299b643b55377483af01&format=json&limit=10",
    true
  );
  xmlhttp.send();
}, 30000);
