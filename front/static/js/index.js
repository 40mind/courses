let admin_buttons = document.getElementById("admin_buttons");
if (getCookie("admin-session") !== undefined) {
    let elemAdminPanel = document.createElement("a");
    elemAdminPanel.className = "nav-link active d-flex";
    elemAdminPanel.setAttribute("aria-current", "page");
    elemAdminPanel.setAttribute("href", "/admin_panel.html");
    elemAdminPanel.innerText = "Панель администратора";
    admin_buttons.appendChild(elemAdminPanel);

    let elemAdminLogout = document.createElement("a");
    elemAdminLogout.className = "nav-link active d-flex";
    elemAdminLogout.setAttribute("aria-current", "page");
    elemAdminLogout.setAttribute("href", "/logout.html");
    elemAdminLogout.innerText = "Выход";
    admin_buttons.appendChild(elemAdminLogout);
} else {
    let elem = document.createElement("a");
    elem.className = "nav-link active d-flex";
    elem.setAttribute("aria-current", "page");
    elem.setAttribute("href", "/login.html");
    elem.innerText = "Вход для администратора";
    admin_buttons.appendChild(elem);
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
            elem.innerHTML = `<div class="card">
                        <div class="card-body">
                            <h5 class="card-title">${course.name}</h5>
                            <ul class="list-group list-group-flush">
                                <li class="list-group-item">Направление: ${course.direction_name}</li>
                                <li class="list-group-item">Цена: ${course.price}</li>
                                <a class="btn btn-primary" href="/course_detail.html?course=${course.id}" role="button">Подробнее</a>
                            </ul>
                        </div>
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
