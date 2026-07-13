const API_BASE_URL = "http://localhost:8080/api";

const form = document.getElementById("loginForm");
const loginButton = document.getElementById("loginButton");
const errorMessage = document.getElementById("errorMessage");

if (localStorage.getItem("token")) {
    window.location.replace("../user/user.html");
}

form.addEventListener("submit", login);

async function login(e) {

    e.preventDefault();

    errorMessage.style.display = "none";

    loginButton.disabled = true;
    loginButton.textContent = "Signing in...";

    const username =
        document.getElementById("username").value.trim();

    const password =
        document.getElementById("password").value;

    try {

        const response =
            await fetch(
                `${API_BASE_URL}/login`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        username,
                        password,
                    }),
                },
            );

        const data =
            await response.json();

        if (!response.ok) {
            throw new Error(
                data.error ||
                "Login failed",
            );
        }

        localStorage.setItem(
            "token",
            data.token,
        );

        localStorage.setItem(
            "user",
            JSON.stringify({
                id: data.user_id,
                username: data.username,
                email: data.email,
                roleLevel: data.role_level,
            }),
        );

        window.location.replace("../user/user.html");

    } catch (error) {

        console.error(error);

        errorMessage.textContent = error.message;
        errorMessage.style.display = "block";

    } finally {

        loginButton.disabled = false;
        loginButton.textContent = "Login";
    }
}