function formtime(typ) {
  if (typ < 10) {
    return (typ = "0" + typ);
  } else {
    return typ;
  }
}
setInterval(() => {
  var today = new Date();
  var day = formtime(today.getDate());
  var month = formtime(today.getMonth() + 1);
  var hours = formtime(today.getHours());
  var minutes = formtime(today.getMinutes());
  var seconds = formtime(today.getSeconds());
  var date = day + "." + month + "." + today.getFullYear();
  var time = hours + ":" + minutes + ":" + seconds;
  postMessage([date, time]);
}, 0.5);
