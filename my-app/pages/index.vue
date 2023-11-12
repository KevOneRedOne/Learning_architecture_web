<template>
  <div >
    <h2>Home page</h2>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Quod ipsam accusantium vel. Vero laudantium quos nam! Quam, cumque saepe! Facere cum iure quos ipsum voluptatum enim consequatur error eos placeat!</p>

  </div>
</template>

<script setup>
    const config = useRuntimeConfig();

    if (process.server) {
        console.log("server");
    } else {
        console.log("client");
    }
    console.log(config.public.baseAPIGo);
    console.log(config.public.baseAPIFass);

</script>

<style scoped>

</style>
