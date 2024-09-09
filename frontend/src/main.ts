import './assets/main.css'

import { createApp } from "vue";
import { createRouter, createWebHistory } from "vue-router";

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core';

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

/* import specific icons */
import { faMagnifyingGlass, faCartShopping, faUser, faPaperPlane } from '@fortawesome/free-solid-svg-icons';

/* add icons to the library */
library.add(faMagnifyingGlass);
library.add(faCartShopping);
library.add(faUser);
library.add(faPaperPlane);

import App from "./App.vue";
import LoginPage from "./pages/LoginPage.vue";
import RegisterPage from "./pages/RegisterPage.vue";
import HomePage from "./pages/HomePage.vue";
import StorePage from './pages/StorePage.vue';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: HomePage},
        {path: '/login', component: LoginPage},
        {path: '/register', component: RegisterPage},
        {path: '/store', component: StorePage},
    ]
});

const app = createApp(App)

app.component('font-awesome-icon', FontAwesomeIcon)
app.use(router)
app.mount("#app");
