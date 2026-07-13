const layoutBaseUrl = new URL('.', document.currentScript.src);

const fallbackDrawerHtml = `
  <div class="drawer-backdrop" id="drawerBackdrop"></div>

  <aside class="drawer" aria-label="Dashboard drawer">
    <div class="brand">
      <h2>PanExpress</h2>
      <p>Admin Dashboard</p>
    </div>

    <nav class="nav">
      <a href="../user/user.html" data-page="user">Users Table</a>
      <a href="../menu/menu.html" data-page="menu">Menu Table</a>
      <a href="../menu/menu.html" data-page="makanan">Makanan Table</a>
    </nav>

    <div class="drawer-footer">
      <span id="drawerUser">Signed in as Admin</span>
      <button type="button" class="logout-button" id="logoutButton">Logout</button>
    </div>
  </aside>
`;

const fallbackTopbarHtml = `
  <header class="topbar">
    <button type="button" class="menu-toggle" id="menuToggle" aria-label="Open drawer">=</button>
    <div class="page-title">
      <h1 id="pageTitle">Page Title</h1>
      <p id="pageDescription">Page description goes here.</p>
    </div>
    <div class="admin-chip" id="adminChip">Admin</div>
  </header>
`;

async function loadLayout(activePage, title, description) {
  const drawerHtml = await loadLayoutPartial(
    'drawer.html',
    fallbackDrawerHtml
  );

  const topbarHtml = await loadLayoutPartial(
    'topbar.html',
    fallbackTopbarHtml
  );

  document.getElementById('drawerContainer').innerHTML = drawerHtml;
  document.getElementById('topbarContainer').innerHTML = topbarHtml;

  const activeLink = document.querySelector(`.nav a[data-page="${activePage}"]`);
  if (activeLink) activeLink.classList.add('active');

  const titleEl = document.getElementById('pageTitle');
  const descEl = document.getElementById('pageDescription');
  if (titleEl) titleEl.textContent = title;
  if (descEl) descEl.textContent = description;

  initDrawerEvents();
}

async function loadLayoutPartial(fileName, fallbackHtml) {
  try {
    const response = await fetch(new URL(fileName, layoutBaseUrl));

    if (!response.ok) {
      throw new Error(`${fileName} returned status ${response.status}`);
    }

    return response.text();
  } catch (error) {
    console.warn(`${fileName} could not be loaded. Using fallback layout.`, error);
    return fallbackHtml;
  }
}

function initDrawerEvents() {
  const menuToggle = document.getElementById('menuToggle');
  const drawerBackdrop = document.getElementById('drawerBackdrop');

  menuToggle?.addEventListener('click', () => {
    document.body.classList.add('drawer-open');
  });

  drawerBackdrop?.addEventListener('click', () => {
    document.body.classList.remove('drawer-open');
  });
}
