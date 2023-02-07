function onLoad() {
    const fileName = window.location.href.split('/').pop();
    fetch(`/v0.1/document?id=` + fileName, {
            method: 'GET',
        })
        .then(response => {
            if (response.status !== 200) {
                window.location.href = '/login';
            }
            return response.blob();
        })
        .then(async blob => {
            const content = document.createElement('zero-md');
            const script = document.createElement('script');
            script.type = 'text/markdown';
            script.innerHTML = await blob.text();
            content.appendChild(script);
            document.body.appendChild(content);
        });
}