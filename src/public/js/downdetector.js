const url =
  "http://192.168.1.240:8081/data/detector?key=sg.lyceum130.ru,www.google.com,www.vk.com,www.github.com,www.instagram.com,www.youtube.com,www.twitter.com";

var xmlhttp = new XMLHttpRequest();
xmlhttp.onreadystatechange = function () {
  if (this.readyState == 4 && this.status == 200) {
    var myObj = JSON.parse(this.responseText);

    postMessage(myObj);
  }
};
xmlhttp.open("GET", url, true);
xmlhttp.send();
setInterval(() => {
  var xmlhttp = new XMLHttpRequest();
  xmlhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      var myObj = JSON.parse(this.responseText);

      postMessage(myObj);
    }
  };
  xmlhttp.open("GET", url, true);
  xmlhttp.send();
}, 60000);
