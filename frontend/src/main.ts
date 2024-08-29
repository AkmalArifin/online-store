// import './assets/main.css'

import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";

import App from "./App.vue";
import LoginPage from "./pages/LoginPage.vue";
import RegisterPage from "./pages/RegisterPage.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/login', component: LoginPage},
        {path: '/register', component: RegisterPage}
    ]
});

const app = createApp(App)

app.use(router)
app.mount("#app");
