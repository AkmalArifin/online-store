<template>
	<div>
		<form v-on:submit="loginClicked">
			<label for="email">Login</label>
			<input type="text" id="login_email" name="email" v-model="email">
			<label for="password">Password</label>
			<input type="text" id="login_password" name="password" v-model="password"/>
			<input type="submit" value="Login">
		</form>

		<button @click="getCookie">Get Cookie</button>
		<p>{{ email }}</p>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

const email = ref("");
const password = ref("");


async function loginClicked(event: Event) {
	event.preventDefault();
	try {
		const data = {
			"email": email.value,
			"password": password.value
		};

		const response = await axios.post("http://localhost:8090/login", data);
		localStorage.setItem("token", response.data.token);
		console.log("Token acquired!");
	} catch (e) {
		console.error(e);
	};
}
</script>

<script lang="ts">
export default {
  name: 'LoginPage',
};
</script>

<style scoped></style>
