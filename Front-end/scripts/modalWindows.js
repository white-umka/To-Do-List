var modal = document.getElementById("myModal");
var modalError = document.getElementById("myModalError");
var span = document.getElementsByClassName("close")[0];

// Функция для открытия модального окна с переданным текстом
function openModal(text) {
    document.getElementById("modalText").innerText = text;
    modal.style.display = "block";
    modal.classList.add("show"); // Добавляем класс show
    modal.querySelector(".modal-content").classList.add("show"); // Добавляем класс show для содержимого
    console.log("open modal window");
}


function openModalError(text) {
    document.getElementById("modalTextError").innerText = text;
    modalError.style.display = "block";
    modalError.classList.add("show");
    modalError.querySelector(".modal-content").classList.add("show");
    console.log("open modal window error");
}


// Функция для закрытия модального окна с ошибкой
function closeModalError() {
    modalError.style.display = "none";
    modalError.classList.remove("show");
    modalError.querySelector(".modal-content").classList.remove("show");
}

// Закрытие модального окна при клике вне его области
window.onclick = function(event) {
    if (event.target == modal) {
        closeModal();
    }
    if (event.target == modalError) {
        closeModalError();
    }
};

// Функция для перенаправления на сайт
function redirectToWebsite() {
    window.location.href = 'https://github.com';
}


