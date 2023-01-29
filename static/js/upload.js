
function uploadFile() {
    const title = document.getElementById('title').value;
    const selectedFile = document.getElementById('file').files[0];

    const data = new FormData()
    data.append('title', title)
    data.append('file', selectedFile)
    
    fetch('/v0.1/documents', {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: data
    }).then(response => response.text())
    .then(raw => {
        const data = JSON.parse(raw);
        window.location.href = '/' + data.document_id;
    });
}