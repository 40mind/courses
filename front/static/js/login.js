function loginButton() {
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

    let login = document.getElementById("input_login").value;
    let password = document.getElementById("input_password").value;
    let admin = {
        login: login,
        password: password
    };
    fetch(`/api/v1/admin/login`, {
        method: "POST",
        body: JSON.stringify(admin)
    })
        .then(response => {
            if (response.status === 200) {
                window.location.replace("/")
            } else if (response.status === 400) {
                showDangerToast("Проверьте правильность введенных данных", false);
            } else if (response.status === 401) {
                for (let elem of elems) {
                    elem.className = "form-control form-control-sm border-danger";
                }
                let danger_text = document.getElementById("password_help");
                danger_text.innerText = "Неверные логин или пароль";
            } else if (response.status === 500) {
                showDangerToast("Серверная ошибка, попробуйте позже", true);
            }
        })
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
