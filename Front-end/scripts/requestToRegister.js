document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('registrationForm');
    if (form) {
        form.addEventListener('submit', sendData);
    } else {
        console.error("Форма не найдена!");
    }
});

function sendData(event) {
    event.preventDefault(); // Предотвращаем стандартное поведение формы

    console.log("Form submitted");

    const username = document.getElementById('username').value; // Получаем значение из поля ввода
    const password = document.getElementById('password').value; // Получаем значение из поля ввода

    fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    })
    .then(response => {
        if (!response.ok) {
            // Если ответ не успешен, выбрасываем ошибку
            return response.json().then(errorData => {
                throw new Error(errorData.error || 'error registration');
            });
        }
        return response.json(); 
    })
    .then(data => {
        openModal(data.message)
        console.log("open modal window")
    })
    .catch(error => {
        console.log("open modal window error")
        console.error(error);
        openModalError(error.message)
    });
}
