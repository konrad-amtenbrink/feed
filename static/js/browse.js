function onLoad() {
    const fileName = window.location.href.split('/').pop();
    fetch(`/v0.1/documents`, {
            method: 'GET',
        })
        .then(async response => {
            const documents = await response.json();

            documents.forEach(doc => {
                console.log(doc)
                const listElement = document.createElement('li');
                listElement.className = 'pb-3 sm:pb-4';

                const mainContainer = document.createElement('div');
                mainContainer.className = 'flex items-center space-x-4';

                const infoContainer = document.createElement('div');
                infoContainer.className = 'flex-1 min-w-0';

                const linkElement = document.createElement('a');
                linkElement.className = 'text-sm font-medium text-gray-900 truncate dark:text-white';
                linkElement.href = `/${doc.id}`;
                linkElement.innerHTML = doc.title;

                infoContainer.appendChild(linkElement);

                const tagElement = document.createElement('p');
                tagElement.className = 'text-sm text-gray-500 truncate dark:text-gray-400';
                tagElement.innerHTML = '#tag1 #tag2';
                infoContainer.appendChild(tagElement);

                mainContainer.appendChild(infoContainer);

                
                const priceElement = document.createElement('div');
                priceElement.className = 'inline-flex items-center text-base font-semibold text-gray-900 dark:text-white';
                priceElement.innerHTML = '03-20-2021';

                mainContainer.appendChild(priceElement);

                listElement.appendChild(mainContainer);
                
                document.getElementById('files').appendChild(listElement);
            });
        }); 
}