<template>
    <header>
        <div class="container">
            <img class="logo" src="/logo.png" alt="logo-title">
            <lu class="container-nav">
                <li>
                    <router-link class="p1" to="/">Homepage</router-link>
                </li>
                <li>
                    <router-link class="p1" to="/store">Store</router-link>
                </li>
                <li>
                    <router-link class="cart" to="/">
                        <font-awesome-icon icon="fa-solid fa-cart-shopping" />
                    </router-link>
                </li>
                <li>
                    <!-- Icon for user is still not able to make -->
                    <router-link class="profile" to="/">
                        <font-awesome-icon icon="fa-solid fa-user" />
                        <!-- <font-awesome-icon icon="fa-regular fa-user" /> -->
                    </router-link>
                </li>
            </lu>
        </div>
    </header>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { ref } from 'vue';

const profileIcon = 'fa-regular fa-user';
const profileIconHover = 'fa-solid fa-user';

const profileClass = ref(profileIcon);

const onHover = () => {
    profileClass.value = profileIconHover;
}

const onLeave = () => {
    profileClass.value = profileIcon;
}

</script>

<script lang="ts">
export default {
    name: 'Navbar',
}
</script>

<style scoped>

header {
    position: fixed;
    background-color: var(--white);
    height: 80px;
    overflow: visible;
    width: 100%;
}

.container {
    /* max-width: 1440px; */
    max-width: 1080px;
    width: 90%;
    height: 100%;
    margin: auto;

    display: flex;
    align-items: center;
    justify-content: space-between;
}

.container-nav {
    display: flex;
    gap: 4em;
    justify-content: flex-end;
    text-align: center;
    list-style-type: none;
}

.container-nav li {
    padding: 0 8px;
    /* min-width: 100%; */
    text-align: center;
}

.p1 {
    text-decoration: none;
    color: var(--black);
}

.p1:hover {
    color: #000000;
    font-weight: bold;
    transition: var(--transition);
}

.logo {
    max-width: 128px;
}

.cart {
    font-size: 32px;
    color: var(--color-primary);
}
.cart:hover {
    color: #607662;
}

.profile {
    font-size: 32px;
    color: var(--color-primary);
}
.profile:hover {
    color: #607662;
}

</style>