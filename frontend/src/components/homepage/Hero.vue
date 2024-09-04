<template>
    <div class="hero">
        <div class="container">
            <h1>Find Your <span style="color: var(--color-primary); font-weight: bold;">Personal Goods</span></h1>
            <h4>Unique Artisanal Goods, Made by Hand, Made with Heart</h4>
            <form class="search" autocomplete="off">
                <span style="font-size: 28px; color: var(--grey);">
                    <!-- <font-awesome-icon :icon="['fas', 'magnifying-glass']"/> -->
                    <font-awesome-icon icon="fa-solid fa-magnifying-glass"/>
                </span>
                <input class="p1 search-input" type="text" name="search" id="search" placeholder="Macrame Wall Hangings">
            </form>
            <div class="button">
                <p class="p3">Show All Products</p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

</script>

<script lang="ts">
export default {
    name: 'Hero'
}
</script>

<style scoped>

.hero {
    text-align: center;
    background-image: url('/images/hero-bg-2.jpg');
    background-position: top;
    background-size: cover;
    background-repeat: no-repeat;

    max-width: 1440px;
    margin: auto;
    height: 920px;
    padding: 1px;
}

.container {
    margin: 17rem auto;
}

.search {
    margin: auto;
    margin-top: 32px;
    
    height: 60px;
    width: min(80vw, 500px);
    padding: 0px 32px;
    border-radius: 50px;
    border: var(--color-secondary) solid 3px;
    background-color: white;
    
    display: flex;
    align-items: center;
    column-gap: 20px;
}

.magnify {
    border-color: var(--grey);
}

.search-input {
    border: none;
    width: 100%;
}

.search-input:focus {
    outline: none;
}

.button {
    width: 140px;
    height: 36px;
    background-color: var(--color-primary);
    
    margin: auto;
    margin-top: 20px;
    border-radius: 50px;
    padding: 6px 16px;

    cursor: pointer;
}

.button:hover {
    /* transition: var(--transition); */
    box-shadow: var(--shadow);
    /* border-color: var(--black); */
}

.button p {
    font-weight: bold;
    color: var(--white);
}

</style>