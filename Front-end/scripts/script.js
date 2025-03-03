document.addEventListener('DOMContentLoaded', () => {
    const eyeIcon = document.getElementById('togglePassword');
    eyeIcon.addEventListener('click', togglePassword);
});

function togglePassword() {
    const passwordInput = document.getElementById('password');
    const eyeIcon = document.getElementById('togglePassword');
    
    if (passwordInput.type === 'password') {
        passwordInput.type = 'text'; // Меняем тип на текст
        eyeIcon.src = '/Front-end/images/open-eye.png'; // Путь к изображению открытого глаза
    } else {
        passwordInput.type = 'password'; // Меняем тип обратно на пароль
        eyeIcon.src = '/Front-end/images/closed-eye.png'; // Путь к изображению закрытого глаза
    }
}