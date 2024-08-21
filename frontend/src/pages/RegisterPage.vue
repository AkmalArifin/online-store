<template>
    <div>
        <form v-on:submit="registerClicked">
			<label for="username">Username</label>
			<input type="text" id="login_email" name="username" v-model="username">
			<label for="email">Email</label>
			<input type="text" id="login_email" name="email" v-model="email">
			<label for="password">Password</label>
			<input type="text" id="login_password" name="password" v-model="password"/>
			<label for="re_password">Re-type Password</label>
			<input type="text" id="login_password" name="re_password" v-model="rePassword"/>
			<input type="submit" value="Register">
		</form>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const username = ref("");
const email = ref("");
const password = ref("");
const rePassword = ref("")

function registerClicked() {

}

</script>

<script>
export default {
  name: 'RegisterPage',
};
</script>

<style scoped>

</style>