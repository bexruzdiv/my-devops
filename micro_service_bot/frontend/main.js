document.addEventListener("DOMContentLoaded", function () {
    const delete_first = document.querySelector(".text");

    const delete_button = document.querySelector(".button");
  
    delete_button.addEventListener("click", function () {
      const first_name = delete_first.value;

  
      const userData = { first_name: first_name};
      console.log(userData)
      alert("User delete from postgres successfuly! ")

      // POST so'rovi uchun fetch ishlatilishi
      fetch('http://65.21.241.48:30111/send', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*',
          'Access-Control-Allow-Methods': 'POST, PUT, DELETE, GET, OPTIONS',
          'Access-Control-Request-Method': '*',
          'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept, Authorization',
        },
        body: JSON.stringify(userData),
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
  
  
  