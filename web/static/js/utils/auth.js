function logout() {
    fetch('/logout', {
        method: 'POST',
    })
    .then((response) => {
        if (response.status === 200) {
            window.location.href = '/login';
        }
    });
}