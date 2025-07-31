document.addEventListener('DOMContentLoaded', () => {
    const createForm = document.getElementById('create-form');
    const contextList = document.getElementById('context-list');
    const updateModal = document.getElementById('update-modal');
    const updateForm = document.getElementById('update-form');
    const closeModal = document.querySelector('.close-button');

    const API_URL = '/mcp'; // Use relative URL since the API is on the same server

    // Fetch and display all contexts
    const fetchContexts = async () => {
        try {
            const response = await fetch(API_URL);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const contexts = await response.json();
            contextList.innerHTML = ''; // Clear the list
            if (contexts) {
                contexts.forEach(context => {
                    const li = document.createElement('li');
                    li.innerHTML = `
                        <div class="context-data">
                            <strong>ID:</strong> ${context.id}<br>
                            <strong>Created At:</strong> ${new Date(context.created_at).toLocaleString()}<br>
                            <strong>Data:</strong> <pre>${JSON.stringify(context.data, null, 2)}</pre>
                        </div>
                        <div class="context-actions">
                            <button class="update-btn" data-id="${context.id}" data-data='${JSON.stringify(context.data)}'>Update</button>
                            <button class="delete-btn" data-id="${context.id}">Delete</button>
                        </div>
                    `;
                    contextList.appendChild(li);
                });
            }
        } catch (error) {
            console.error('Error fetching contexts:', error);
            contextList.innerHTML = '<li>Error loading contexts. Is the server running?</li>';
        }
    };

    // Create a new context
    createForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const dataText = document.getElementById('create-data').value;
        try {
            const data = JSON.parse(dataText);
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            });
            if (response.ok) {
                fetchContexts(); // Refresh the list
                createForm.reset();
            } else {
                const errorText = await response.text();
                alert(`Error creating context: ${errorText}`);
            }
        } catch (error) {
            alert('Invalid JSON data. Please make sure it is a valid JSON object.');
            console.error('Error creating context:', error);
        }
    });

    // Handle clicks on update and delete buttons
    contextList.addEventListener('click', async (e) => {
        const target = e.target;
        const id = target.dataset.id;

        if (target.classList.contains('delete-btn')) {
            if (confirm('Are you sure you want to delete this context?')) {
                try {
                    const response = await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
                    if (response.ok) {
                        fetchContexts(); // Refresh the list
                    } else {
                        const errorText = await response.text();
                        alert(`Error deleting context: ${errorText}`);
                    }
                } catch (error) {
                    console.error('Error deleting context:', error);
                }
            }
        }

        if (target.classList.contains('update-btn')) {
            const data = JSON.parse(target.dataset.data);
            document.getElementById('update-id').value = id;
            document.getElementById('update-data').value = JSON.stringify(data, null, 2);
            updateModal.style.display = 'block';
        }
    });

    // Close the update modal
    closeModal.addEventListener('click', () => {
        updateModal.style.display = 'none';
    });

    window.addEventListener('click', (e) => {
        if (e.target == updateModal) {
            updateModal.style.display = 'none';
        }
    });

    // Update a context
    updateForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const id = document.getElementById('update-id').value;
        const dataText = document.getElementById('update-data').value;
        try {
            const data = JSON.parse(dataText);
            const response = await fetch(`${API_URL}/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(data)
            });
            if (response.ok) {
                fetchContexts(); // Refresh the list
                updateModal.style.display = 'none';
            } else {
                const errorText = await response.text();
                alert(`Error updating context: ${errorText}`);
            }
        } catch (error) {
            alert('Invalid JSON data. Please make sure it is a valid JSON object.');
            console.error('Error updating context:', error);
        }
    });

    // Initial fetch
    fetchContexts();
});
