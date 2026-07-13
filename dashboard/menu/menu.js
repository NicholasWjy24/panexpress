const API_BASE_URL = 'http://localhost:8080/api';
const token = localStorage.getItem('token');
const storedUser = JSON.parse(localStorage.getItem('user') || 'null');

const state = {
  menus: [],
  error: '',
};

let elements = {};

if (!token) {
  window.location.href = '../login/login.html';
}

async function initMenuPage() {
  await loadLayout('menu', 'Menu Table', 'View menu rows from the menu API.');

  elements = {
    adminChip: document.getElementById('adminChip'),
    drawerUser: document.getElementById('drawerUser'),
    tableHead: document.getElementById('tableHead'),
    tableBody: document.getElementById('tableBody'),
    emptyState: document.getElementById('emptyState'),
    errorText: document.getElementById('errorText'),
    totalMenus: document.getElementById('totalMenus'),
    menuModal: document.getElementById('menuModal'),
    menuForm: document.getElementById('menuForm'),
    saveMenuButton: document.getElementById('saveMenuButton'),
  };

  if (storedUser) {
    elements.adminChip.textContent = storedUser.username || 'Admin';
    elements.drawerUser.textContent = `Signed in as ${storedUser.username || 'Admin'}`;
  }

  document.getElementById('logoutButton').addEventListener('click', () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    window.location.href = '../login/login.html';
  });

  document.getElementById('refreshButton').addEventListener('click', loadMenus);

  document.getElementById('addButton').addEventListener('click', () => {
    elements.menuModal.hidden = false;
  });

  document.getElementById('closeModal').addEventListener('click', () => {
    elements.menuModal.hidden = true;
  });

  elements.menuForm.addEventListener('submit', handleMenuSubmit);

  loadMenus();
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

function renderMenus() {
  elements.totalMenus.textContent = state.menus.length;

  elements.tableHead.innerHTML = `
    <tr>
      <th>Menu ID</th>
      <th>Menu Name</th>
      <th>Route</th>
      <th>Min Role</th>
      <th>Max Role</th>
    </tr>
  `;

  elements.tableBody.innerHTML = state.menus.map((menu) => `
    <tr>
      <td>${menu.id}</td>
      <td>${menu.name}</td>
      <td>${menu.route}</td>
      <td>${menu.minRoleLevel}</td>
      <td>${menu.maxRoleLevel}</td>
    </tr>
  `).join('');

  elements.emptyState.hidden = state.menus.length > 0;
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

async function loadMenus() {
  setError('');

  try {
    const data = await authorizedGet('/menus');
    state.menus = getArrayPayload(data, 'menu').map((menu) => ({
      id: menu.id,
      name: menu.menu_name,
      route: menu.route,
      minRoleLevel: menu.min_role_level,
      maxRoleLevel: menu.max_role_level,
    }));
  } catch (error) {
    setError(error.message);
  }

  renderMenus();
}

async function handleMenuSubmit(event) {
  event.preventDefault();

  try {
    const minRoleLevel = Number(document.getElementById('minRoleLevel').value);
    const maxRoleLevel = Number(document.getElementById('maxRoleLevel').value);

    if (maxRoleLevel < minRoleLevel) {
      throw new Error('Max role level must be greater than or equal to min role level');
    }

    elements.saveMenuButton.disabled = true;
    elements.saveMenuButton.textContent = 'Saving...';

    const response = await fetch(`${API_BASE_URL}/menu`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        menu_name: document.getElementById('menuName').value,
        route: document.getElementById('route').value,
        min_role_level: minRoleLevel,
        max_role_level: maxRoleLevel,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || 'Failed to create menu');
    }

    alert(data.message);
    elements.menuModal.hidden = true;
    elements.menuForm.reset();
    await loadMenus();
  } catch (error) {
    alert(error.message);
  } finally {
    elements.saveMenuButton.disabled = false;
    elements.saveMenuButton.textContent = 'Save Menu';
  }
}

initMenuPage();
