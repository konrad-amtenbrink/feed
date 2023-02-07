function onLoad() {
    fetch(`/v0.1/documents`, {
            method: 'GET',
        })
        .then(async response => {
            const [latest, ...documents] = (await response.json()).splice(0, 5);
            createListElement(latest, true);
            documents.forEach(doc => {
                createListElement(doc);
            });
        }); 
}

function createListElement(doc, isLatest = false) {
    const icon = document.createElement('span');
    icon.className = 'absolute flex items-center justify-center w-6 h-6 bg-blue-100 rounded-full -left-3 ring-8 ring-white dark:ring-gray-900 dark:bg-blue-900';

    const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    svg.setAttribute('aria-hidden', 'true');
    svg.setAttribute('class', 'w-3 h-3 text-blue-800 dark:text-blue-300');
    svg.setAttribute('fill', 'currentColor');
    svg.setAttribute('viewBox', '0 0 20 20');
    svg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');

    const path = document.createElementNS('http://www.w3.org/2000/svg', 'path');
    path.setAttribute('fill-rule', 'evenodd');
    path.setAttribute('d', 'M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z');
    path.setAttribute('clip-rule', 'evenodd');

    svg.appendChild(path);
    icon.appendChild(svg);

    const title = document.createElement('h3');
    title.className = 'mb-1 text-lg font-semibold text-gray-900 dark:text-white';
    title.innerHTML = doc.title;

    if (isLatest) {
        const latest = document.createElement('span');
        latest.className = 'bg-blue-100 text-blue-800 text-sm font-medium mr-2 px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300 ml-3';
        latest.innerHTML = 'Latest';
        title.appendChild(latest);
    }

    const time = document.createElement('time');
    time.className = 'block mb-2 text-sm font-normal leading-none text-gray-400 dark:text-gray-500';
    time.innerHTML = doc.created_at;

    const description = document.createElement('p');
    description.className = 'text-base font-normal text-gray-500 dark:text-gray-400';
    description.innerHTML = 'All of the changes in the latest release of Flowbite Figma.';

    const listElement = document.createElement('li');
    listElement.className = 'mb-10 ml-6';
    listElement.appendChild(icon);
    listElement.appendChild(title);
    listElement.appendChild(time);
    listElement.appendChild(description);

    document.getElementById('timeline').appendChild(listElement);
}