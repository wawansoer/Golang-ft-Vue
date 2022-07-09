<template>
  <form @submit.prevent="submit">
    <h1 class="h3 mb-3 fw-normal">Please sign in</h1>

    <input v-model="data.email" type="email" class="form-control" placeholder="Email" required>

    <input v-model="data.password" type="password" class="form-control" placeholder="Password" required>

    <button class="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
  </form>
</template>

<script lang="ts">
import {reactive} from 'vue';
import {useRouter} from "vue-router";
import {useStore} from "vuex";

export default {
  name: "Login",
  setup() {
    const data = reactive({
      email: '',
      password: ''
    });
    
    const store = useStore();
    const router = useRouter();

    const submit = async () => {
      const response = await fetch('http://localhost:3000/api/login', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        credentials: 'include',
        body: JSON.stringify(data)
      });
      const content = await response.json();

      // jikalau gagal mendapatkan data maka kembali ke halaman utama
      if (content.message != 'success'){
          await router.push('/');
      }else{
        await store.dispatch('setAuth', true);
        await router.push('/dashboard');  
      }
    }

    return {
      data,
      submit
    }
  }
}
</script>