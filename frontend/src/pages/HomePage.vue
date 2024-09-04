<template>
    <div>
        <Navbar/>
        <Hero/>
        <div class="section col">
            <div class="text">
                <h2>Personal and Unique</h2>
                <p class="p2">Discover the unique charm of handcrafted items—where every piece tells its own story. Unlike mass-produced products, these one-of-a-kind creations offer unmatched personality.</p>
            </div>
            <div class="image">
                <img src="/images/section-personal.jpg" alt="">
            </div>
        </div>
        <div class="section col">
            <div class="image">
                <img src="/images/section-artisans.jpg" alt="">
            </div>
            <div class="text right">
                <h2>Support for Artisans</h2>
                <p class="p2">Celebrate the craftsmanship of local artisans by showcasing their high-quality handmade goods. Support local artisans and attract customers who value handmade goods and the spirit of small businesses.</p>
            </div>
        </div>
        <div class="section col">
            <div class="text">
                <h2>Emotional Connection</h2>
                <p class="p2">The stories behind the products and the people who make them create an emotional connection, making the shopping experience more meaningful.</p>
            </div>
            <div class="image">
                <img src="/images/section-connect.jpg" alt="">
            </div>
        </div>
        <div class="section">

        </div>
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

.section {
    margin: 0px auto;
    padding: 0px 20px;
    max-width: 1080px;
    height: max(100vh, 900px);
    scroll-snap-align: start;
}

.col {
    display: flex;
    gap: 80px;
    flex-wrap: wrap;

    align-items: center;
    align-content: center;
    justify-content: center;
}

.col > div {
    flex : 1;
    min-width: 400px;
}

.image {
    display: flex;
    justify-content: center;
}       

img {
    max-width: 100%;
}
.text h2 {
    margin-bottom: 28px;
}

.right {
    text-align: right;
}


@media (max-width: 920px) {
    .img {
        order: 1;
    }
    .text {
        order: 2;   
    }
}

</style>