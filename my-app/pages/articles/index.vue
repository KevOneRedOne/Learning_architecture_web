<script setup onload>
    definePageMeta({
        layout: 'default',
    })
    const apiUrl = `${process.env.API_PATH_GO}`;
    const { pending, data: articles, error, refresh } = await useFetch(`${apiUrl}/api/v1/articles/`, {
        onResponse({ request, response, options }) {
        console.log("[fetch response]", request, response.status, response.body);
        },
        onRequestError({ request, error, options }) {
        console.error("[fetch request error]", request, error);
        },
    });
    const fetchedArticles = await articles?.value;
</script>

<template>
  <div class="container">
    <div class="grid">
      <div v-if="articles && !pending" v-for="articles in fetchedArticles" :key="article.id">
        <div class="item">
          <h2><NuxtLink :to="`/articles/${articles.id}`">{{ articles.title }}</NuxtLink></h2>
          <p>{{ articles.description }}</p>
          <p>Date : {{ articles.date }}</p>
        </div> 
      </div>
      <div v-if="pending">Chargement en cours...</div>
    </div>
  </div>
</template>

