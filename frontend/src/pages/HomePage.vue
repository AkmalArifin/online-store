<template>
    <div>
        <Navbar/>
        <Hero/>
    </div>
</template>

<script setup lang="ts">
import Hero from '../components/homepage/Hero.vue';
import Navbar from '../components/navbar/Navbar.vue';


</script>


<script lang="ts">
export default {
    name: 'HomePage',
};
</script>

<style scoped>

</style>