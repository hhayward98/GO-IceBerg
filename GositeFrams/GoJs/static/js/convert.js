window.addEventListener("DOMContentLoaded", domLoaded);

function domLoaded() {
   let Conv = document.getElementById('convertButton');
   let Cel = document.querySelector("#cInput");
   let Fe = document.getElementById('fInput');
   let errmsg = document.querySelector('#errorMessage');
   

   //on key press run event to remove text from fInput
   Cel.addEventListener("input", function() {
      Fe.value = "";
   });

   // on key press run event to remove text from cInput
   Fe.addEventListener("input", function() {
      Cel.value = "";
   });


   Conv.addEventListener("click", function(){
      errmsg.textContent = "";
      if (Cel.value === "") {
         try {
            var temp = parseInt(Fe.value);
         } catch (err){
            console.log("Error: "+err);
         }
         if (isNum(Fe.value) === false) {
            errmsg.textContent = Fe.value + " is not a number";
            return;
         }
         let NCe = convertFtoC(temp);
         changeIMG(temp);
         Cel.value = NCe;
         return;



      } else if (Fe.value === "") {
         try {
            var temp = parseInt(Cel.value);
         } catch (err){
            console.log(err);
         }
         if (isNum(Cel.value) === false) {
            errmsg.textContent = Cel.value + " is not a number";
            return;
         }
         let NFe = convertCtoF(temp);
         changeIMG(NFe);
         Fe.value = NFe;
         return;
      }

   });

}

function isNum(val) {
   return !isNaN(val)
}

function convertCtoF(degreesCelsius) {
   let F = degreesCelsius * 9/5 + 32;
   return F;
}

function convertFtoC(degreesFahrenheit) {
   let C = (degreesFahrenheit - 32) * 5/9;
   return C;
}

// temp is in F
function changeIMG(tempF) {
   console.log(tempF)
   let img = document.getElementById('weatherImage');


   if (tempF <= 32) {
      img.src = "/static/images/cold.png";
   }else if (tempF <= 50) {
      img.src = "/static/images/cool.png";
   }else if (tempF > 50) {
      img.src = "/static/images/warm.png";
   }

}