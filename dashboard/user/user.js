const API_BASE_URL = 'http://localhost:8080/api';
const token = localStorage.getItem('token');
const storedUser = JSON.parse(localStorage.getItem('user') || 'null');

const state = {
  users: [],
  error: '',
};

const elements = {
  adminChip: document.getElementById('adminChip'),
  drawerUser: document.getElementById('drawerUser'),
  tableHead: document.getElementById('tableHead'),
  tableBody: document.getElementById('tableBody'),
  emptyState: document.getElementById('emptyState'),
  errorText: document.getElementById('errorText'),
  totalUsers: document.getElementById('totalUsers'),
};

if (!token) {
  window.location.href = '../login/login.html';
}

if (storedUser) {
  elements.adminChip.textContent = storedUser.username || 'Admin';
  elements.drawerUser.textContent = `Signed in as ${storedUser.username || 'Admin'}`;
}

document.getElementById('menuToggle').addEventListener('click', () => {
  document.body.classList.add('drawer-open');
});

document.getElementById('drawerBackdrop').addEventListener('click', closeDrawer);

document.getElementById('logoutButton').addEventListener('click', () => {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  window.location.href = '../login/login.html';
});

document.getElementById('refreshButton').addEventListener('click', loadUsers);

function closeDrawer() {
  document.body.classList.remove('drawer-open');
}

function getArrayPayload(data, key) {
  if (Array.isArray(data)) {
    return data;
  }

  if (data && Array.isArray(data[key])) {
    return data[key];
  }

  if (data && Array.isArray(data.data)) {
    return data.data;
  }

  return [];
}

function setError(message) {
  state.error = message || '';
  elements.errorText.textContent = state.error;
  elements.errorText.hidden = !state.error;
}

function renderUsers() {
  elements.totalUsers.textContent = state.users.length;

  elements.tableHead.innerHTML = `
    <tr>
      <th>User ID</th>
      <th>Username</th>
      <th>Email</th>
      <th>Role Level</th>
      <th>CRUD</th>
    </tr>
  `;

  elements.tableBody.innerHTML = state.users.map((user) => `
    <tr>
      <td>${user.id}</td>
      <td>${user.username}</td>
      <td>${user.email}</td>
      <td>${user.roleLevel}</td>
      <td>
        <div class="action-group">
          <button class="action-button edit-button" onclick="editUser(${user.id})">
            Edit
          </button>
          <button class="action-button delete-button" onclick="deleteUser(${user.id})">
            Delete
          </button>
        </div>
      </td>
    </tr>
  `).join('');

  elements.emptyState.hidden = state.users.length > 0;
}

async function authorizedGet(path) {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (response.status === 401) {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    window.location.href = "../login/login.html";
    return null;
  }

  if (!response.ok) {
    throw new Error(`GET /api${path} failed with status ${response.status}`);
  }

  return response.json();
}

async function loadUsers() {
  setError('');

  try {
    const data = await authorizedGet('/users');
    state.users = getArrayPayload(data, 'users').map((user) => ({
      id: user.id,
      username: user.username,
      email: user.email,
      roleLevel: user.role_level,
    }));
  } catch (error) {
    setError(error.message);
  }

  renderUsers();
}

async function deleteUser(id) {
  const confirmed = confirm('Delete this user?');

  if (!confirmed) {
    return;
  }

  try {
    const response = await fetch(`${API_BASE_URL}/users/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to delete user');
    }

    await loadUsers();
    alert('User deleted successfully');
  } catch (error) {
    console.error(error);
    alert(error.message);
  }
}

async function editUser(id) {
  const user = state.users.find((currentUser) => currentUser.id === id);

  if (!user) {
    alert('User not found');
    return;
  }

  const username = prompt('Edit username:', user.username);
  if (!username) {
    return;
  }

  const email = prompt('Edit email:', user.email);
  if (!email) {
    return;
  }

  const roleLevel = prompt('Edit role level:', user.roleLevel);
  if (!roleLevel) {
    return;
  }

  try {
    const response = await fetch(`${API_BASE_URL}/users/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        username,
        email,
        role_level: Number(roleLevel),
      }),
    });

    if (!response.ok) {
      throw new Error('Failed to update user');
    }

    await loadUsers();
    alert('User updated successfully');
  } catch (error) {
    console.error(error);
    alert(error.message);
  }
}

loadUsers();
