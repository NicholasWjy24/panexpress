const API_BASE_URL = 'http://localhost:8080/api';
const token = localStorage.getItem('token');
const storedUser = JSON.parse(localStorage.getItem('user') || 'null');

const state = {
  users: [],
  error: '',
};

let elements = {};

if (!token) {
  window.location.href = '../login/login.html';
}

async function initUserPage() {
  await loadLayout(
    'user',
    'Users Table',
    'View user rows from the users API.'
  );

  elements = {
    adminChip: document.getElementById('adminChip'),
    drawerUser: document.getElementById('drawerUser'),
    tableHead: document.getElementById('tableHead'),
    tableBody: document.getElementById('tableBody'),
    emptyState: document.getElementById('emptyState'),
    errorText: document.getElementById('errorText'),
    totalUsers: document.getElementById('totalUsers'),

    userModal: document.getElementById('userModal'),
    userForm: document.getElementById('userForm'),
    saveUserButton: document.getElementById('saveUserButton'),
  };

  if (storedUser) {
    elements.adminChip.textContent =
      storedUser.username || 'Admin';

    elements.drawerUser.textContent =
      `Signed in as ${storedUser.username || 'Admin'}`;
  }

  document
    .getElementById('logoutButton')
    .addEventListener('click', logout);

  document
    .getElementById('addButton')
    .addEventListener('click', openUserModal);

  document
    .getElementById('refreshButton')
    .addEventListener('click', loadUsers);

  document
    .getElementById('closeUserModal')
    .addEventListener('click', closeUserModal);

  elements.userForm.addEventListener(
    'submit',
    createUser
  );

  await loadUsers();
}

function logout() {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
  window.location.href = '../login/login.html';
}

function openUserModal() {
  elements.userModal.hidden = false;
}

function closeUserModal() {
  elements.userModal.hidden = true;
  elements.userForm.reset();
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
    window.location.href = '../login/login.html';
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

async function createUser(event) {
  event.preventDefault();

  const username = document.getElementById('newUsername').value.trim();
  const email = document.getElementById('newEmail').value.trim();
  const password = document.getElementById('newPassword').value;
  const roleLevel = Number(document.getElementById('newRoleLevel').value);

  if (!Number.isInteger(roleLevel) || roleLevel < 1 || roleLevel > 3) {
    alert("Role level must be an integer between 1 and 3.");
    return;
  }

  try {
    elements.saveUserButton.disabled = true;
    elements.saveUserButton.textContent = 'Saving...';

    const registerResponse = await fetch(`${API_BASE_URL}/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username,
        email,
        password,
      }),
    });

    const registerData = await registerResponse.json();

    if (!registerResponse.ok) {
      throw new Error(registerData.error || 'Failed to create user');
    }

    const createdUserId = registerData.user_id;

    if (createdUserId && roleLevel !== 3) {
      const roleResponse = await fetch(`${API_BASE_URL}/users/${createdUserId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          username,
          email,
          role_level: roleLevel,
        }),
      });

      if (!roleResponse.ok) {
        throw new Error('User was created, but failed to update role level');
      }
    }

    closeUserModal();
    await loadUsers();
    alert(registerData.message || 'User created successfully');
  } catch (error) {
    console.error(error);
    alert(error.message);
  } finally {
    elements.saveUserButton.disabled = false;
    elements.saveUserButton.textContent = 'Save User';
  }
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
initUserPage();
