<template>
    <div class="footbar">
        <div class="container">
            <div class="section">
                <img class="logo" src="/logo.png" alt="logo-title">
            </div>
            <div class="section">
                <h4>Resources</h4>
                <p><router-link to="/" class="p2">Home</router-link></p>
                <p><router-link to="/store" class="p2">Store</router-link></p>
            </div>
            <div class="section">
                <h4>Company</h4>
                <p><router-link to="/" class="p2">About</router-link></p>
                <p><router-link to="/" class="p2">Careers</router-link></p>
                <p><router-link to="/" class="p2">Terms of service</router-link></p>
            </div>
            <div class="section">
                <h4>Stay updated</h4>
                <input class="email p2" type="text" name="email" id="email" placeholder="Email address">
                <button class="email-button" :click="submitClicked">
                    <font-awesome-icon icon="fa-solid fa-paper-plane" />
                </button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

async function submitClicked() {

}

</script>

<script lang="ts">
export default {
    name: "Footbar"
}
</script>

<style scoped>
.footbar {
    background-color: var(--white);
    height: 340px;
}

.container {
    margin: auto;
    margin-top: 70px;
    max-width: 1028px;

    display: flex;
    justify-content: space-between;

}

.logo {
    max-width: 128px;
}

h4 {
    margin-bottom: 16px;
}

p {
    margin-bottom: 8px;
}

.p2 {
    text-decoration: none;
    color: var(--black);
    cursor: pointer;
}

.email {
    height: 45px;
    width: 200px;
    border-radius: 5px 0px 0px 5px;
    border: 2px solid var(--color-primary);
    padding: 0px 16px;
}

.email-button {
    height: 45px;
    width: 45px;
    background-color: var(--color-primary);
    border-radius: 0px 5px 5px 0px;
    border: 2px solid var(--color-primary);
    padding: 0px;

    color: var(--white);
    font-size: 18px;
}

</style>