function onLoad() {
    const fileName = window.location.href.split('/').pop();
    fetch(`/v0.1/documents`, {
            method: 'GET',
        })
        .then(async response => {
            const documents = await response.json();

            documents.forEach(doc => {
                const linkElement = document.createElement('a');
                linkElement.className = 'lock max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow-md hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700';
                linkElement.href = `/${doc.id}`;

                const titleElement = document.createElement('h5');
                titleElement.className = 'mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white';
                titleElement.innerHTML = doc.title;

                const descriptionElement = document.createElement('p');
                descriptionElement.className = 'font-normal text-gray-700 dark:text-gray-400';
                descriptionElement.innerHTML = "Here are the biggest enterprise technology acquisitions of 2021 so far, in reverse chronological order.";

                linkElement.appendChild(titleElement);
                linkElement.appendChild(descriptionElement);

                document.getElementById('files').appendChild(linkElement);
            });
        }); 
}
