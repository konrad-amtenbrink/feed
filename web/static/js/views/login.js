
function login() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    var data = { username, password };
    fetch('/login', {
        method: 'POST',
        headers: new Headers({'content-type': 'application/json'}),
        body: JSON.stringify(data)
    })
    .then((response) => {
        if (response.status !== 200) {
            window.location.href = '/login';
        }
        window.location.href = '/';
    });
}