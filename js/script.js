document.addEventListener('DOMContentLoaded', function() {
    fetchUpdates();

    const form = document.getElementById('update-form');
    form.addEventListener('submit', function(event) {
        event.preventDefault();
        
        const formData = new FormData(form);
        const updateData = {
            userId: formData.get('user-id'),
            update: formData.get('update'),
            blockers: formData.get('blockers')
        };

        submitUpdate(updateData);
        form.reset();
    });
});

function fetchUpdates() {
    fetch('/updates')
    .then(response => response.json())
    .then(updates => {
        const updatesContainer = document.getElementById('updates-container');
        updatesContainer.innerHTML = '';
        updates.forEach(update => {
            const updateElement = document.createElement('div');
            updateElement.innerHTML = `
                <div>ID: ${update.id}</div>
                <div>User ID: ${update.userId}</div>
                <div>Update: ${update.update}</div>
                <div>Blockers: ${update.blockers}</div>
                <div>Created At: ${update.createdAt}</div>
                <hr>
            `;
            updatesContainer.appendChild(updateElement);
        });
    })
    .catch(error => console.error('Error fetching updates:', error));
}

function submitUpdate(updateData) {
    fetch('/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(updateData)
    })
    .then(fetchUpdates)
    .catch(error => console.error('Error submitting update:', error));
}
