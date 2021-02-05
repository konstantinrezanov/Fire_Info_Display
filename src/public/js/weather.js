var xmlhttp = new XMLHttpRequest();
const key = "90a74e7934e549966c49076196f53888";
xmlhttp.onreadystatechange = function () {
  if (this.readyState == 4 && this.status == 200) {
    var myObj = JSON.parse(this.responseText);

    postMessage(myObj);
  }
};
xmlhttp.open("GET", "https://api.openweathermap.org/data/2.5/weather?q=Ekaterinburg&units=metric&appid="+key, true);
xmlhttp.send();
setInterval(() => {
  var xmlhttp = new XMLHttpRequest();
  xmlhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      var myObj = JSON.parse(this.responseText);

      postMessage(myObj);
    }
  };
  xmlhttp.open("GET", "https://api.openweathermap.org/data/2.5/weather?q=Ekaterinburg&units=metric&appid="+key, true);
  xmlhttp.send();
}, 60000);
