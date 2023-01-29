function onLoad() {
    const fileName = window.location.href.split('/').pop();
    fetch(`/v0.1/documents`, {
            method: 'GET',
        })
        .then(async response => {
            const documents = await response.json();

            documents.forEach(doc => {
                const listElement = document.createElement('li');
                listElement.className = 'dark:text-gray-300';
                const linkElement = document.createElement('a');
                linkElement.href = `/${doc.id}`;
                linkElement.innerHTML = doc.title;
                listElement.appendChild(linkElement);
                document.getElementById('files').appendChild(listElement);
            });
        }); 
}