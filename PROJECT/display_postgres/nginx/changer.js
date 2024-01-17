document.addEventListener("DOMContentLoaded", function () {
    const change_first = document.querySelector(".change_first");
    const change_class = document.querySelector(".change_class");
    const change_button = document.querySelector(".change_button");
  
    change_button.addEventListener("click", function () {
      const first_name = change_first.value;

      const class_number = change_class.value;
  
      const userData = { first_name: first_name, class_number: class_number };
      console.log(userData)
      alert("User class change successfully!");


      // POST so'rovi uchun fetch ishlatilishi
      fetch('http://65.21.241.48:5000/change_class', {
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
  
  
  