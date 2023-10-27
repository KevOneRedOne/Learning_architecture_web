<!-- <template>
    <head>
        <meta charset="utf-8">
        <title>My wonderfull app</title>
        <link rel="preconnect" href="https://fonts.googleapis.com">
        <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
        <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600&family=Space+Grotesk&display=swap" rel="stylesheet">
    </head>
    <body>
        <Header/>
        <main>
           
        </main>
    </body>
</template>

<script>
    import Header from '~/components/Header/Header.vue'
</script> -->

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