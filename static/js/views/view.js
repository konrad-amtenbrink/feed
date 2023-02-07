function onLoad() {
    const fileName = window.location.href.split('/').pop();
    fetch(`/v0.1/document?id=` + fileName, {
            method: 'GET',
        })
        .then(response => response.blob())
        .then(async blob => {
            const content = document.createElement('zero-md');
            const script = document.createElement('script');
            script.type = 'text/markdown';
            script.innerHTML = await blob.text();
            content.appendChild(script);
            document.body.appendChild(content);
        });
}