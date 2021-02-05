var xmlhttp= new XMLHttpRequest();
xmlhttp.onreadystatechange=function() {
    if (this.readyState == 4 && this.status == 200) {
      var myObj = JSON.parse(this.responseText);

      postMessage(myObj);
    }
}
xmlhttp.open("GET", "http://192.168.1.240:8081/data", true);
xmlhttp.send();
setInterval(()=>{
   var xmlhttp = new XMLHttpRequest();
   xmlhttp.onreadystatechange = function () {
     if (this.readyState == 4 && this.status == 200) {
       var myObj = JSON.parse(this.responseText);

       postMessage(myObj);
     }
   };
   xmlhttp.open("GET", "http://192.168.1.240:8081/data", true);
   xmlhttp.send(); 
},60000)