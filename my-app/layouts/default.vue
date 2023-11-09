<template>
    <div id="app">
        <Navbar />
        <slot/>
    </div>
</template>

<script>

    import Navbar from '~/components/Navbar/Navbar.vue';

    export default {
       components: {
           Navbar
       }
    }

</script>