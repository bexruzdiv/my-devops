document.addEventListener("DOMContentLoaded", function () {
    const add_first = document.querySelector(".add_first");
    const add_last = document.querySelector(".add_last");
    const add_class = document.querySelector(".add_class");
    const add_button = document.querySelector(".add_button");
  
    add_button.addEventListener("click", function () {
      const first_name = add_first.value;
      const last_name = add_last.value;
      const class_number = add_class.value;
  
      const userData = { first_name: first_name, last_name: last_name, class_number: class_number };
      console.log(userData)
      // POST so'rovi uchun fetch ishlatilishi
      alert("User add to postgres successfuly! ")
      fetch('http://localhost:5000/add', {
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
  
  
  