function onLoad() {
    fetch(`/v0.1/admin/documents`, {
            method: 'GET',
        })
        .then(async response => {
            if (response.status !== 200) {
                window.location.href = '/';
            }
            const documents = await response.json();

            documents.forEach(doc => {
                const row = document.createElement('tr');
                row.className = 'bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600';

                const title = document.createElement('th');
                title.scope = 'row';
                title.className = 'px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white';
                title.innerHTML = doc.title;

                const topics = document.createElement('td');
                topics.className = 'px-6 py-4';
                topics.innerHTML = doc.id;

                const date = document.createElement('td');
                date.className = 'px-6 py-4';
                date.innerHTML = doc.created_at;

                const action = document.createElement('td');
                action.className = 'px-6 py-4 text-right';

                const actionLink = document.createElement('button');
                actionLink.className = 'text-indigo-600 hover:text-indigo-900';
                actionLink.innerHTML = 'Delete';
                actionLink.addEventListener("click", onDelete.bind(this, doc));

                action.appendChild(actionLink);

                row.appendChild(title);
                row.appendChild(topics);
                row.appendChild(date);
                row.appendChild(action);

                document.getElementById('files').appendChild(row);
            });
        }); 
}

function onDelete(doc) {
    fetch(`/v0.1/document?id=` + doc.id, {
            method: 'DELETE',
        })
        .then(() => {
            document.getElementById('files').innerHTML = '';
            onLoad();
        })
}