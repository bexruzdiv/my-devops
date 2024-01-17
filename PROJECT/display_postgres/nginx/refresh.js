document.addEventListener("DOMContentLoaded", function () {
    const refresh_button = document.querySelector(".refresh_button");
  
    refresh_button.addEventListener("click", function () {
    

      alert("Data list refresh successfuly! ")

      // POST so'rovi uchun fetch ishlatilishi
      fetch('http://65.21.241.48:5000/data', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*',
          'Access-Control-Allow-Methods': 'POST, PUT, DELETE, GET, OPTIONS',
          'Access-Control-Request-Method': '*',
          'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept, Authorization',
        },
        body: JSON.stringify(),
      })
      .then(response => response.json())
      .then(data => {
        // Serverdan qaytgan javobni ishlash
        console.log('Serverdan qaytgan javob:', data);
      })
      .catch(error => {
        console.error('Xatolik yuz berdi:', error);
      });
  
    });
  });
  
  
  