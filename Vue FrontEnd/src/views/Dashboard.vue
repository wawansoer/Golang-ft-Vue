<template>
  {{ message }}
</template>

<script lang="ts">
import {onMounted, ref} from 'vue';
import {useStore} from "vuex";
import {useRouter} from "vue-router";

export default {
  name: "Dashboard",
  setup() {
    const message = ref('You are not logged in!');
    const store = useStore();
    const router = useRouter();

    onMounted(async () => {
      // cek apakah sudah login 
      try {
        const response = await fetch('http://localhost:3000/api/user', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include'
        });
        const content = await response.json();
        // jikalau belum arahkan ke login 
        if (!content.name){
          await router.push('/login');
        }else{
          message.value = `Selamat Datang ${content.name}`;
          await store.dispatch('setAuth', true);
        }
      } catch (e) {
        await store.dispatch('setAuth', false);
      }
    });

    return {
      message
    }
  }
}
</script>
