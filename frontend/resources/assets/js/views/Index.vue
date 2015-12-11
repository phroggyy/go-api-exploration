<template>
<div class="container">
    <h1 class="text-center">Welcome to a Go Test</h1>
    <h2 class="text-center">We have a few users of our application</h2>
    <table class="table">
        <thead>
            <tr>
                <th>Name</th>
                <th>Email</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="user in users">
                <td>{{ user.name }}</td>
                <td>{{ user.email }}</td>
            </tr>
        </tbody>
    </table>
</div>
</template>

<script>
    export default {
        data() {
            var users;
            return this.$http.get('http://localhost:8080/api', function (data) {
                console.log(data);
                return {users:data}
            });
        }
    }
</script>