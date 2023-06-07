fetch(`/api/v1/editor`)
    .then(response => {
        if (response.status === 400) {
            showDangerToast("Проверьте правильность введенных данных", false);
        } if (response.status === 401) {
            window.location.replace("/editor/login");
        } else if (response.status === 500) {
            showDangerToast("Серверная ошибка, попробуйте позже", true);
        }
    });

function courseButton() {
    let buttons = document.querySelectorAll(".list-group-item");
    for (let b of buttons) {
        b.className = "list-group-item list-group-item-action";
    }
    let activeButton = document.getElementById("course_button");
    activeButton.className += " active";

    let searchButton = document.getElementById("search_button");
    searchButton.setAttribute("onclick", "searchCourse()");

    let search_string = document.getElementById("search_string");
    search_string.setAttribute("placeholder", "Поиск");

    if (document.getElementById("select_course") !== null) {
        document.getElementById("courses_select").innerHTML = ``;
    }

    let elems_table = document.getElementById("elems_table");

    fetch(`/api/v1/editor/courses`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    printCourses(elems_table, info);
                })
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                window.location.replace("/editor/login");
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function printCourses(elems_table, info) {
    let table_head = document.querySelector("thead");
    if (table_head !== null) {
        elems_table.removeChild(table_head);
    }

    let table_body = document.querySelector("tbody");
    if (table_body !== null) {
        elems_table.removeChild(table_body);
    }

    if (info !== null){
        info.sort((x, y) => x.id - y.id);

        let table_head_elem = document.createElement("thead");
        table_head_elem.innerHTML = `<tr>
            <th scope="col">#</th>
            <th scope="col">ID</th>
            <th scope="col">Курс</th>
            <th scope="col">Кол-во занятий</th>
            <th scope="col">Длительность занятия</th>
            <th scope="col">Дни недели</th>
            <th scope="col">Дата первого занятия</th>
            <th scope="col">Дата последнего занятия</th>
            <th scope="col">Цена</th>
            <th scope="col">Направление</th>
            <th scope="col">Информация</th>
            <th scope="col"></th>
        </tr>`;

        elems_table.appendChild(table_head_elem);
        let no_records_elem = document.getElementById("no_records");
        if (no_records_elem !== null) {
            let struct_list = document.getElementById("struct_list");
            struct_list.removeChild(no_records_elem);
        }

        let table_body_elem = document.createElement("tbody");
        for (let i = 0; i < info.length; i++) {
            let elem = document.createElement("tr");
            elem.innerHTML = `<th scope="col">${i + 1}</th>
                <td>${info[i].id}</td>
                <td>${info[i].name}</td>
                <td>${info[i].num_of_classes}</td>
                <td>${info[i].class_time}</td>
                <td>${info[i].week_days}</td>
                <td>${info[i].first_class_date.slice(8, 10) + "." + info[i].first_class_date.slice(5, 7) + "." + info[i].first_class_date.slice(0, 4)}</td>
                <td>${info[i].last_class_date.slice(8, 10) + "." + info[i].last_class_date.slice(5, 7) + "." + info[i].last_class_date.slice(0, 4)}</td>
                <td>${info[i].price} рублей</td>
                <td>${info[i].direction_name}</td>
                <td>${(info[i].info !== null) ? info[i].info : ""}</td>
                <td width="1%"><button class="btn btn-primary btn-sm" type="button" onclick="updateModalCourse(${info[i].id})">Редактировать</button></td>`;
            table_body_elem.appendChild(elem);
        }
        elems_table.appendChild(table_body_elem);
    } else {
        let no_records_elem = document.getElementById("no_records");
        if (no_records_elem === null) {
            let struct_list = document.getElementById("struct_list");
            let elem = document.createElement("div");
            elem.className = "row justify-content-center text-center";
            elem.setAttribute("id", "no_records");
            elem.innerHTML = "<p>Соответствующие курсы не были найдены</p>";
            struct_list.appendChild(elem);
        }
    }
}

function updateModalCourse(id) {
    fetch(`/api/v1/editor/courses/${id}`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    let struct_list = document.getElementById("struct_list");
                    let modal = document.getElementById("modal");
                    if (modal !== null) {
                        struct_list.removeChild(modal);
                    }

                    let modal_elem = document.createElement("div");
                    modal_elem.className = "modal fade";
                    modal_elem.id = "modal";
                    modal_elem.setAttribute("data-bs-backdrop", "static");
                    modal_elem.setAttribute("data-bs-keyboard", "false");
                    modal_elem.setAttribute("tabindex", "-1");
                    modal_elem.setAttribute("aria-labelledby", "modalLabel");
                    modal_elem.setAttribute("aria-hidden", "true");
                    modal_elem.innerHTML = `<div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h1 class="modal-title fs-5" id="exampleModalLabel">Редактирование курса</h1>
                            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Закрыть"></button>
                        </div>
                        <div class="modal-body">
                            <form novalidate>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <select class="form-select form-select-sm validation" aria-label="Выбор направления" disabled readonly>
                                            <option value="${info.direction_id}" selected>${info.direction_name}</option>
                                        </select>
                                        <div id="select_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_name" class="col-form-label-sm">Название курса*</label>
                                        <input type="text" class="form-control form-control-sm validation" id="input_name" value="${info.name === null ? "" : info.name}">
                                        <div id="name_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_first_date" class="col-form-label-sm">Дата начала*</label>
                                        <input type="date" class="form-control form-control-sm validation" id="input_first_date" value="${info.first_class_date === null ? "" : info.first_class_date.slice(0, 10)}">
                                        <div id="first_date_help" class="form-text text-danger"></div>
                                    </div>
                                    <div class="col mb-1">
                                        <label for="input_last_date" class="col-form-label-sm">Дата окончания*</label>
                                        <input type="date" class="form-control form-control-sm validation" id="input_last_date" value="${info.last_class_date === null ? "" : info.last_class_date.slice(0, 10)}">
                                        <div id="last_date_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_num" class="col-form-label-sm">Кол-во занятий*</label>
                                        <input type="number" class="form-control form-control-sm validation" id="input_num" value="${info.num_of_classes === null ? "" : info.num_of_classes}">
                                        <div id="num_help" class="form-text text-danger"></div>
                                    </div>
                                    <div class="col mb-1">
                                        <label for="input_class_time" class="col-form-label-sm">Длительность занятия* (минут)</label>
                                        <input type="number" class="form-control form-control-sm validation" id="input_class_time" value="${info.class_time === null ? "" : info.class_time}">
                                        <div id="class_time_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_week_days" class="col-form-label-sm">Дни недели* (через запятую)</label>
                                        <input type="text" class="form-control form-control-sm validation" id="input_week_days" value="${info.week_days === null ? "" : info.week_days}">
                                        <div id="week_days_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_price" class="col-form-label-sm">Цена*</label>
                                        <input type="number" class="form-control form-control-sm validation" id="input_price" value="${info.price === null ? "" : info.price}">
                                        <div id="price_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_info" class="col-form-label-sm">Информация*</label>
                                        <textarea class="form-control form-control-sm validation" id="input_info" rows="3">${info.info === null ? "" : info.info}</textarea>
                                        <div id="info_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                            </form>
                        </div>
                        <div class="modal-footer">
                            <button id="button_close_modal" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Закрыть</button>
                            <button type="button" class="btn btn-primary" onclick="saveUpdateModalCourse(${id})">Сохранить</button>
                        </div>
                        </div>
                    </div>`
                    struct_list.appendChild(modal_elem);

                    let modalToShow = new bootstrap.Modal(document.getElementById('modal'));
                    modalToShow.show();
                });
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                window.location.replace("/editor/login");
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function saveUpdateModalCourse(id) {
    let elems = document.querySelectorAll(".validation");
    for (let elem of elems) {
        let ind = elem.className.indexOf("border-danger");
        if (ind !== -1) {
            elem.className = elem.className.substring(0, ind - 1);
            let danger_text = elem.nextElementSibling;
            danger_text.innerText = "";
        }
    }

    let is_validated = true;
    for (let elem of elems) {
        if (elem.value === null || elem.value === "") {
            elem.className += " border-danger";
            let danger_text = elem.nextElementSibling;
            danger_text.innerText = "Поле не может быть пустым";
            is_validated = false;
        }
    }

    if (!is_validated) {
        return;
    }

    let course = {
        name: document.getElementById("input_name").value,
        num_of_classes: document.getElementById("input_num").value,
        class_time: document.getElementById("input_class_time").value,
        week_days: document.getElementById("input_week_days").value,
        first_class_date: document.getElementById("input_first_date").value === null ? null : new Date(document.getElementById("input_first_date").value),
        last_class_date: document.getElementById("input_last_date").value === null ? null : new Date(document.getElementById("input_last_date").value),
        price: document.getElementById("input_price").value,
        info: document.getElementById("input_info").value,
        direction_id: Number(document.querySelector("select.form-select").value),
    }

    fetch(`/api/v1/editor/courses/${id}`, {
        method: "PATCH",
        body: JSON.stringify(course)
    })
        .then(response => {
            if (response.status === 200) {
                document.getElementById('button_close_modal').click();
                showSuccessToast("Курс успешно отредактирован");
                courseButton();
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                window.location.replace("/editor/login");
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function searchCourse() {
    let search_string = document.getElementById("search_string").value;

    fetch(`/api/v1/editor/courses?search=${search_string}`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    let elems_table = document.getElementById("elems_table");
                    let table_body = document.querySelector("tbody");
                    if (table_body !== null) {
                        elems_table.removeChild(table_body)
                    }

                    printCourses(elems_table, info);
                })
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                window.location.replace("/editor/login");
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function studentButton() {
    let buttons = document.querySelectorAll(".list-group-item");
    for (let b of buttons) {
        b.className = "list-group-item list-group-item-action";
    }
    let activeButton = document.getElementById("student_button");
    activeButton.className += " active";


    let searchButton = document.getElementById("search_button");
    searchButton.setAttribute("onclick", "searchStudent()");

    let search_string = document.getElementById("search_string");
    search_string.setAttribute("placeholder", "Поиск по фамилии");

    if (document.getElementById("select_course") === null) {
        fetch(`/api/v1/editor/courses`)
            .then(response => {
                if (response.status === 200) {
                    response.json().then(info => {
                        let courses_select_elem = document.getElementById("courses_select");
                        courses_select_elem.innerHTML = `<select id="select_course" class="form-select" aria-label="Курс">
                            <option selected value="-1">&#60;Курс&#62;</option>
                        </select>`;
                        let select_course = document.getElementById("select_course");
                        for (let course of info) {
                            let elem = document.createElement("option");
                            elem.innerText = course.name;
                            elem.setAttribute("value", course.id);
                            select_course.appendChild(elem);
                        }
                    })
                } else if (response.status === 400) {
                    showDangerToast("Проверьте правильность введенных данных", false);
                } else if (response.status === 401) {
                    window.location.replace("/editor/login");
                } else if (response.status === 500) {
                    showDangerToast("Серверная ошибка, попробуйте позже", true);
                }
            })
    }

    let elems_table = document.getElementById("elems_table");

    fetch(`/api/v1/editor/students`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    printStudents(elems_table, info);
                })
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                window.location.replace("/editor/login");
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function printStudents(elems_table, info) {
    let table_head = document.querySelector("thead");
    if (table_head !== null) {
        elems_table.removeChild(table_head);
    }

    let table_body = document.querySelector("tbody");
    if (table_body !== null) {
        elems_table.removeChild(table_body);
    }

    if (info !== null){
        info.sort((x, y) => x.id - y.id);

        let table_head_elem = document.createElement("thead");
        table_head_elem.innerHTML = `<tr>
            <th scope="col">#</th>
            <th scope="col">ID</th>
            <th scope="col">Фамилия</th>
            <th scope="col">Имя</th>
            <th scope="col">Отчество</th>
            <th scope="col">Электронная почта</th>
            <th scope="col">Номер телефона</th>
            <th scope="col">Оплата</th>
            <th scope="col">Дата оплаты</th>
            <th scope="col">Курс</th>
            <th scope="col">Комментарий</th>
            <th scope="col" colspan="2"></th>
        </tr>`;

        elems_table.appendChild(table_head_elem);
        let no_records_elem = document.getElementById("no_records");
        if (no_records_elem !== null) {
            let struct_list = document.getElementById("struct_list");
            struct_list.removeChild(no_records_elem);
        }

        let table_body_elem = document.createElement("tbody");
        for (let i = 0; i < info.length; i++) {
            let elem = document.createElement("tr");
            elem.innerHTML = `<th scope="col">${i + 1}</th>
                <td>${info[i].id}</td>
                <td>${info[i].surname === null ? "" : info[i].surname}</td>
                <td>${info[i].name === null ? "" : info[i].name}</td>
                <td>${info[i].patronymic === null ? "" : info[i].patronymic}</td>
                <td>${info[i].email === null ? "" : info[i].email}</td>
                <td>${info[i].phone === null ? "" : info[i].phone}</td>
                <td>${info[i].payment === null || info[i].payment === false ? "Нет" : "Да"}</td>
                <td>${info[i].date_of_payment === null ? "" : info[i].date_of_payment.slice(8, 10) + "." + info[i].date_of_payment.slice(5, 7) + "." + info[i].date_of_payment.slice(0, 4)}</td>
                <td>${info[i].course_name === null ? "" : info[i].course_name}</td>
                <td>${info[i].comment === null ? "" : info[i].comment}</td>
                <td width="1%"><button class="btn btn-primary btn-sm" type="button" onclick="updateModalStudent(${info[i].id})">Редактировать</button></td>
                <td width="1%"><button class="btn btn-danger btn-sm" type="button" onclick="deleteStudent(${info[i].id})">Удалить</button></td>`;
            table_body_elem.appendChild(elem);
        }
        elems_table.appendChild(table_body_elem);
    } else {
        let no_records_elem = document.getElementById("no_records");
        if (no_records_elem === null) {
            let struct_list = document.getElementById("struct_list");
            let elem = document.createElement("div");
            elem.className = "row justify-content-center text-center";
            elem.setAttribute("id", "no_records");
            elem.innerHTML = "<p>Соответствующие студенты не были найдены</p>";
            struct_list.appendChild(elem);
        }
    }
}

function updateModalStudent(id) {
    fetch(`/api/v1/editor/students/${id}`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    let struct_list = document.getElementById("struct_list");
                    let modal = document.getElementById("modal");
                    if (modal !== null) {
                        struct_list.removeChild(modal);
                    }

                    let modal_elem = document.createElement("div");
                    modal_elem.className = "modal fade";
                    modal_elem.id = "modal";
                    modal_elem.setAttribute("data-bs-backdrop", "static");
                    modal_elem.setAttribute("data-bs-keyboard", "false");
                    modal_elem.setAttribute("tabindex", "-1");
                    modal_elem.setAttribute("aria-labelledby", "modalLabel");
                    modal_elem.setAttribute("aria-hidden", "true");
                    modal_elem.innerHTML = `<div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h1 class="modal-title fs-5" id="exampleModalLabel">Редактирование студента</h1>
                            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Закрыть"></button>
                        </div>
                        <div class="modal-body">
                            <form novalidate>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_surname" class="col-form-label-sm">Фамилия*</label>
                                        <input type="text" class="form-control form-control-sm validation" id="input_surname" value="${info.surname === null ? "" : info.surname}">
                                        <div id="surname_help" class="form-text text-danger"></div>
                                    </div>
                                    <div class="col mb-1">
                                        <label for="input_name" class="col-form-label-sm">Имя*</label>
                                        <input type="text" class="form-control form-control-sm validation" id="input_name" value="${info.name === null ? "" : info.name}">
                                        <div id="name_help" class="form-text text-danger"></div>
                                    </div>
                                    <div class="col mb-1">
                                        <label for="input_patronymic" class="col-form-label-sm">Отчество*</label>
                                        <input type="text" class="form-control form-control-sm validation" id="input_patronymic" value="${info.patronymic === null ? "" : info.patronymic}">
                                        <div id="patronymic_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_email" class="col-form-label-sm">Электронная почта*</label>
                                        <input type="email" class="form-control form-control-sm validation" id="input_email" placeholder="example@mail.com" value="${info.email === null ? "" : info.email}">
                                        <div id="email_help" class="form-text text-danger"></div>
                                    </div>
                                    <div class="col mb-1">
                                        <label for="input_phone" class="col-form-label-sm">Телефон*</label>
                                        <input type="text" class="form-control form-control-sm validation" id="input_phone" placeholder="+79000000000" value="${info.phone === null ? "" : info.phone}">
                                        <div id="phone_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1">
                                        <label for="input_comment" class="col-form-label-sm">Комментарий</label>
                                        <textarea class="form-control form-control-sm" id="input_comment" rows="3">${info.comment === null ? "" : info.comment}</textarea>
                                        <div id="comment_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                                <div class="row justify-content-center">
                                    <div class="col mb-1 d-flex align-items-center justify-content-center">
                                        <label for="input_payment" class="form-check-label">Оплата</label>
                                        <input class="form-check-input" type="checkbox" id="input_payment" ${info.payment === null || info.payment === false ? "" : "checked" }>
                                        <div id="payment_help" class="form-text text-danger"></div>
                                    </div>
                                    <div class="col mb-1">
                                        <label for="input_date_payment" class="col-form-label-sm">Дата оплаты</label>
                                        <input class="form-control form-control-sm" type="date" id="input_date_payment" value="${info.date_of_payment === null ? "" : info.date_of_payment.slice(0, 10)}">
                                        <div id="date_payment_help" class="form-text text-danger"></div>
                                    </div>
                                </div>
                            </form>
                        </div>
                        <div class="modal-footer">
                            <button id="button_close_modal" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Закрыть</button>
                            <button type="button" class="btn btn-primary" onclick="saveUpdateModalStudent(${id})">Сохранить</button>
                        </div>
                        </div>
                    </div>`
                    struct_list.appendChild(modal_elem);
                    let modalToShow = new bootstrap.Modal(document.getElementById('modal'));
                    modalToShow.show();
                });
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function saveUpdateModalStudent(id) {
    let elems = document.querySelectorAll(".validation");
    for (let elem of elems) {
        let ind = elem.className.indexOf("border-danger");
        if (ind !== -1) {
            elem.className = elem.className.substring(0, ind - 1);
            let danger_text = elem.nextElementSibling;
            danger_text.innerText = "";
        }
    }

    let is_validated = true;
    for (let elem of elems) {
        if (elem.value === null || elem.value === "") {
            elem.className += " border-danger";
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
        comment: document.getElementById("input_comment").value,
        payment: document.getElementById("input_payment").checked,
        date_of_payment: document.getElementById("input_date_payment").value === null ? null : new Date(document.getElementById("input_date_payment").value)
    }

    fetch(`/api/v1/editor/students/${id}`, {
        method: "PATCH",
        body: JSON.stringify(student)
    })
        .then(response => {
            if (response.status === 200) {
                document.getElementById('button_close_modal').click();
                showSuccessToast("Студент успешно отредактирован");
                studentButton();
            } else if (response.status === 400) {
                response.text().then(err => {
                    if (err.includes("mail")) {
                        let elem = document.getElementById("input_email");
                        elem.className += " border-danger";
                        let danger_text = elem.nextElementSibling;
                        danger_text.innerText = "Неверный формат электронной почты";
                    } else if (err.includes("phone")) {
                        let elem = document.getElementById("input_phone");
                        elem.className += " border-danger";
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

function searchStudent() {
    let search_string = document.getElementById("search_string").value;
    let course = document.getElementById("select_course").value;

    fetch(`/api/v1/editor/students?search=${search_string}&course=${course}`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    let elems_table = document.getElementById("elems_table");
                    let table_body = document.querySelector("tbody");
                    if (table_body !== null) {
                        elems_table.removeChild(table_body)
                    }

                    printStudents(elems_table, info);
                })
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                window.location.replace("/editor/login");
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function deleteStudent(id) {
    fetch(`/api/v1/editor/students/${id}`, {
        method: "DELETE"
    })
        .then(response => {
            if (response.status === 200) {
                showSuccessToast("Студент был успешно удален");
                studentButton();
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                window.location.replace("/editor/login");
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

function showSuccessToast(message) {
    let toast_div = document.querySelector("div.toast-container");
    let elem = document.createElement("div");
    elem.className = "toast align-items-center text-bg-success";
    let btn_style = "btn-close-white";
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

