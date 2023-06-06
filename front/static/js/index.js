let admin_buttons = document.getElementById("admin_buttons");
if (getCookie("admin-session") !== undefined) {
    admin_buttons.innerHTML = `<div class="btn-group-vertical border-0" role="group" aria-label="Группа вертикальных кнопок">
        <button type="button" class="btn btn-outline-primary" onclick="location.href='/admin/panel'">Панель администратора</button>
        <button type="button" class="btn btn-outline-primary" onclick="logout()">Выход</button>
    </div>`
} else if (getCookie("editor-session") !== undefined) {
    admin_buttons.innerHTML = `<div class="btn-group-vertical border-0" role="group" aria-label="Группа вертикальных кнопок">
        <button type="button" class="btn btn-outline-primary" onclick="location.href='/editor/panel'">Панель редактора</button>
        <button type="button" class="btn btn-outline-primary" onclick="logout()">Выход</button>
    </div>`
} else {
    admin_buttons.innerHTML = `<div class="btn-group-vertical border-0" role="group" aria-label="Группа вертикальных кнопок">
        <button type="button" class="btn btn-outline-primary" onclick="location.href='/admin/login'">Вход для администратора</button>
        <button type="button" class="btn btn-outline-primary" onclick="location.href='/editor/login'">Вход для редактора</button>
    </div>`
}

fetch("/api/v1")
    .then(response => {
        if (response.status === 200) {
            response.json().then(info => {
                let direction_select = document.getElementById("select_direction");
                if (info.directions !== null) {
                    for (let direction of info.directions) {
                        let elem = document.createElement("option");
                        elem.innerText = direction.name;
                        elem.setAttribute("value", direction.id);
                        direction_select.appendChild(elem);
                    }
                }

                let courses_row = document.getElementById("courses_row");
                printCourses(courses_row, info);
            })
        } else if (response.status === 400) {
            showDangerToast("Проверьте правильность введенных данных", false);
        } else if (response.status === 500) {
            showDangerToast("Серверная ошибка, попробуйте позже", true);
        }
    });

function searchButton() {
    let direction = document.getElementById("select_direction").value;
    let search_string = document.getElementById("search_string").value;
    fetch(`/api/v1?direction=${direction}&search=${search_string}`)
        .then(response => {
            if (response.status === 200) {
                response.json().then(info => {
                    let courses_row = document.getElementById("courses_row");
                    let remove_rows = document.querySelectorAll("div.col.course");
                    for (let remove_row of remove_rows) {
                        courses_row.removeChild(remove_row);
                    }

                    printCourses(courses_row, info);
                })
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        });
}

function printCourses(courses_row, info) {
    if (info.courses !== null){
        let no_courses_elem = document.getElementById("no_courses_found");
        if (no_courses_elem !== null) {
            document.body.removeChild(no_courses_elem);
        }

        for (let course of info.courses) {
            let elem = document.createElement("div");
            elem.className = "col course";
            elem.innerHTML = `<div class="card h-100">
                        <div class="card-body">
                            <h5 class="card-title">${course.name}</h5>
                            <ul class="list-group list-group-flush">
                                <li class="list-group-item">Направление: ${course.direction_name}</li>
                                <li class="list-group-item">Цена: ${course.price}</li>
                            </ul>
                        </div>
                        <a class="btn btn-primary rounded-top-0" href="/course/${course.id}" role="button">Подробнее</a>
                    </div>`;
            courses_row.appendChild(elem);
        }
    } else {
        let no_courses_elem = document.getElementById("no_courses_found");
        if (no_courses_elem === null) {
            let elem = document.createElement("div");
            elem.className = "container-fluid text-center";
            elem.setAttribute("id", "no_courses_found");
            elem.innerHTML = "<p>Соответствующие курсы не были найдены</p>";
            document.body.appendChild(elem);
        }
    }
}

function logout() {
    if (getCookie("admin-session") !== undefined) {
        fetch(`/api/v1/admin/logout`, {
            method: "POST"
        })
            .then(response => {
                if (response.status === 500) {
                    showDangerToast("Серверная ошибка, попробуйте позже", true);
                    return;
                }

                location.reload();
            });
    }

    if (getCookie("editor-session") !== undefined) {
        fetch(`/api/v1/editor/logout`, {
            method: "POST"
        })
            .then(response => {
                if (response.status === 500) {
                    showDangerToast("Серверная ошибка, попробуйте позже", true);
                    return;
                }

                location.reload();
            });
    }
}

function getCookie(name) {
    let matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
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
