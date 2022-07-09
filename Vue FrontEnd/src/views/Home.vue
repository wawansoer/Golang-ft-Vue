<template>
  {{ message }}
</template>

<script lang="ts">
import {onMounted, ref} from 'vue';
import {useStore} from "vuex";
import {useRouter} from "vue-router";

export default {
  name: "Home",
  setup() {
    const message = ref('Anda Belum Login atau Salah Menginputakan Data');
    const store = useStore();
    const router = useRouter();

    onMounted(async () => {
      try {
        const response = await fetch('http://localhost:3000/api/user', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include'
        });

        const content = await response.json();
        // jikalau sudah login paksa ke dashboard
        if (content.name){
          await router.push('/dashboard');
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
