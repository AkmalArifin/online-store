<template>
	<div>
		<form action="http://localhost:8090/login" method="POST">
			<label for="email">Login</label>
			<input type="text" name="email" id="login_email" />
			<label for="password">Password</label>
			<input type="text" name="password" id="login_password" />
			<input type="submit" value="Submit" />
		</form>
	</div>
</template>

<script setup lang="ts">

	
</script>

<style scoped></style>
