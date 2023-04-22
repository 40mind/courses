const urlParams = new URLSearchParams(window.location.search);
let student_id = urlParams.get('student');

fetch(`/api/v1/payment/confirm/${student_id}`, {
    method: "POST"
}).then(response => {
    if (response.status === 200) {
        let column = document.querySelector("div.col-5.text-center");
        column.innerHTML = `<h3>Оплачено успешно!</h3>
                            <p>Информация о курсе поступит на вашу электронную почту. Спасибо за покупку!</p>`
    } else if (response.status === 400) {
        showDangerToast("Курс не был оплачен, сначала оплатите курс", false);
    } else if (response.status === 500) {
        showDangerToast("Серверная ошибка, попробуйте позже", true);
    }
});

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