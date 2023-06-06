let pathArray = window.location.pathname.split("/")
let course_id = pathArray[pathArray.length - 1];
const container = document.getElementById("course_container");

if (course_id === null) {
    showDangerToast("Проверьте правильность введенных данных", false);
} else {
    fetch(`/api/v1/course/${course_id}`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    document.title = info.name;
                    printCourse(info);
                    container.className = "container mb-4 visible"
                })
            } else if (response.status === 204) {
                showDangerToast("По данному курсу не найдена информация", false);
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function printCourse(info) {
    let card_body = document.getElementById("card_body");
    card_body.innerHTML = `<h3 class="card-title text-center">${info.name}</h3>
                            <table class="table">
                                <tbody>
                                    <tr>
                                        <th scope="row">Направление</th>
                                        <td>${info.direction_name}</td>
                                    </tr>
                                    <tr>
                                        <th scope="row">Количество занятий</th>
                                        <td>${info.num_of_classes}</td>
                                    </tr>
                                    <tr>
                                        <th scope="row">Время одного занятия</th>
                                        <td>${info.class_time}</td>
                                    </tr>
                                    <tr>
                                        <th scope="row">Дни недели</th>
                                        <td>${info.week_days}</td>
                                    </tr>
                                    <tr>
                                        <th scope="row">Даты проведения курса</th>
                                        <td>${info.first_class_date.slice(8, 10) + "." + info.first_class_date.slice(5, 7) + "." + info.first_class_date.slice(0, 4)} - 
                                            ${info.last_class_date.slice(8, 10) + "." + info.last_class_date.slice(5, 7) + "." + info.last_class_date.slice(0, 4)}</td>
                                    </tr>
                                </tbody>
                            </table>
                            <h5 class="text-center">Цена: ${info.price}</h5>`;

    let card_info = document.getElementById("card_info");
    card_info.innerHTML = `<h5 class="text-center">Описание</h5>
    <p>${info.info}</p>`
}

function recordButton() {
    let elems = document.querySelectorAll("input.form-control");

    for (let elem of elems) {
        elem.className = "form-control form-control-sm";
        let danger_text = elem.nextElementSibling;
        danger_text.innerText = "";
    }

    let is_validated = true;
    for (let elem of elems) {
        if (elem.value === null || elem.value === "") {
            elem.className = "form-control form-control-sm border-danger";
            let danger_text = elem.nextElementSibling;
            danger_text.innerText = "Поле не может быть пустым";
            is_validated = false;
        }
    }

    if (!is_validated) {
        return;
    }

    let student = {
        name: document.getElementById("input_name").value,
        surname: document.getElementById("input_surname").value,
        patronymic: document.getElementById("input_patronymic").value,
        email: document.getElementById("input_email").value,
        phone: document.getElementById("input_phone").value,
        comment: document.getElementById("input_comment").value
    }
    fetch(`/api/v1/course/${course_id}`, {
        method: "POST",
        body: JSON.stringify(student)
    })
        .then(response => {
            if (response.status === 201) {
                response.json().then(id => {
                    window.location.replace(`/student/record/${id.student_id}`);
                });
            } else if (response.status === 400) {
                response.text().then(err => {
                    if (err.includes("mail")) {
                        let elem = document.getElementById("input_email");
                        elem.className = "form-control form-control-sm border-danger";
                        let danger_text = elem.nextElementSibling;
                        danger_text.innerText = "Неверный формат электронной почты";
                    } else if (err.includes("phone")) {
                        let elem = document.getElementById("input_phone");
                        elem.className = "form-control form-control-sm border-danger";
                        let danger_text = elem.nextElementSibling;
                        danger_text.innerText = "Неверный формат номера телефона";
                    } else {
                        showDangerToast("Проверьте правильность введенных данных", false);
                    }
                });
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function showDangerToast(message, is_server) {
    let toast_div = document.querySelector("div.toast-container");
    let elem = document.createElement("div");
    let btn_style
    if (is_server) {
        elem.className = "toast align-items-center text-bg-danger";
        btn_style = "btn-close-white"
    } else {
        elem.className = "toast align-items-center border-danger";
        btn_style = "btn-close-black"
    }
    elem.setAttribute("role", "alert");
    elem.setAttribute("aria-live", "assertive");
    elem.setAttribute("aria-atomic", "true");
    elem.innerHTML = `<div class="d-flex">
                    <div class="toast-body">
                        ${message}
                    </div>
                    <button type="button" class="btn-close ${btn_style} me-2 m-auto" data-bs-dismiss="toast" aria-label="Закрыть"></button>
                </div>`;
    toast_div.appendChild(elem);
    let toast = new bootstrap.Toast(elem);
    toast.show();
}