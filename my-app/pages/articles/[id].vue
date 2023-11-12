<template>
    <div>
        <h1>Article {{ id }}</h1>
        <p>{{ article.title }}</p>
        <p>{{ article.description }}</p>
        <p>{{ article.date }}</p>
    </div>
</template>

<script setup>
    const { id } =  useRoute().params;
    const apiUrlGo = `${process.env.API_PATH_GO}/api/v1/articles/` + id;

    const { data: article, error, refresh } = await useFetch(apiUrlGo, { key: id }, {
        onResponse({ request, response, options }) {
        console.log("[fetch response]", request, response.status, response.body);
        },
        onRequestError({ request, error, options }) {
        console.error("[fetch request error]", request, error);
        },
    });

    definePageMeta({
        layout: 'default',
    })
</script>

<style scoped>

</style>